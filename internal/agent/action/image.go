package action

import (
	pb "TzTages/api/proto/v1/images"
	"bufio"
	"context"
	"fmt"
	"github.com/sqweek/dialog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadImage(client pb.ImageServiceClient) {
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

func DownloadImage(client pb.ImageServiceClient, reader *bufio.Reader) {
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

func ListImages(client pb.ImageServiceClient) {
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
