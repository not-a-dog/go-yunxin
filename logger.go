package yunxin

import "log"

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type emptyLogger struct{}

var _ Logger = new(emptyLogger)

func (e *emptyLogger) Info(args ...interface{})  {}
func (e *emptyLogger) Error(args ...interface{}) {}

type stdLogger struct{}

var _ Logger = new(stdLogger)

func (e *stdLogger) Info(args ...interface{}) {
	log.Println(args...)
}
func (e *stdLogger) Error(args ...interface{}) {
	log.Println(args...)
}

func StdLogger() Logger {
	return &stdLogger{}
}
