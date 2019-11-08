package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	vapour "vapour/cache"
	"vapour/lib"
	"vapour/util"

	"github.com/gorilla/mux"
)

//GetKey handles the get key endpoint
func GetKey(res http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	res.WriteHeader(http.StatusOK)
	value, err := util.GetBytestream(vapour.MasterCache.Get(key))
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	res.Write([]byte(value))
}

//SetKey handles the get key endpoint
func SetKey(res http.ResponseWriter, req *http.Request) {
	var keyInstace lib.KeySetter
	json.NewDecoder(req.Body).Decode(&keyInstace)
	fmt.Println(keyInstace.Value)
	if err := keyInstace.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}

	vapour.MasterCache.Set(keyInstace.Key, keyInstace.Value)
	util.SendSuccessReponse(res, map[string]string{})
}
