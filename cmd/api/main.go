package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories"
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/routers"
	"github.com/ESSantana/jogo-do-bicho/internal/services"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ESSantana/jogo-do-bicho/packages/log"
)

var logger log.Logger
var repoManager repo_contracts.RepositoryManager
var serviceManager svc_contracts.ServiceManager
var router *chi.Mux

func main() {
	logger = log.NewLogger(log.DEBUG)
	singletonRepository(context.Background())
	singletonService(logger)
	setupRouter()

	defer startServer(router)
	fmt.Println("Server listening on port :3000")
}

func setupRouter() {
	router = chi.NewRouter()
	router.Use(middleware.Logger)

	routers.ConfigBetRouter(router, logger, serviceManager)
}

func startServer(router *chi.Mux) {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	if err := http.Serve(listen, router); err != nil {
		panic(err)
	}
}

func singletonRepository(ctx context.Context) {
	if repoManager != nil {
		return
	}
	repoManager = repositories.NewRepositoryManager(ctx)
}

func singletonService(logger log.Logger) {
	if serviceManager != nil {
		return
	}
	serviceManager = services.NewServiceManager(logger, repoManager)
}
