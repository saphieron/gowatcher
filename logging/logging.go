package logging

import (
	"log"
	"os"
)

var ErrorLog log.Logger

func Init() {
	ErrorLog = *log.New(os.Stderr, "gowatcher: ", 0)
}
