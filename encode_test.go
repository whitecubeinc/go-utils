package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodePhone(t *testing.T) {
	phone := "01012345678"
	phone2 := phone
	EncodePhone(&phone2)

	assert.NotEqual(t, phone, phone2)
}

func TestDecodePhone(t *testing.T) {
	phone := "01012345678"
	phone2 := phone
	EncodePhone(&phone2)
	DecodePhone(&phone2)

	assert.Equal(t, phone, phone2)
}
