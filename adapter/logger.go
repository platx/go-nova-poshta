package adapter

import "log"

type logger interface {
	Debugf(format string, args ...interface{})
}

type noopLogger struct{}

func (l *noopLogger) Debugf(format string, args ...interface{}) {}

type defaultLogger struct{}

func (l *defaultLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}
