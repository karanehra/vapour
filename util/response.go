package util

import (
	"encoding/json"
	"net/http"
)

//SendServerErrorResponse send a status 500 to the provided responseWriter
//with errorMessage in response data
func SendServerErrorResponse(res http.ResponseWriter, errorMessage string) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(errorMessage))
}

//SendSuccessCreatedResponse sends as status 201 to the provided responseWriter
func SendSuccessCreatedResponse(res http.ResponseWriter, responseBody interface{}) {
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(responseBody)
}

//SendSuccessReponse sends a status 200 to the provided responseWriter
func SendSuccessReponse(res http.ResponseWriter, responseBody interface{}) {
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(responseBody)
}

//SendBadRequestResponse sends a status 400 to the provided responseWriter
func SendBadRequestResponse(res http.ResponseWriter, responseBody interface{}) {
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(responseBody)
}

//SendUnauthorizedResponse sends a status 401 to the provided responseWriter
func SendUnauthorizedResponse(res http.ResponseWriter, message string) {
	responseBody := make(map[string]interface{})
	responseBody["message"] = message
	res.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(res).Encode(responseBody)
}
