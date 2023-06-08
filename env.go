package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ReadDotEnv(fileName string, dest any) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	envMap := make(map[string]string)
	for idx, line := range strings.Split(string(file), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		splitLines := strings.SplitN(line, "=", 2)
		if len(splitLines) != 2 {
			panic(fmt.Sprintf("wrong .env file in %d line", idx))
		}

		key := strings.TrimSpace(splitLines[0])
		value := strings.TrimSpace(splitLines[1])

		envMap[key] = value
	}

	err = json.Unmarshal(MarshalMust(envMap), dest)
	if err != nil {
		panic(err)
	}
}
