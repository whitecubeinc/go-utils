package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

var logger = log.New(os.Stderr, "[LOG] ", log.Ldate|log.Ltime)

// Log default logger
func Log(v ...any) {
	logger.Println(v...)
}

func Logf(format string, v ...any) {
	logger.Printf(format, v...)
}

func Error(v ...any) {
	str := fmt.Sprintf(ErrorColor, fmt.Sprintln(v...))
	logger.Print(str)
}

func Errorf(format string, v ...any) {
	str := fmt.Sprintf(ErrorColor, fmt.Sprintln(fmt.Sprintf(format, v...)))
	logger.Printf(str)
}

func Info(v ...any) {
	str := fmt.Sprintf(InfoColor, fmt.Sprintln(v...))
	logger.Print(str)
}

func Warning(v ...any) {
	str := fmt.Sprintf(WarningColor, fmt.Sprintln(v...))
	logger.Print(str)
}

func Debug(v ...any) {
	logger.Print(fmt.Sprintf(DebugColor, fmt.Sprintln(v...)))
}

func Fatal(v ...any) {
	logger.Fatal(fmt.Sprintf(WarningColor, fmt.Sprintln(v...)))
}

func PrintStruct(v any) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}

	Log(string(b))
}

func PrintStackTrace() {
	Log(string(debug.Stack()))
}
