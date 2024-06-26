package utils

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
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
	if req.Header.Get("Content-Type") != "application/json" {
		req.Header.Add("Content-Type", "application/json")
	}

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

func Get[T any](url string, param url.Values, header http.Header) (resBody T) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	if len(param) != 0 {
		req.URL.RawQuery = param.Encode()
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
	if req.Header.Get("Content-Type") != "application/json" {
		req.Header.Add("Content-Type", "application/json")
	}

	// send request
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func Put[T any](url string, body any, header http.Header) (resBody T) {
	if body == nil {
		body = map[string]string{}
	}

	client := new(http.Client)

	bodyBytes := MarshalMust(body)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	// set header
	req.Header = header
	if req.Header.Get("Content-Type") != "application/json" {
		req.Header.Add("Content-Type", "application/json")
	}

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

func Patch[T any](url string, body any, header http.Header) (resBody T) {
	if body == nil {
		body = map[string]string{}
	}

	client := new(http.Client)

	bodyBytes := MarshalMust(body)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	// set header
	req.Header = header
	if req.Header.Get("Content-Type") != "application/json" {
		req.Header.Add("Content-Type", "application/json")
	}

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
