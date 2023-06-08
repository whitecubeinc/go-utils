package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	url := "https://httpbin.org/post"

	type Data struct {
		A string
		B int
		C float64
	}

	reqBody := Data{
		A: "A",
		B: 5049,
		C: 12393.2,
	}

	resBody := Post[struct {
		Json Data
	}](url, reqBody, http.Header{})
	resData := resBody.Json

	assert.Equal(t, reqBody, resData)
}
