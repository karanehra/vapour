package handlers

import (
	"encoding/json"
	"net/http"
	vapour "vapour/cache"
	"vapour/lib"
	"vapour/util"

	"github.com/gorilla/mux"
)

//GetKey handles the get key endpoint
func GetKey(res http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	value := vapour.MasterCache.Get(key)
	util.SendSuccessValueReponse(res, value)
}

//SetKey handles the get key endpoint
func SetKey(res http.ResponseWriter, req *http.Request) {
	var keyInstance lib.KeySetter
	json.NewDecoder(req.Body).Decode(&keyInstance)
	if err := keyInstance.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	vapour.MasterCache.Set(&keyInstance)
	util.SendSuccessReponse(res, map[string]string{})
}
