package main

import (
	"log"
	"zajunaApi/internal/server"
)

func main() {
	srv := server.New()
	if err := srv.Run(); err != nil {
		log.Fatal("Error iniciando el servidor:", err)
	}
}
