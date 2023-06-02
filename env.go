package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadDotEnv(fileName string, dest any) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	envMap := make(map[string]string)
	for idx, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			continue
		}
		splitLines := strings.SplitN(line, "=", 2)
		if len(splitLines) != 2 {
			return errors.New(fmt.Sprintf("wrong .env file in %d line", idx))
		}

		key := strings.TrimSpace(splitLines[0])
		value := strings.TrimSpace(splitLines[1])

		envMap[key] = value
	}

	err = json.Unmarshal(MarshalMust(envMap), dest)
	if err != nil {
		return err
	}

	return nil
}
