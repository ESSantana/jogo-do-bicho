package contracts

import "net/http"

type BetController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
}
