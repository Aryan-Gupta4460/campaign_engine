package utils

import (
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func InitLogger() {
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(msg string) {
	infoLog.Println(msg)
}

func LogInfof(format string, v ...any) {
	infoLog.Printf(format, v...)
}

func LogError(err error) {
	if err != nil {
		errorLog.Println(err)
	}
}
func LogErrorf(format string, err error) {
	if err != nil {
		errorLog.Printf(format, err)
	}
}
