package lib

//LRU is a struct that implements an LRU cache
type LRU struct {
	Entries []map[string]interface{}
	Size    int
}
