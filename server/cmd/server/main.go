package main

import (
	"database/sql"
	"github.com/Alexander3006/architecture-lab-3/server/infrastructure/db"
	"os"
	"os/signal"
)

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName: "lb3",
		User: "postgres",
		Host: "localhost",
		Password: "30062001",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	services, err := ComposeServices()
	if err != nil {
		panic(err)
	}
	services.Register()

	server := ComposeHttpServer()
	go server.Start(3000)

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel
}
