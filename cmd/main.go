package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "//", log.LstdFlags)
	serv := server.LaunchServer(logger)
	logger.Println("Launching server on the port 8080...")
	err := serv.Server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Launch server error: %v", err)
	}
}
