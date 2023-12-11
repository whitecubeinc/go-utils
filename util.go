package utils

func Ternary[T any](condition bool, true T, false T) (result T) {
	if condition {
		return true
	}
	return false
}
