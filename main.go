package main

func main() {
	masterCache = CreateNewCache()
}

//Cache defines the Cache type
type Cache struct {
	keys map[string]interface{}
}

//Get fetches the provided keys value
func (cache *Cache) Get(key string) interface{} {
	return cache.keys[key]
}

//Set allots the provided key the provided value
func (cache *Cache) Set(key string, value interface{}) {
	cache.keys[key] = value
}

var masterCache *Cache

//CreateNewCache inits a new cache
func CreateNewCache() *Cache {
	return &Cache{
		keys: make(map[string]interface{}),
	}
}
