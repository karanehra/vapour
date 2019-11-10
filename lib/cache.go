package lib

import (
	"sync"
	"vapour/util"
)

//The Cache struct defines a crude cache implementation
type Cache struct {
	Items      map[string]interface{}
	Maintainer *ExpiryMaintainer
	SyncLock   sync.Mutex
}

//Get fetches the provided keys value
func (cache *Cache) Get(key string) interface{} {
	cache.SyncLock.Lock()
	defer cache.SyncLock.Unlock()
	return cache.Items[key]
}

//Set allots the provided key the provided value
func (cache *Cache) Set(keyset *KeySetter) {
	cache.SyncLock.Lock()
	cache.Items[keyset.Key] = keyset.Value
	cache.SyncLock.Unlock()
	if keyset.Expiry > 0 {
		keyset := ExpiryKey{
			ExpiryEpoch: util.GetMsSinceEpoch() + int64(keyset.Expiry),
			Keyname:     keyset.Key,
		}
		cache.Maintainer.Add(keyset)
	}
}

//Delete removes the key-value pair from the cache
func (cache *Cache) Delete(key string) {
	cache.SyncLock.Lock()
	delete(cache.Items, key)
	cache.SyncLock.Unlock()
}
