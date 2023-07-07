package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	intValue := 1
	intValuePtr := ToPtr(intValue)

	assert.Equal(t, intValue, *intValuePtr)

	stringValue := "asd"
	stringValuePtr := ToPtr(stringValue)

	assert.Equal(t, stringValue, *stringValuePtr)

	assert.NotEqual(t, stringValuePtr, intValuePtr)
}
