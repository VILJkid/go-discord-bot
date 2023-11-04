package helper

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fatih/color"
)

const (
	logErrorLevel = iota
	logInfoLevel
)

func LogError(format string, a ...any) {
	logGeneric(logErrorLevel, format, a...)
}

func LogInfo(format string, a ...any) {
	logGeneric(logInfoLevel, format, a...)
}

func logGeneric(l int, format string, a ...any) {
	var prefix string
	info := fmt.Sprintf(format, a...)

	switch l {
	case logErrorLevel:
		prefix = "[ERROR]"
		info = color.RedString(info)
	case logInfoLevel:
		prefix = "[INFO]"
		info = color.MagentaString(info)
	}

	if _, file, line, found := runtime.Caller(2); found {
		log.Printf("%s %s:%d %v", prefix, file, line, info)
	}
}
