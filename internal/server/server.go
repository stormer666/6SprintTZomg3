package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type homeServer struct {
	Loger  *log.Logger
	Server http.Server
}

func LaunchServer(logger *log.Logger) *homeServer {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.FirstHandler)
	router.HandleFunc("/upload", handlers.SecondHandler)

	serv := &homeServer{
		Loger: logger,
		Server: http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
	return serv
}
