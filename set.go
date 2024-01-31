package utils

// GetDifferenceSet 차집합, set1 - set1
func GetDifferenceSet[V comparable](set1, set2 []V) (result []V) {
	result = make([]V, 0, len(set1))
	existMap := make(map[V]bool)

	for _, v := range set2 {
		existMap[v] = true
	}

	for _, v := range set1 {
		if !existMap[v] {
			result = append(result, v)
		}
	}
	return result
}

// GetIntersectionSet 교집합
func GetIntersectionSet[V comparable](set1, set2 []V) (result []V) {
	result = make([]V, 0, len(set1))
	table := make(map[V]bool)

	for _, ele := range set1 {
		table[ele] = true
	}

	for _, ele := range set2 {
		if table[ele] {
			result = append(result, ele)
		}
	}

	return result
}

func GetUnionSet[V comparable](set1, set2 []V) (result []V) {
	existMap := make(map[V]bool)
	result = make([]V, 0, len(set1))
	for _, v := range set1 {
		if _, exist := existMap[v]; exist {
			continue
		}

		result = append(result, v)
		existMap[v] = true
	}

	for _, v := range set2 {
		if _, exist := existMap[v]; exist {
			continue
		}

		result = append(result, v)
		existMap[v] = true
	}
	return
}
