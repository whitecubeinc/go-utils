package utils

type M map[string]any

func Struct2M(v any) map[string]any {
	b := MarshalMust(v)
	return ReturnUnmarshal[M](b)
}

// Struct2MWithCase string case function 으로 key 값을 변경하여 반환
func Struct2MWithCase(v any, keyCaseFunction func(string) string) map[string]any {
	mapData := make(map[string]any)
	for key, value := range Struct2M(v) {
		mapData[keyCaseFunction(key)] = value
		delete(mapData, key)
	}

	return mapData
}
