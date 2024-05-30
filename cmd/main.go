package main

import (
	"github.com/rwtnitin/gitstore/internal/config"
	"github.com/rwtnitin/gitstore/internal/server"
)

func main() {
	config := config.LoadConfig()
	apiserver := server.NewApiServer(config)
	apiserver.Run()
}
