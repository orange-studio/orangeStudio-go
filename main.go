package main

import (
	"go-study-server/src/config"
	"go-study-server/src/router"
	"log"
)

func main() {
	config.Init()
	log.Print("start server...")
	router.StartHttpServer(9527)
}
