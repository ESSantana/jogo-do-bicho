package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ESSantana/jogo-do-bicho/internal/routers"
	"github.com/go-chi/chi/v5"
)

func main() {
	log := initLogger()

	chi_router := chi.NewRouter()
	routers.ConfigBetRouter(chi_router, log)

	defer startServer(chi_router)
	fmt.Println("Server listening on port :8080")
}

func startServer(router *chi.Mux) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	if err := http.Serve(listen, router); err != nil {
		panic(err)
	}
}

func initLogger() *log.Logger {
	var out bytes.Buffer
	logger := log.New(&out, "jogo-do-bicho", log.Lshortfile)
	return logger
}
