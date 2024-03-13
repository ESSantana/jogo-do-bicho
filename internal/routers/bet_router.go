package routers

import (
	"github.com/ESSantana/jogo-do-bicho/internal/controllers"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
	"github.com/go-chi/chi/v5"
)

func ConfigBetRouter(router *chi.Mux, logger log.Logger, serviceManager svc_contracts.ServiceManager) {
	configGambler(router, logger, serviceManager)
	configBet(router, logger, serviceManager)
}

func configBet(router *chi.Mux, logger log.Logger, serviceManager svc_contracts.ServiceManager) {
	controller := controllers.NewBetController(logger, serviceManager)

	router.Get("/bet", controller.GetAll)
	router.Get("/bet?id={id}", controller.Get)
	router.Post("/bet", controller.Create)
	router.Put("/bet", controller.Update)
	router.Delete("/bet", controller.Delete)
}

func configGambler(router *chi.Mux, logger log.Logger, serviceManager svc_contracts.ServiceManager) {
	controller := controllers.NewGamblerController(logger, serviceManager)

	router.Get("/gambler", controller.GetAll)
	router.Get("/gambler?id={id}", controller.Get)
	router.Post("/gambler", controller.Create)
	router.Put("/gambler", controller.Update)
	router.Delete("/gambler", controller.Delete)
}
