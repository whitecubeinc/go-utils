package utils

import (
	"bytes"
	"io"
	"net/http"
)

func Post[T any](url string, body any, header http.Header) (resBody T) {
	if body == nil {
		body = map[string]string{}
	}

	client := new(http.Client)

	bodyBytes := MarshalMust(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	// set header
	req.Header = header

	// send request
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read body
	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	resBody = ReturnUnmarshal[T](resBodyBytes)

	return
}

func Get[T any](url string, param map[string]string, header http.Header) (resBody T) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	if param != nil {
		query := req.URL.Query()
		for k, v := range param {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	// set header
	req.Header = header

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	// read body
	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	resBody = ReturnUnmarshal[T](resBodyBytes)

	return
}

func PostWithoutResponse(url string, body any, header http.Header) {
	client := new(http.Client)

	bodyBytes := MarshalMust(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	// set header
	req.Header = header

	// send request
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
