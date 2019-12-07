package lib

import (
	"crypto/sha1"
	"encoding/hex"
)

//The Cache struct defines a crude cache implementation
type Cache struct {
	KeyCount   int32
	Shards     map[string]*CacheShard
	Counters   map[string]int32
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
	cache.KeyCount++
	shard.Set(keyset)
}

//SetCounter creates a new counter with the provided name in the cache
func (cache *Cache) SetCounter(counterName string) {
	cache.Counters[counterName] = 0
}

//GetCounter returns the counters value
func (cache *Cache) GetCounter(counterName string) int32 {
	return cache.Counters[counterName]
}

//IncrementCounter ups the counter by unity
func (cache *Cache) IncrementCounter(counterName string) {
	cache.Counters[counterName]++
}

//Delete removes the key-value pair from the cache
func (cache *Cache) Delete(key string) {
	shard := cache.GetShard(key)
	cache.KeyCount--
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
