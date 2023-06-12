package utils

import (
	"encoding/base64"
	"strings"
)

func EncodePhone(phone *string) {
	*phone = strings.ReplaceAll(*phone, "-", "")
	*phone = base64.StdEncoding.EncodeToString([]byte(*phone))
}

func DecodePhone(phone *string) {
	phoneByte, err := base64.StdEncoding.DecodeString(*phone)
	if err != nil {
		panic(err)
	}

	*phone = string(phoneByte)
}
