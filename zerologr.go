// zerologr is a logr.Logger implementation that uses zerolog for logging.
// It provides a simple interface to log messages with different verbosity levels and
// structured logging capabilities. It supports console output and caller information.
// It is designed to be used as a drop-in replacement for logr.Logger in applications that
// already use logr for logging.
//
// zerologr does not adapt to zerolog levels, it only logs Info and Error level messages.
// It does not support Debug or Trace levels, as it is designed to follow
// the logr interface as closely as possible. It does not make much sense to translate
// between verbosity and zerolog levels if a package wants to make use of logr for logging.
//
// zerologr will change the global zerolog level to InfoLevel, so that it does not log
// Debug or Trace level messages by default. This is to ensure that zerologr behaves
// consistently with the logr interface. It will also set the zerolog
// MessageFieldName and ErrorFieldName, both and more are possible to override using the Set*FieldName
// functions of this package.
package zerologr

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/rs/zerolog"
)

type (
	sink struct {
		// The underlying zerolog.Logger
		logger    *zerolog.Logger
		name      string
		v         int
		callDepth int
	}
	// Opts is the options for the zerologr logger.
	Opts struct {
		// Set to true to log to prettily to console. If false, logs are formatted
		// as JSON.
		Console bool
		// Set to true to log the caller file, functiona and line number.
		Caller bool
		// Set the verbosity level. This is used to filter logs at runtime.
		V int
	}
)

var (
	//nolint:gochecknoglobals
	nameFieldName = "name"
	//nolint:gochecknoglobals
	vFieldName = "v"
)

const staticDepth = 2

//nolint:gochecknoinits
func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//nolint:reassign
	zerolog.MessageFieldName = "msg"
	//nolint:reassign
	zerolog.ErrorFieldName = "err"
}

// SetNameFieldName sets the name field name for loggers.
func SetNameFieldName(name string) {
	nameFieldName = name
}

// SetVFieldName sets the verbosity field name for loggers.
func SetVFieldName(name string) {
	vFieldName = name
}

// SetTimestampFieldName sets the timestamp field name for loggers.
func SetTimestampFieldName(name string) {
	//nolint:reassign
	zerolog.TimestampFieldName = name
}

// SetErrorFieldName sets the error field name for loggers.
func SetErrorFieldName(name string) {
	//nolint:reassign
	zerolog.ErrorFieldName = name
}

// SetMessageFieldName sets the message field name for loggers.
func SetMessageFieldName(name string) {
	//nolint:reassign
	zerolog.MessageFieldName = name
}

// New creates a new logr.Logger with the specified options.
func New(opts *Opts) logr.Logger {
	zerologger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if opts.Caller {
		zerologger = zerologger.With().Caller().Logger()
	}

	if opts.Console {
		zerologger = zerologger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return logr.New(&sink{
		logger: &zerologger,
		v:      opts.V,
	})
}

func (s *sink) SetV(v int) {
	s.v = v
}

func (s *sink) Init(info logr.RuntimeInfo) {
	s.callDepth = info.CallDepth + staticDepth
}

func (s *sink) Enabled(v int) bool {
	return v <= s.v
}

func (s *sink) Info(v int, msg string, keysAndValues ...any) {
	e := s.logger.Info()
	if v > 0 {
		e.Int(vFieldName, v)
	}
	s.msg(e, msg, keysAndValues...)
}

func (s *sink) Error(err error, msg string, keysAndValues ...any) {
	e := s.logger.Err(err)
	s.msg(e, msg, keysAndValues...)
}

func (s *sink) WithValues(keysAndValues ...any) logr.LogSink {
	ns := *s
	nl := ns.logger.With().Fields(keysAndValues).Logger()
	ns.logger = &nl
	return &ns
}

func (s *sink) WithName(name string) logr.LogSink {
	ns := *s
	if ns.name == "" {
		ns.name = name
	} else {
		ns.name = ns.name + "/" + name
	}
	return &ns
}

func (s *sink) WithCallDepth(depth int) logr.LogSink {
	ns := *s
	ns.callDepth += depth
	return &ns
}

func (s *sink) msg(e *zerolog.Event, msg string, keysAndValues ...any) {
	if s.name != "" {
		e.Str(nameFieldName, s.name)
	}

	e.CallerSkipFrame(s.callDepth)
	e.Fields(keysAndValues)
	e.Msg(msg)
}
