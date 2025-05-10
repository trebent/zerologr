package zerologr

import "github.com/go-logr/logr"

var logger = New(&Opts{})

func Set(newLogger logr.Logger) {
	logger = newLogger.WithCallDepth(1)
}

func Enabled() bool {
	return logger.Enabled()
}

func V(level int) logr.Logger {
	return logger.V(level).WithCallDepth(-1)
}

func Info(msg string, keysAndValues ...interface{}) {
	logger.Info(msg, keysAndValues...)
}

func Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Error(err, msg, keysAndValues...)
}

func WithName(name string) logr.Logger {
	return logger.WithName(name)
}

func WithValues(keysAndValues ...interface{}) logr.Logger {
	return logger.WithValues(keysAndValues...)
}

func WithCallDepth(depth int) logr.Logger {
	return logger.WithCallDepth(depth)
}
