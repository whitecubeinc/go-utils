package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodePhone(t *testing.T) {
	phone := "01012345678"

	assert.NotEqual(t, phone, EncodePhone(phone))
}

func TestDecodePhone(t *testing.T) {
	phone := "01012345678"
	encodePhone := EncodePhone(phone)

	assert.Equal(t, phone, DecodePhone(encodePhone))
}
