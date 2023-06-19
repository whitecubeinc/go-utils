package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNotEmptyFields(t *testing.T) {
	type Data struct {
		Exist1 *string
		Exist2 *int
		Exist3 *[]int
		Exist4 string
		Nil1   *string
		Nil2   *int
		Nil3   *[]int
	}

	exist1 := ""
	exist2 := 1
	exist3 := []int{1, 2, 3}
	data := Data{
		Exist1: &exist1,
		Exist2: &exist2,
		Exist3: &exist3,
		Nil1:   nil,
		Nil2:   nil,
		Nil3:   nil,
	}

	assert.Equal(t, []string{"Exist1", "Exist2", "Exist3"}, GetNotEmptyFields(data))

	data.Exist4 = "any"
	assert.Equal(t, []string{"Exist1", "Exist2", "Exist3", "Exist4"}, GetNotEmptyFields(data))
}
