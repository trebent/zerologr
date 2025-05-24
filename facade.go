package zerologr

import "github.com/go-logr/logr"

//nolint:gochecknoglobals
var logger = New(&Opts{})

// Set replaces the global logger with a new one.
func Set(newLogger logr.Logger) {
	// Set replaces the global logger with a new one.
	// We need to adjust the call depth because the logger is called from this package, meaning an
	// extra frame is added to the stack trace.
	logger = newLogger.WithCallDepth(1)
}

// Enabled returns true if the global logger is enabled for the current verbosity level.
func Enabled() bool {
	return logger.Enabled()
}

// V returns the global logger with the specified verbosity level.
func V(level int) logr.Logger {
	// Have to remove one frame of call depth because the stack depth is reduced from using the
	// returned logger.
	return logger.V(level).WithCallDepth(-1)
}

// Info logs an info message with the global logger.
func Info(msg string, keysAndValues ...any) {
	logger.Info(msg, keysAndValues...)
}

// Error logs an error message with the global logger.
func Error(err error, msg string, keysAndValues ...any) {
	logger.Error(err, msg, keysAndValues...)
}

// WithName returns a new logger with the specified name, using the global logger as a base.
func WithName(name string) logr.Logger {
	// Have to remove one frame of call depth because the stack depth is reduced from using the
	// returned logger.
	return logger.WithName(name).WithCallDepth(-1)
}

// WithValues returns a new logger with the specified key-value pairs, using the global logger as
// a base.
func WithValues(keysAndValues ...any) logr.Logger {
	// Have to remove one frame of call depth because the stack depth is reduced from using the
	// returned logger.
	return logger.WithValues(keysAndValues...).WithCallDepth(-1)
}

// WithCallDepth returns a new logger with the specified call depth, using the global logger as a
// base.
func WithCallDepth(depth int) logr.Logger {
	return logger.WithCallDepth(depth)
}
