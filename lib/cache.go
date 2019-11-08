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

//Validate validates the keySetter values before addition to the cache
func (keySetter *KeySetter) Validate() []string {
	var errorData []string = []string{}
	if keySetter.Key == "" {
		errorData = append(errorData, "field 'key' is required")
	}
	if keySetter.Value == "" {
		errorData = append(errorData, "field 'value' is required")
	}
	if keySetter.Expiry < 0 {
		errorData = append(errorData, "Enter a valid numerical expiry in ms")
	}
	return errorData
}

//Get fetches the provided keys value
func (cache *Cache) Get(key string) interface{} {
	return cache.Items[key]
}

//Set allots the provided key the provided value
func (cache *Cache) Set(key string, value interface{}) {
	cache.Items[key] = value
}
