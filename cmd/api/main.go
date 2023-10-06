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
)

var repoManager repo_contracts.RepositoryManager
var serviceManager svc_contracts.ServiceManager

func main() {
	chi_router := chi.NewRouter()
	chi_router.Use(middleware.Logger)
	SingletonRepository(context.Background())
	SingletonService()

	routers.ConfigBetRouter(chi_router, serviceManager)

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

func SingletonRepository(ctx context.Context) {
	if repoManager != nil {
		return
	}
	repoManager = repositories.NewRepositoryManager(ctx)
}

func SingletonService() {
	if serviceManager != nil {
		return
	}
	serviceManager = services.NewServiceManager(repoManager)
}
