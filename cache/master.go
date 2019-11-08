package vapour

import "vapour/lib"

//MasterCache the master cache instance. Run get and Set on this
var MasterCache *lib.Cache

//InitCache initializes the cache
func InitCache() {
	MasterCache = &lib.Cache{
		Items: make(map[string]interface{}),
	}
}
