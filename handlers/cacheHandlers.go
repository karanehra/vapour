package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	vapour "vapour/cache"
	"vapour/lib"
	"vapour/util"

	"github.com/gorilla/mux"
)

//GetKey handles the get key endpoint
func GetKey(res http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	fmt.Printf("Getting: %v\n", key)
	value := vapour.MasterCache.Get(key)
	if value == nil {
		atomic.AddInt64(&vapour.MasterCache.Misses, 1)
	} else {
		atomic.AddInt64(&vapour.MasterCache.Hits, 1)
	}
	util.SendSuccessValueReponse(res, value)
}

//SetKey handles the get key endpoint
func SetKey(res http.ResponseWriter, req *http.Request) {
	var keyInstance lib.KeySetter
	json.NewDecoder(req.Body).Decode(&keyInstance)
	fmt.Printf("Setting: %v\n", keyInstance.Key)
	if err := keyInstance.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	vapour.MasterCache.Set(&keyInstance)
	util.SendSuccessReponse(res, map[string]string{})
}

//GetCounter handles the get counter endpoint
func GetCounter(res http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["name"]
	count := vapour.MasterCache.GetCounter(key)
	util.SendSuccessReponse(res, count)
}

//IncrementCounter handles the get increase counter endpoint
func IncrementCounter(res http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["name"]
	vapour.MasterCache.IncrementCounter(key)
	util.SendSuccessReponse(res, map[string]string{})
}

//GetStatus is a dummy handler for sending a status 200
func GetStatus(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessReponse(res, nil)
}

//CreateQueue handles the queue creator endpoint
//TODO
func CreateQueue(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessReponse(res, nil)
}

//AddToQueue handles the data enqueue endpoint
//TODO
func AddToQueue(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessReponse(res, nil)
}

//RemoveFromQueue handles the data dequeue endpoint
//TODO
func RemoveFromQueue(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessReponse(res, nil)
}

//GetAllShards handles the getting shard endpoint
func GetAllShards(res http.ResponseWriter, req *http.Request) {
	var shards map[string]*lib.CacheShard = vapour.MasterCache.Shards
	var responseBody map[string]interface{} = map[string]interface{}{}
	responseBody["totalKeyCount"] = vapour.MasterCache.KeyCount
	responseBody["hits"] = vapour.MasterCache.Hits
	responseBody["misses"] = vapour.MasterCache.Misses
	responseBody["gets"] = vapour.MasterCache.Gets
	responseBody["sets"] = vapour.MasterCache.Sets
	responseBody["startupMS"] = vapour.MasterCache.StartupTimeMS
	responseBody["shards"] = []int64{}
	for i := range shards {
		responseBody["shards"] = append(responseBody["shards"].([]int64), shards[i].KeyCount)
	}
	util.SendSuccessReponse(res, responseBody)
}
