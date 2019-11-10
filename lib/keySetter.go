package lib

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