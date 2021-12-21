package utils

import "encoding/json"

func ToJson(o interface{}) (string, error) {
	if o == nil {
		return "{}", nil
	}

	if b, err := json.Marshal(o); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
