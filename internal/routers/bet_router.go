package routers

import (
	"log"

	"github.com/ESSantana/jogo-do-bicho/internal/controllers"
	"github.com/go-chi/chi/v5"
)

func ConfigBetRouter(router *chi.Mux, logger *log.Logger) {
	controller := controllers.NewBetController(logger)

	router.Get("/bet", controller.GetAll)
	router.Get("/bet/{id}", controller.Get)
	router.Post("/bet", controller.Create)
	router.Put("/bet", controller.Update)
	router.Delete("/bet", controller.Delete)
}
