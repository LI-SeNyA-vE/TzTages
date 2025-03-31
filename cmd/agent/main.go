package main

import (
	"TzTages/pgk"
	"bufio"
	"context"
	"fmt"
	"github.com/sqweek/dialog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	pb "TzTages/api/proto/v1/images"
)

func main() {
	var port int
	var err error

	runtime.LockOSThread() //обязательно для GUI

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите номер порта  от 1 до 65535 который прописан на сервере: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		port, err = pgk.ValidPort(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	conn, err := grpc.NewClient(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewImageServiceClient(conn)

	for {
		fmt.Print(`
Выберите действие:
	1. Загрузить изображение
	2. Скачать изображение
	3. Показать список
	0. Выйти
> `)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			uploadImage(client)
		case "2":
			downloadImage(client, reader)
		case "3":
			listImages(client)
		case "0":
			fmt.Println("Завершение.")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}

func uploadImage(client pb.ImageServiceClient) {
	path, err := dialog.File().Title("Выбери изображение").Load()
	if err != nil {
		fmt.Println("Возможно пользователь нажал Отмена:", err)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.UploadImage(ctx, &pb.UploadImageRequest{
		FileName: filepath.Base(path),
		Data:     data,
	})
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	fmt.Printf("Загружено: ID=%s | Размер=%d байт\n", resp.ImageMetadata.ImageId, resp.ImageMetadata.SizeBytes)
}

func downloadImage(client pb.ImageServiceClient, reader *bufio.Reader) {
	fmt.Print("Введите ID изображения: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DownloadImage(ctx, &pb.DownloadImageRequest{ImageId: id})
	if err != nil {
		fmt.Println("Ошибка скачивания:", err)
		return
	}

	path, err := dialog.File().Title("Сохранить файл").SetStartFile(resp.ImageMetadata.FileName).Save()
	if err != nil {
		fmt.Println("Отмена или ошибка сохранения:", err)
		return
	}

	err = os.WriteFile(path, resp.Data, 0644)
	if err != nil {
		fmt.Println("Возможно пользователь нажал Отмена:", err)
		return
	}

	fmt.Println("Изображение сохранено в хранилище:", path)
}

func listImages(client pb.ImageServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.ListImages(ctx, &pb.Empty{})
	if err != nil {
		fmt.Println("Ошибка получения списка:", err)
		return
	}

	for _, img := range resp.ImageMetadata {
		fmt.Printf("ID: %s | Файл: %s | Размер: %d байт\n", img.ImageId, img.FileName, img.SizeBytes)
	}
}
