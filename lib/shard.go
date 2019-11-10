package lib

import (
	"vapour/util"
)

//A CacheShard defines a partition in the cache
type CacheShard struct {
	Items  map[string]interface{}
	Parent *Cache
}

//Get fetches the provided keys value
func (shard *CacheShard) Get(key string) interface{} {
	return shard.Items[key]
}

//Set allots the provided key the provided value
func (shard *CacheShard) Set(keyset *KeySetter) {
	shard.Items[keyset.Key] = keyset.Value
	if keyset.Expiry > 0 {
		volatileKey := ExpiryKey{
			ExpiryEpoch: util.GetMsSinceEpoch() + int64(keyset.Expiry),
			Keyname:     keyset.Key,
		}
		shard.Parent.Maintainer.Add(volatileKey)
	}
}

//Delete removes the key-value pair from the cache
func (shard *CacheShard) Delete(key string) {
	delete(shard.Items, key)
}
