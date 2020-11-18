package main

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/Alexander3006/architecture-lab-3/server/infrastructure/db"
)

func NewDbConnection(conf *Config) (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     conf.Db.Name,
		User:       conf.Db.User,
		Host:       conf.Db.Host,
		Password:   conf.Db.Password,
		DisableSSL: conf.Db.DisableSSL,
	}
	return conn.Open()
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	services, err := ComposeServices()
	handle(err)
	services.Register()

	server, err := ComposeHttpServer()
	handle(err)
	go server.Start()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel
}
