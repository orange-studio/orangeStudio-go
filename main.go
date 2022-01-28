package main

import (
	"log"
	"orangestudio-go/src/config"
	"orangestudio-go/src/router"
)

func main() {
	config.Init()
	log.Print("start server...")
	router.StartHttpServer(9527)
}
