package logging

import (
	"log"
)

const (
	ERR  = "[err] "
	INFO = "[info] "
)

func Fatal(v string) {
	log.Fatalf(ERR+"%s", v)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(ERR+format, v...)
}

func Info(v string) {
	log.Printf(INFO+"%s", v)
}

func Infof(format string, v ...interface{}) {
	log.Printf(INFO+format, v...)
}