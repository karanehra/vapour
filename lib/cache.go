package lib

import (
	"crypto/sha1"
	"encoding/hex"
)

//The Cache struct defines a crude cache implementation
type Cache struct {
	Shards     map[string]*CacheShard
	Maintainer *ExpiryMaintainer
}

//Get fetches the provided keys value
func (cache *Cache) Get(key string) interface{} {
	shard := cache.GetShard(key)
	return shard.Get(key)
}

//Set allots the provided key the provided value
func (cache *Cache) Set(keyset *KeySetter) {
	shard := cache.GetShard(keyset.Key)
	shard.Set(keyset)
}

//Delete removes the key-value pair from the cache
func (cache *Cache) Delete(key string) {
	shard := cache.GetShard(key)
	shard.Delete(key)
}

//GetShard returns a pointer to the shard allocated to the provided key
func (cache *Cache) GetShard(key string) *CacheShard {
	shardID := GetShardIdentifierFromKey(key)
	if cache.Shards[shardID] == nil {
		return cache.CreateShard(key)
	}
	return cache.Shards[shardID]
}

//CreateShard spawns a new shard inside the cache
func (cache *Cache) CreateShard(key string) *CacheShard {
	shardID := GetShardIdentifierFromKey(key)
	cache.Shards[shardID] = &CacheShard{
		Items:  make(map[string]interface{}),
		Parent: cache,
	}
	return &CacheShard{
		Items:  make(map[string]interface{}),
		Parent: cache,
	}
}

//GetShardIdentifierFromKey returns the identifier
//of a shard based on the key value.
func GetShardIdentifierFromKey(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))[:2]
}
