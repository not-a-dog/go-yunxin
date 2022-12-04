package yunxin

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type emptyLogger struct{}

var _ Logger = new(emptyLogger)

func (e *emptyLogger) Info(args ...interface{})  {}
func (e *emptyLogger) Error(args ...interface{}) {}
