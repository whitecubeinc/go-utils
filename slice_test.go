package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMapValueFilter(t *testing.T) {
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
		if i%2 == 0 {
			intArr = append(intArr, i)
		}
	}

	assert.Equal(t, intArr, MapValueFilter(dataArr, func(element Data) *int {
		v, _ := strconv.Atoi(element.A)
		if v%2 == 0 {
			return &v
		}
		return nil
	}))

	assert.Equal(t, intArr, MapValueFilter(dataArr, func(element Data) *int {
		if element.B%2 == 0 {
			return &element.B
		}
		return nil
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

func TestDivideSlicePerSize(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	dividedArr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7},
	}

	assert.Equal(t, DivideSlicePerSize(arr, 3), dividedArr)

	dividedArr = [][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
	}

	assert.Equal(t, DivideSlicePerSize(arr, 1), dividedArr)

	dividedArr = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7},
	}

	assert.Equal(t, DivideSlicePerSize(arr, 2), dividedArr)
}

func TestDivideMapPerSize(t *testing.T) {
	mapData := map[int]string{
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
	}

	divideMap := DivideMapPerSize(mapData, 1)
	for key, value := range mapData {
		var exist bool
		for _, divide := range divideMap {
			if v, ok := divide[key]; ok {
				assert.Equal(t, v, value)
				exist = true
				break
			}
		}
		assert.Equal(t, exist, true)
	}

	divideMap = DivideMapPerSize(mapData, 2)
	for key, value := range mapData {
		var exist bool
		for _, divide := range divideMap {
			if v, ok := divide[key]; ok {
				assert.Equal(t, v, value)
				exist = true
				break
			}
		}
		assert.Equal(t, exist, true)
	}

	divideMap = DivideMapPerSize(mapData, 3)
	for key, value := range mapData {
		var exist bool
		for _, divide := range divideMap {
			if v, ok := divide[key]; ok {
				assert.Equal(t, v, value)
				exist = true
				break
			}
		}
		assert.Equal(t, exist, true)
	}
}

func TestSlice2String(t *testing.T) {
	intArray := []int{1, 2, 3}
	intArrayString := Slice2String(intArray)
	assert.Equal(t, "1, 2, 3", intArrayString)

	stringArray := []string{"a", "b", "c"}
	stringArrayString := Slice2String(stringArray)
	assert.Equal(t, "'a', 'b', 'c'", stringArrayString)
}
