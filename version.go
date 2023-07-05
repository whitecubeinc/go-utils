package utils

import (
	"strconv"
	"strings"
)

type AppVersion struct {
	version string
}

func NewAppVersion(version string) *AppVersion {
	return &AppVersion{
		version: version,
	}
}

// IsSupportedVersion 빈 배열은 해당 버전 이상이라고 간주
func (o *AppVersion) IsSupportedVersion(minVersion string) bool {
	if minVersion == o.version || o.version == "" {
		return true
	}

	minVersionTokens := strings.Split(minVersion, ".")
	currentVersionTokens := strings.Split(o.version, ".")
	max := len(minVersionTokens)
	if max < len(currentVersionTokens) {
		max = len(currentVersionTokens)
	}

	v1 := make([]int, max)
	v2 := make([]int, max)

	var err error
	for i, str := range minVersionTokens {
		v1[i], err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
	}

	for i, str := range currentVersionTokens {
		v2[i], err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
	}

	for i := range v1 {
		if v1[i] == v2[i] {
			continue
		}

		return v2[i]-v1[i] > 0
	}

	return false
}
