package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadDotEnv(t *testing.T) {
	envMap := map[string]any{
		"A": "A",
		"B": "B=B",
		"C": 123,
	}

	type Env struct {
		A string
		B string
		C int `json:",string"`
	}

	envString := ""
	for key, value := range envMap {
		envString += fmt.Sprintf("%s=%v\n", key, value)
	}

	err := os.WriteFile(".env", []byte(envString), 0644)

	assert.Nil(t, err)

	env := Env{}
	err = ReadDotEnv(".env", &env)

	assert.Nil(t, err)
	assert.Equal(t, env.A, envMap["A"])
	assert.Equal(t, env.B, envMap["B"])
	assert.Equal(t, env.C, envMap["C"])

	err = os.Remove(".env")
	assert.Nil(t, err)
}
