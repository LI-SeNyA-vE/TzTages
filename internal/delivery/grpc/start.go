package grpc

import (
	"TzTages/internal/delivery/grpc/middleware"
	"TzTages/internal/repository"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"

	pb "TzTages/api/proto/v1/images"
	"TzTages/internal/delivery/grpc/handler"
)

func StartServerRPC(port int, storage repository.StorageImage, log *logrus.Entry) {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	limiter := middleware.NewClientLimiters(10, 100, log) // 1 из возможных мидделвейров

	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer(
		grpc.UnaryInterceptor(limiter.UnaryInterceptor()),
	)
	// регистрируем сервис
	pb.RegisterImageServiceServer(s, handler.NewImageServer(storage, log)) // NewImageServer(storage, log)

	log.Info("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
