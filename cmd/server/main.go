package main

import (
	"TzTages/internal/delivery/grpc"
	"TzTages/internal/repository/database/liteSQL"
	"TzTages/pgk"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {
	var port int
	var err error

	l := logrus.New()
	log := logrus.NewEntry(l)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите номер порта от 1 до 65535: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		port, err = pgk.ValidPort(input)
		if err != nil {
			log.Info(err.Error())
			continue
		}
		break
	}

	db, err := liteSQL.NewSQLiteStorage()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	grpc.StartServerRPC(port, db, log)
}
