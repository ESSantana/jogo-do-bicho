package controllers

import (
	"context"
	"fmt"
	"strconv"

	// "log"
	"net/http"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/controllers/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/utils"
)

var betService svc_contracts.BetService

type BetController struct {
	// logger         *log.Logger
	serviceManager svc_contracts.ServiceManager
}

func NewBetController( /*logger *log.Logger,*/ serviceManager svc_contracts.ServiceManager) contracts.BetController {
	return &BetController{
		// logger:         logger,
		serviceManager: serviceManager,
	}
}

func (ctlr *BetController) Create(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	betDTO := utils.ReadBody[dto.Bet](request, response)
	if betDTO.BetType == "" {
		return
	}

	vm, err := betService.Create(ctx, betDTO)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error creating new bet",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"id":   vm.ID,
		"data": vm,
	}
	utils.CreateResponse(&response, http.StatusCreated, responseBody)
}

func (ctlr *BetController) GetAll(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	vm, err := betService.GetAllBets(ctx)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error getting all bets",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	if len(vm) < 1 {
		utils.CreateResponse(&response, http.StatusNoContent, nil)
		return
	}

	responseBody := map[string]interface{}{
		"data": vm,
	}
	utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *BetController) Get(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	queryValues := request.URL.Query()
	unparsedID := queryValues["id"][0]
	id, err := strconv.Atoi(unparsedID)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "id format not supported",
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
		return
	}

	vm, err := betService.GetByID(ctx, int32(id))
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error getting bet by id",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	if vm.ID == 0 {
		utils.CreateResponse(&response, http.StatusNoContent, nil)
		return
	}

	responseBody := map[string]interface{}{
		"data": vm,
	}
	utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *BetController) Update(response http.ResponseWriter, request *http.Request) {

	fmt.Println(request.RequestURI)
	// betService := ctlr.serviceManager.NewBetService()
	// ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	// defer cancel()

	// vm, err := betService.GetAllBets(ctx)
	// if err != nil {
	// 	responseBody := map[string]interface{}{
	// 		"message": "error getting all bets",
	// 	}
	// 	utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
	// 	return
	// }

	// if len(vm) < 1 {
	// 	utils.CreateResponse(&response, http.StatusNoContent, nil)
	// 	return
	// }

	// responseBody := map[string]interface{}{
	// 	"data": vm,
	// }
	// utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *BetController) Delete(response http.ResponseWriter, request *http.Request) {

	fmt.Println(request.RequestURI)
	// betService := ctlr.serviceManager.NewBetService()
	// ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	// defer cancel()

	// vm, err := betService.GetAllBets(ctx)
	// if err != nil {
	// 	responseBody := map[string]interface{}{
	// 		"message": "error getting all bets",
	// 	}
	// 	utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
	// 	return
	// }

	// if len(vm) < 1 {
	// 	utils.CreateResponse(&response, http.StatusNoContent, nil)
	// 	return
	// }

	// responseBody := map[string]interface{}{
	// 	"data": vm,
	// }
	// utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *BetController) getService() {
	if betService != nil {
		return
	}

	betService = ctlr.serviceManager.NewBetService()
}
