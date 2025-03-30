package main

import (
	"TzTages/internal/delivery/grpc"
	"TzTages/internal/repository/database/liteSQL"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	log := logrus.NewEntry(l)

	db := liteSQL.NewSQLiteStorage()
	grpc.StartServerRPC(db, log)
}
