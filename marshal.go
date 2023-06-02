package utils

import "encoding/json"

func MarshalMust(v any) (b []byte) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return
}

func ReturnUnmarshal[T any](b []byte) (v T) {
	err := json.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}

	return
}
