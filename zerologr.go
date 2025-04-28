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
	Opts struct {
		// Set to true to log to prettily to console. If false, logs are formatted
		// as JSON.
		Console bool
		// Set to true to log the caller file, functiona and line number.
		Caller bool
	}
)

var (
	nameFieldName = "name"
	vFieldName    = "v"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorFieldName = "err"
}

func SetNameFieldName(name string) {
	nameFieldName = name
}

func SetVFieldName(name string) {
	vFieldName = name
}

func SetTimestampFieldName(name string) {
	zerolog.TimestampFieldName = name
}

func SetErrorFieldName(name string) {
	zerolog.ErrorFieldName = name
}

func SetMessageFieldName(name string) {
	zerolog.MessageFieldName = name
}

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
	})
}

func (s *sink) Init(info logr.RuntimeInfo) {
	s.callDepth = info.CallDepth + 2
}

func (s *sink) Enabled(v int) bool {
	return v <= s.v
}

func (s *sink) Info(v int, msg string, keysAndValues ...any) {
	e := s.logger.Info()
	s.msg(e, msg, keysAndValues...)
}

func (s *sink) Error(err error, msg string, keysAndValues ...any) {
	e := s.logger.Err(err)
	s.msg(e, msg, keysAndValues...)
}

func (s *sink) WithValues(keysAndValues ...any) logr.LogSink {

	return s
}

func (s *sink) WithName(name string) logr.LogSink {
	return s
}

func (s *sink) WithCallDepth(depth int) logr.LogSink {
	return s
}

func (s *sink) msg(e *zerolog.Event, msg string, keysAndValues ...any) {
	if s.name != "" {
		e.Str(nameFieldName, s.name)
	}

	if s.v > 0 {
		e.Int(vFieldName, s.v)
	}

	if len(keysAndValues) > 0 {
		e.Interface("fields", keysAndValues)
	}

	e.Msg(msg)
}
