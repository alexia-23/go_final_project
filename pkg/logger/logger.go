package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func Init() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
}

func Get() *log.Logger {
	return logger
}
