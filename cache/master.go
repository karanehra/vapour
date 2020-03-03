package cache

import (
	"fmt"
	"time"
	"vapour/lib"
	"vapour/util"
)

//MasterCache the master cache instance. Run get and Set on this
var MasterCache *lib.Cache

//InitCache initializes the cache
func InitCache(expiryTime time.Duration) {
	maintainer := &lib.ExpiryMaintainer{
		Items: make([]lib.ExpiryKey, 0),
		Timer: time.NewTicker(expiryTime),
	}
	MasterCache = &lib.Cache{
		Shards:        make(map[string]*lib.CacheShard),
		Maintainer:    maintainer,
		StartupTimeMS: time.Now().Unix() * 1000,
	}

	go func() {
		for {
			select {
			case t := <-MasterCache.Maintainer.Timer.C:
				items := MasterCache.Maintainer.Items
				var nonExpiredItems []lib.ExpiryKey
				deletionCount := 0
				for _, v := range items {
					if v.ExpiryEpoch < util.GetMsSinceEpoch() {
						MasterCache.Delete(v.Keyname)
						deletionCount++
					} else {
						nonExpiredItems = append(nonExpiredItems, v)
					}
				}
				MasterCache.Maintainer.Items = nonExpiredItems
				fmt.Printf("Cleared %d Entries at %d\n", deletionCount, t.UnixNano())
			}
		}
	}()
}

//GetCache returns a pointer to the master cache
func GetCache() *lib.Cache {
	return MasterCache
}
