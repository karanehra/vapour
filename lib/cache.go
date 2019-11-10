package lib

import (
	"fmt"
	"vapour/util"
)

//The Cache struct defines a crude cache implementation
type Cache struct {
	Items      map[string]interface{}
	Maintainer *ExpiryMaintainer
}

//Get fetches the provided keys value
func (cache *Cache) Get(key string) interface{} {
	return cache.Items[key]
}

//Set allots the provided key the provided value
func (cache *Cache) Set(keyset *KeySetter) {
	cache.Items[keyset.Key] = keyset.Value
	if keyset.Expiry > 0 {
		fmt.Printf("Expiry: %d \n", keyset.Expiry)
		keyset := ExpiryKey{
			ExpiryEpoch: util.GetMsSinceEpoch() + int64(keyset.Expiry),
			Keyname:     keyset.Key,
		}
		cache.Maintainer.Add(keyset)
	}
}

//Delete removes the key-value pair from the cache
func (cache *Cache) Delete(key string) {
	delete(cache.Items, key)
}
