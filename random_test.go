package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	arr := make([]string, 1000)

	for i := range arr {
		v := GenerateRandomNumber(10)
		assert.NotContains(t, arr, v)
		arr[i] = v

	}
}

func TestGenerateRandomString(t *testing.T) {
	arr := make([]string, 1000)

	for i := range arr {
		v := GenerateRandomString(10)
		assert.NotContains(t, arr, v)
		arr[i] = v

	}
}
