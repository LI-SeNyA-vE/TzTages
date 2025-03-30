package middleware

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"strings"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type semaphore chan struct{}

func newSemaphore(limit int) semaphore {
	return make(semaphore, limit)
}

func (s semaphore) acquire(ctx context.Context) error {
	select {
	case s <- struct{}{}:
		return nil // заняли 1 слот
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s semaphore) release() {
	select {
	case <-s: // освободили 1 слот
	default:
	}
}

type clientLimiters struct {
	mu        sync.Mutex
	clients   map[string]*clientSemaphores
	upLimit   int
	listLimit int

	log *logrus.Entry
}

type clientSemaphores struct {
	upload semaphore
	list   semaphore
}

func NewClientLimiters(uploadLimit, listLimit int, log *logrus.Entry) *clientLimiters {
	return &clientLimiters{
		clients:   make(map[string]*clientSemaphores),
		upLimit:   uploadLimit,
		listLimit: listLimit,
		log:       log,
	}
}

func (cl *clientLimiters) getClientSemaphores(ip string) *clientSemaphores {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if s, ok := cl.clients[ip]; ok {
		return s // Возвращаем счётчик если уже есть
	}

	sem := &clientSemaphores{
		upload: newSemaphore(cl.upLimit),
		list:   newSemaphore(cl.listLimit),
	} //Создаём новый счётчик

	cl.clients[ip] = sem //Записываем его в структуру
	return sem
}

func (cl *clientLimiters) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		pr, ok := peer.FromContext(ctx)
		if !ok {
			return nil, errors.New("peer не найден")
		}

		ip, _, err := net.SplitHostPort(pr.Addr.String()) //Извлекаем IP
		if err != nil {
			return nil, err
		}

		client := cl.getClientSemaphores(ip) //Получаем/создаём счётчик для этого IP

		cl.log.Printf("Пришёл запрос с IP=%s, method=%s", ip, info.FullMethod)

		switch {
		case strings.Contains(info.FullMethod, "UploadImage"),
			strings.Contains(info.FullMethod, "DownloadImage"):
			err = client.upload.acquire(ctx)
			if err != nil {
				cl.log.Printf("Лимит превышен: IP=%s, method=%s", ip, info.FullMethod)
				return nil, status.Errorf(codes.ResourceExhausted, "слишком много запросов на загрузку/скачивание")
			}
			defer client.upload.release()

		case strings.Contains(info.FullMethod, "ListImages"):
			err = client.list.acquire(ctx)
			if err != nil {
				cl.log.Printf("Лимит превышен: IP=%s, method=%s", ip, info.FullMethod)
				return nil, status.Errorf(codes.ResourceExhausted, "слишком много запросов на список")
			}
			defer client.list.release()

		default:
			return handler(ctx, req)
		}

		return handler(ctx, req)
	}
}
