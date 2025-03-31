package main

import (
	pb "TzTages/api/proto/v1/images"
	"TzTages/internal/agent/action"
	"TzTages/pgk"
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"runtime"
	"strings"
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
			action.UploadImage(client)
		case "2":
			action.DownloadImage(client, reader)
		case "3":
			action.ListImages(client)
		case "0":
			fmt.Println("Завершение.")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
