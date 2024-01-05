package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init() {
	file, err := os.OpenFile("./logger/wzh.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		panic(err)
	}
	Logger = log.New(file, "wzh's bk:", log.Ldate|log.Ltime|log.Lshortfile)
}
