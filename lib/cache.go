package lib

//The Cache struct defines a crude cache implementation
type Cache struct {
	Items map[string]interface{}
}

//An ExpiryKey defines a struct to maintain volatile key expiries
type ExpiryKey struct {
	ExpiryEpoch int64
	Keyname     string
}

//An ExpiryMaintainer acts as a store for volatile keys
type ExpiryMaintainer struct {
	Items []ExpiryKey
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
	if keySetter.Value == "" || keySetter.Value == nil {
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
func (cache *Cache) Set(keyset *KeySetter) {
	cache.Items[keyset.Key] = keyset.Value
}
