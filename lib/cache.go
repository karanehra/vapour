package lib

//Cache defines the Cache type
type Cache struct {
	Items map[string]interface{}
}

//KeySetter defines the request body to set a key
type KeySetter struct {
	Key    string      `json:"key"`
	Value  interface{} `json:"value"`
	Expiry int32       `json:"expiry"`
}

//Get fetches the provided keys value
func (cache *Cache) Get(key string) interface{} {
	return cache.Items[key]
}

//Set allots the provided key the provided value
func (cache *Cache) Set(key string, value interface{}) {
	cache.Items[key] = value
}
