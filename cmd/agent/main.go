package main

import (
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
	runtime.LockOSThread() //обязательно для GUI

	conn, err := grpc.NewClient(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewImageServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(`
Выберите действие:
	1. Загрузить изображение
	2. Скачать изображение
	3. Показать список
	0. Выйти
> `)
		choice, _ := reader.ReadString('\n')
		switch strings.TrimSpace(choice) {
		case "1":
			{
				path, err := dialog.File().Title("Выбери изображение").Load()
				path = strings.TrimSpace(path)
				data, err := os.ReadFile(path)
				if err != nil {
					fmt.Println("Ошибка чтения файла:", err)
					continue
				}

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				resp, err := client.UploadImage(ctx, &pb.UploadImageRequest{
					FileName: filepath.Base(path),
					Data:     data,
				})
				if err != nil {
					fmt.Println("Ошибка загрузки:", err)
					continue
				}
				fmt.Printf("Загружено: ID=%s, Размер=%d байт\n", resp.ImageMetadata.ImageId, resp.ImageMetadata.SizeBytes)
			}
		case "2":
			{
				fmt.Print("Введите ID изображения: ")
				id, _ := reader.ReadString('\n')
				id = strings.TrimSpace(id)

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				resp, err := client.DownloadImage(ctx, &pb.DownloadImageRequest{ImageId: id})
				if err != nil {
					fmt.Println("Ошибка скачивания:", err)
					continue
				}

				runtime.LockOSThread()
				path, err := dialog.File().
					Title("Сохранить файл").
					SetStartFile(resp.ImageMetadata.FileName).
					Save()

				if err != nil {
					fmt.Println("Возможно пользователь нажал Отмена", err)
					continue
				}
				os.WriteFile(path, resp.Data, 0644)
				fmt.Println("Изображение сохранено в файл:", resp.ImageMetadata.FileName)
			}

		case "3":
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			resp, err := client.ListImages(ctx, &pb.Empty{})
			if err != nil {
				fmt.Println("Ошибка получения списка:", err)
				continue
			}
			for _, img := range resp.ImageMetadata {
				fmt.Printf("ID: %s | Файл: %s | Размер: %d байт\n", img.ImageId, img.FileName, img.SizeBytes)
			}
		case "0":
			fmt.Println("Завершение.")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
