package utils

// GetDifferenceSet 차집합 arr1 - arr2 (comparable type만 가능합니다..), arr1 순서 보장
func GetDifferenceSet[T comparable](arr1, arr2 []T) []T {
	differSet := make([]T, 0, len(arr1))
	existMap := make(map[T]bool, 0)

	for _, v := range arr2 {
		existMap[v] = true
	}

	for _, v := range arr1 {
		if !existMap[v] {
			differSet = append(differSet, v)
		}
	}
	return differSet
}
