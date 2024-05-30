package main

import (
	"log"

	"github.com/jackc/pgx"
	"github.com/rwtnitin/gitstore/internal/config"
	"github.com/rwtnitin/gitstore/internal/server"
)

func main() {
	config := config.LoadConfig()
	conn, err := pgx.Connect(pgx.ConnConfig{
		Database: config.DB_NAME,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
	})
	if err != nil {
		log.Fatal("unable to connect to database")
	} else {
		log.Printf("connect to database %s", config.DB_NAME)
	}
	defer conn.Close()
	apiserver := server.NewApiServer(config)
	apiserver.Run()
}
