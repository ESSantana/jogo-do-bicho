package contracts

import "net/http"

type BaseController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	Create(response http.ResponseWriter, request *http.Request)
}
