package zerologr

import "github.com/go-logr/logr"

//nolint:gochecknoglobals
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

func Info(msg string, keysAndValues ...any) {
	logger.Info(msg, keysAndValues...)
}

func Error(err error, msg string, keysAndValues ...any) {
	logger.Error(err, msg, keysAndValues...)
}

func WithName(name string) logr.Logger {
	return logger.WithName(name).WithCallDepth(-1)
}

func WithValues(keysAndValues ...any) logr.Logger {
	return logger.WithValues(keysAndValues...).WithCallDepth(-1)
}

func WithCallDepth(depth int) logr.Logger {
	return logger.WithCallDepth(depth)
}
