package utils

import (
	"encoding/base64"
	"strings"
)

func EncodePhone(phone string) string {
	phone = strings.ReplaceAll(phone, "-", "")
	return base64.StdEncoding.EncodeToString([]byte(phone))
}

func DecodePhone(phone string) string {
	phoneByte, err := base64.StdEncoding.DecodeString(phone)
	if err != nil {
		panic(err)
	}

	return string(phoneByte)
}
