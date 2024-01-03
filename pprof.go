package utils

import (
	"os"
	"runtime/pprof"
)

func CreateCPUProfFile(filePath string, function func()) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()
	function()
}
