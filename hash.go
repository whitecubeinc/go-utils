package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
)

func HashStruct(target any) (string, error) {
	targetBytes, err := json.Marshal(target)
	if err != nil {
		return "", err
	}

	var targetMap map[string]any
	if err = json.Unmarshal(targetBytes, &targetMap); err != nil {
		return "", err
	}

	s := SortMap(targetMap)
	s = SortSimpleMap(s)
	return hash(s), nil
}

func SortMap(target map[string]any) map[string]any {

	sorted := SortSimpleMap(target)

	res := make(map[string]any)
	for k, v := range sorted {
		if tv, s := v.(map[string]any); s {
			res[k] = SortMap(tv)
		} else if tv, s := v.([]any); s {
			res[k] = SortSlice(tv)
		} else {
			res[k] = v
		}
	}

	return res
}

func SortSlice(target []any) []any {
	hashArr := make(map[string]any)
	for _, i := range target {
		var tmpV any
		var ha string
		if ttv, ts := i.(map[string]any); ts {
			tmpV = SortMap(ttv)
			ha = hash(tmpV)
		} else if ttv, ts := i.([]any); ts {
			tmpV = SortSlice(ttv)
			ha = hash(tmpV)
		} else {
			tmpV = i
			ha = tmpV.(string)
		}

		hashArr[ha] = tmpV
	}

	var r []any
	sor := SortSimpleMap(hashArr)
	sortKeys := getSortKeys(sor)
	for _, v := range sortKeys {
		r = append(r, sor[v])
	}

	return r
}

func SortSimpleMap(target map[string]any) map[string]any {

	keys := getSortKeys(target)
	res := make(map[string]any, len(keys))
	for _, k := range keys {
		res[k] = target[k]
	}

	return res
}

func getSortKeys(target map[string]any) []string {
	keys := make([]string, 0, len(target))
	for k := range target {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func hash(t any) string {
	sortBytes, _ := json.Marshal(t)
	hashBytes := md5.Sum(sortBytes)
	return hex.EncodeToString(hashBytes[:])
}
