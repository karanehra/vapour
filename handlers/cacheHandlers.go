package handlers

import (
	"encoding/json"
	"net/http"
	"vapour/lib"

	"github.com/gorilla/mux"
)

//GetKey handles the get key endpoint
func GetKey(res http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(key))
}

//SetKey handles the get key endpoint
func SetKey(res http.ResponseWriter, req *http.Request) {
	var keyInstace lib.KeySetter
	json.NewDecoder(req.Body).Decode(&keyInstace)
}
