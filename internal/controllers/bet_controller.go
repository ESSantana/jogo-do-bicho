package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ESSantana/jogo-do-bicho/internal/controllers/contracts"
)

type BetController struct {
	logger *log.Logger
}


func NewBetController(logger *log.Logger) contracts.BetController {
	return &BetController{
		logger: logger,
	}
}

func (controller *BetController) GetAll(response http.ResponseWriter, request *http.Request) {
	a := map[string]interface{}{
		"response": true,
	}
	byted, _ := json.Marshal(a)
	response.Header().Set("Content-Type", "application/json")
	response.Write(byted)
	controller.logger.Println("All")
}

func (controller *BetController) Get(response http.ResponseWriter, request *http.Request) {
	a := map[string]interface{}{
		"response": true,
	}
	byted, _ := json.Marshal(a)
	response.Header().Set("Content-Type", "application/json")
	response.Write(byted)
	controller.logger.Println("All")
}

func (controller *BetController) Create(response http.ResponseWriter, request *http.Request) {
	a := map[string]interface{}{
		"response": true,
	}
	byted, _ := json.Marshal(a)
	response.Header().Set("Content-Type", "application/json")
	response.Write(byted)
	controller.logger.Println("All")
}

func (controller *BetController) Update(response http.ResponseWriter, request *http.Request) {
	a := map[string]interface{}{
		"response": true,
	}
	byted, _ := json.Marshal(a)
	response.Header().Set("Content-Type", "application/json")
	response.Write(byted)
	controller.logger.Println("All")
}

func (controller *BetController) Delete(response http.ResponseWriter, request *http.Request) {
	a := map[string]interface{}{
		"response": true,
	}
	byted, _ := json.Marshal(a)
	response.Header().Set("Content-Type", "application/json")
	response.Write(byted)
	controller.logger.Println("All")
}
