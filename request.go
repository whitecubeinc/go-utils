package utils

import (
	"bytes"
	"io"
	"net/http"
)

func Post[T any](url string, body any, header http.Header) (resBody T, err error) {
	if body == nil {
		body = map[string]string{}
	}

	client := new(http.Client)

	bodyBytes := MarshalMust(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	// set header
	req.Header = header

	// send request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// read body
	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	resBody = ReturnUnmarshal[T](resBodyBytes)

	return
}
