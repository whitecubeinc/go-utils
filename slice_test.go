package utils

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestContains(t *testing.T) {
	maxInt := 10

	intArr := make([]int, 0)
	stringArr := make([]string, 0)
	for i := 0; i <= maxInt; i++ {
		// check int
		assert.False(t, Contains(intArr, func(value int) bool {
			return i == value
		}))
		intArr = append(intArr, i)
		assert.True(t, Contains(intArr, func(value int) bool {
			return i == value
		}))

		// check string
		assert.False(t, Contains(stringArr, func(value string) bool {
			return strconv.Itoa(i) == value
		}))
		stringArr = append(stringArr, strconv.Itoa(i))
		assert.True(t, Contains(stringArr, func(value string) bool {
			return strconv.Itoa(i) == value
		}))
	}

	assert.False(t, Contains(intArr, func(value int) bool {
		return value == maxInt+1
	}))
	assert.False(t, Contains(stringArr, func(value string) bool {
		return value == strconv.Itoa(maxInt+1)
	}))
}

func TestMapValue(t *testing.T) {
	type Data struct {
		A string
		B int
	}

	dataArr := make([]Data, 0)
	intArr := make([]int, 0)
	for i := 0; i < 10; i++ {
		dataArr = append(dataArr, Data{
			A: strconv.Itoa(i),
			B: i,
		})
		intArr = append(intArr, i)
	}

	assert.Equal(t, intArr, MapValue(dataArr, func(element Data) int {
		v, _ := strconv.Atoi(element.A)
		return v
	}))

	assert.Equal(t, intArr, MapValue(dataArr, func(element Data) int {
		return element.B
	}))
}

func TestMapValueUnique(t *testing.T) {
	type Data struct {
		A string
		B int
	}

	dataArr := make([]Data, 0)
	intArr := []int{0, 1}
	for i := 0; i < 10; i++ {
		dataArr = append(dataArr, Data{
			A: strconv.Itoa(i),
			B: i,
		})
	}

	assert.Equal(t, intArr, MapValueUnique(dataArr, func(element Data) int {
		v, _ := strconv.Atoi(element.A)

		return v % 2
	}))

	assert.Equal(t, intArr, MapValueUnique(dataArr, func(element Data) int {
		return element.B % 2
	}))
}
