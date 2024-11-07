package controllers

import (
	"context"
	"fmt"
	"strconv"

	"net/http"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/controllers/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/utils"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
	"github.com/go-chi/chi/v5"
)

var betService svc_contracts.BetService

type BetController struct {
	logger         log.Logger
	serviceManager svc_contracts.ServiceManager
}

func NewBetController(logger log.Logger, serviceManager svc_contracts.ServiceManager) contracts.BetController {
	return &BetController{
		logger:         logger,
		serviceManager: serviceManager,
	}
}

func (ctlr *BetController) Create(response http.ResponseWriter, request *http.Request) {
	ctlr.logger.Debug("Create")
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	betDTO := utils.ReadBody[dto.Bet](request, response)
	if betDTO.GamblerID < 1 {

	}

	result, err := betService.Create(ctx, betDTO)
	if err != nil {
		ctlr.logger.Errorf("error creating new bet: %s", err.Error())
		if errVal, ok := err.(*errors.ValidationError); ok {
			responseBody := map[string]interface{}{
				"error":   true,
				"message": errVal.Error(),
			}
			utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
			return
		}
		responseBody := map[string]interface{}{
			"message": "error creating new bet",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"id":   result.ID,
		"data": result,
	}
	utils.CreateResponse(&response, http.StatusCreated, responseBody)
}

func (ctlr *BetController) GetAll(response http.ResponseWriter, request *http.Request) {
	ctlr.logger.Debug("GetAll")
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	result, err := betService.GetAll(ctx)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": fmt.Sprintf("error getting all bets: %s", err.Error()),
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	if len(result) < 1 {
		utils.CreateResponse(&response, http.StatusNoContent, nil)
		return
	}

	responseBody := map[string]interface{}{
		"data": result,
	}
	utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *BetController) Get(response http.ResponseWriter, request *http.Request) {
	ctlr.logger.Debug("Get")
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	unparsedID := chi.URLParam(request, "id")
	id, err := strconv.Atoi(unparsedID)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "id format not supported",
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
		return
	}

	result, err := betService.GetByID(ctx, int64(id))
	if err != nil {
		ctlr.logger.Errorf("error getting bet by ID %s: %s", id, err.Error())
		responseBody := map[string]interface{}{
			"message": "error getting bet by id",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	if result.ID == 0 {
		utils.CreateResponse(&response, http.StatusNoContent, nil)
		return
	}

	responseBody := map[string]interface{}{
		"data": result,
	}
	utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *BetController) Update(response http.ResponseWriter, request *http.Request) {
	ctlr.logger.Debug("Update")
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	betDTO := utils.ReadBody[dto.Bet](request, response)
	updated, err := betService.Update(ctx, betDTO)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error updating bet",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"updated": updated,
	}
	utils.CreateResponse(&response, http.StatusCreated, responseBody)
}

func (ctlr *BetController) GetBetOptions(response http.ResponseWriter, request *http.Request) {
	ctlr.logger.Debug("BetOptions")
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	betDTO := utils.ReadBody[dto.Bet](request, response)
	updated, err := betService.Update(ctx, betDTO)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error updating bet",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"updated": updated,
	}
	utils.CreateResponse(&response, http.StatusCreated, responseBody)
}


func (ctlr *BetController) getService() {
	if betService != nil {
		return
	}

	betService = ctlr.serviceManager.NewBetService()
}
