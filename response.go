package utils

import (
	"errors"
	"net/url"
	"strings"
)

func ParseQuerySkipUnescape(values url.Values, query string) (err error) {
	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")
		if strings.Contains(key, ";") {
			err = errors.New("invalid semicolon separator in query")
			continue
		}
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")
		values[key] = append(values[key], value)
	}
	return err
}
