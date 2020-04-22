package logger

import (
	"fmt"
	"log"
	"os"
)

var loggerErr = log.New(os.Stderr, "[Error]\t|", log.LstdFlags | log.Lshortfile)

func Logln(v ...interface{}) {
	loggerErr.Output(2, fmt.Sprintln(v...))
}

func Logf(format string, v ...interface{}) {
	loggerErr.Output(2, fmt.Sprintf(format, v...))
}
