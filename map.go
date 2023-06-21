package utils

type M map[string]any

func Struct2M(v any) M {
	b := MarshalMust(v)
	return ReturnUnmarshal[M](b)
}
