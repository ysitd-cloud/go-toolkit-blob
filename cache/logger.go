package cache

// Logger interface of getter logging
type Logger interface {
	Debugf(format string, vals ...interface{})
}

type NullLogger struct{}

func (l *NullLogger) Debugf(format string, vals ...interface{}) {}
