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

func Slice2Map[T any, V comparable](slice []T, getKey func(v T) V) map[V]T {
	newMap := make(map[V]T)
	for _, v := range slice {
		newMap[getKey(v)] = v
	}

	return newMap
}

func MapFilterByValue[K comparable, T any](targetMap map[K]T, check func(value T) bool) {
	for key, value := range targetMap {
		if !check(value) {
			delete(targetMap, key)
		}
	}
}

func MapFilterByKey[K comparable, T any](targetMap map[K]T, check func(key K) bool) {
	for key := range targetMap {
		if !check(key) {
			delete(targetMap, key)
		}
	}
}
