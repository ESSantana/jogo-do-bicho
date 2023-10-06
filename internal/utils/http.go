package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func CreateResponse(response *http.ResponseWriter, statusCode int, body map[string]interface{}) {
	bodyResponse, err := json.Marshal(body)
	if err != nil {
		panic(errors.New("something went wrong on http utils "))
	}
	(*response).WriteHeader(statusCode)
	(*response).Write(bodyResponse)
	(*response).Header().Set("Content-Type", "application/json")
}
