package vapour

import (
	"fmt"
	"time"
	"vapour/lib"
)

//MasterCache the master cache instance. Run get and Set on this
var MasterCache *lib.Cache

//InitCache initializes the cache
func InitCache() {
	maintainer := &lib.ExpiryMaintainer{
		Items: make([]lib.ExpiryKey, 0),
		Timer: time.NewTicker(10 * time.Second),
	}
	MasterCache = &lib.Cache{
		Items:      make(map[string]interface{}),
		Maintainer: maintainer,
	}

	go func() {
		for {
			select {
			case t := <-MasterCache.Maintainer.Timer.C:
				fmt.Printf("Master time ticks %d \n", t.UnixNano())
			}
		}
	}()
}
