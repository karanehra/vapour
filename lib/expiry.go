package lib

import "time"

//An ExpiryKey defines a struct to maintain volatile key expiries
type ExpiryKey struct {
	ExpiryEpoch int64
	Keyname     string
}

//An ExpiryMaintainer acts as a store for volatile keys
type ExpiryMaintainer struct {
	Items []ExpiryKey
	Timer *time.Ticker
}

//Add creates a new volatile key entry in the maintainer
func (maintainer *ExpiryMaintainer) Add(key ExpiryKey) {
	maintainer.Items = append(maintainer.Items, key)
}
