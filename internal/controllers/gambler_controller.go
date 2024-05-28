package controllers

import (
	"context"
	"fmt"
	"strconv"

	"net/http"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/controllers/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	svc_contracts "github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/utils"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
	"github.com/go-chi/chi/v5"
)

const DefaultTimeout = 1 * time.Second

var gamblerService svc_contracts.GamblerService

type GamblerController struct {
	logger         log.Logger
	serviceManager svc_contracts.ServiceManager
}

func NewGamblerController(logger log.Logger, serviceManager svc_contracts.ServiceManager) contracts.GamblerController {
	return &GamblerController{
		logger:         logger,
		serviceManager: serviceManager,
	}
}

func (ctlr *GamblerController) Create(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), DefaultTimeout)
	defer cancel()

	gamblerDTO := utils.ReadBody[dto.Gambler](request, response)
	if err := gamblerDTO.Validate(); err != nil {
		responseBody := map[string]interface{}{
			"message": fmt.Sprintf("validation error: %s", err.Error()),
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
		return
	}

	vm, err := gamblerService.Create(ctx, gamblerDTO)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": fmt.Sprintf("error at create gambler: %s", err.Error()),
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

func (ctlr *GamblerController) GetAll(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), DefaultTimeout)
	defer cancel()

	vm, err := gamblerService.GetAll(ctx)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": fmt.Sprintf("error getting all gamblers: %s", err.Error()),
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

func (ctlr *GamblerController) Get(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), DefaultTimeout)
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

	vm, err := gamblerService.GetByID(ctx, int32(id))
	if vm.ID == 0 {
		utils.CreateResponse(&response, http.StatusNotFound, nil)
		return
	}

	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error getting gambler by id",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"data": vm,
	}
	utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *GamblerController) Update(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), DefaultTimeout)
	defer cancel()

	gamblerDTO := utils.ReadBody[dto.Gambler](request, response)
	if err := gamblerDTO.Validate(); err != nil {
		responseBody := map[string]interface{}{
			"message": fmt.Sprintf("validation error: %s", err.Error()),
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
		return
	}

	if gamblerDTO.ID <= 0 {
		responseBody := map[string]interface{}{
			"message": fmt.Sprintf("invalid gambler ID: %d", gamblerDTO.ID),
		}
		utils.CreateResponse(&response, http.StatusBadRequest, responseBody)
		return
	}

	updated, err := gamblerService.Update(ctx, gamblerDTO)
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error updating gambler",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"updated": updated,
	}
	utils.CreateResponse(&response, http.StatusCreated, responseBody)
}

func (ctlr *GamblerController) Delete(response http.ResponseWriter, request *http.Request) {
	ctlr.getService()
	ctx, cancel := context.WithTimeout(request.Context(), DefaultTimeout)
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

	deleted, err := gamblerService.Delete(ctx, int32(id))
	if err != nil {
		responseBody := map[string]interface{}{
			"message": "error deleting gambler",
		}
		utils.CreateResponse(&response, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody := map[string]interface{}{
		"deleted": deleted,
	}
	utils.CreateResponse(&response, http.StatusOK, responseBody)
}

func (ctlr *GamblerController) getService() {
	if gamblerService != nil {
		return
	}

	gamblerService = ctlr.serviceManager.NewGamblerService()
}
