package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalMust(t *testing.T) {
	mapData := map[string]string{
		"A": "A",
		"B": "B",
		"C": "D",
	}
	b := MarshalMust(mapData)

	newMapData := make(map[string]string)
	err := json.Unmarshal(b, &newMapData)

	assert.NotEmpty(t, b)
	assert.Nil(t, err)
	assert.Equal(t, mapData, newMapData)
}

func TestReturnUnmarshal(t *testing.T) {
	mapData := map[string]string{
		"A": "A",
		"B": "B",
		"C": "D",
	}
	b, _ := json.Marshal(mapData)

	newMapData := ReturnUnmarshal[map[string]string](b)

	assert.Equal(t, mapData, newMapData)
}
