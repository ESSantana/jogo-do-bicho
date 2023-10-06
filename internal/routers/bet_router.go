package routers

import (
	"github.com/ESSantana/jogo-do-bicho/internal/controllers"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/go-chi/chi/v5"
)

func ConfigBetRouter(router *chi.Mux, serviceManager svc_contracts.ServiceManager) {
	controller := controllers.NewBetController(serviceManager)

	router.Get("/bet", controller.GetAll)
	router.Get("/bet?id={id}", controller.Get)
	router.Post("/bet", controller.Create)
}
