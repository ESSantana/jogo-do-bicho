package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/ESSantana/jogo-do-bicho/internal/routers"
	"github.com/go-chi/chi/v5"
)

func main() {
	log := initLogger()

	chi_router := chi.NewRouter()

	routers.ConfigBetRouter(chi_router, log)

	err := http.ListenAndServe(":8080", chi_router)
	if err != nil {
		panic(err)
	}
}

func initLogger() *log.Logger {
	var out bytes.Buffer
	logger := log.New(&out, "jogo-do-bicho", log.Lshortfile)
	return logger
}
