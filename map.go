package utils

type M map[string]any

func Struct2M(v any) map[string]any {
	b := MarshalMust(v)
	return ReturnUnmarshal[M](b)
}
