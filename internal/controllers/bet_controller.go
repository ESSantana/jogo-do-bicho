package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	// "log"
	"net/http"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/controllers/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/utils"
)

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
	betService := ctlr.serviceManager.NewBetService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	var betDTO dto.Bet
	body, err := io.ReadAll(request.Body)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "bet payload malformed",
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
		return
	}

	err = json.Unmarshal(body, &betDTO)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "bet payload malformed",
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
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
	betService := ctlr.serviceManager.NewBetService()
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
