package log

import (
	"context"
	"io"
	"os"
	"time"

	"go.uber.org/zap/zapcore"
)

type OUTPUT_FORMAT string

const (
	TextFormat OUTPUT_FORMAT = `text`
	JSONFormat OUTPUT_FORMAT = `json`
)

// logOptions contains all the configuration options for the logger.
type logOptions struct {
	name              string
	colors            bool
	logLevel          Level
	filePath          bool
	funcPath          bool
	skipFrameCount    int
	writer            io.Writer
	output            OUTPUT_FORMAT
	timeEncoder       func(time.Time, zapcore.PrimitiveArrayEncoder)
	levelEncoder      func(Level, zapcore.PrimitiveArrayEncoder)
	ctxTraceExtractor func(ctx context.Context) string
}

type Option func(*logOptions)

// apply applies given configuration values to the logger.
func (lOpts *logOptions) apply(options ...Option) {
	for _, opt := range options {
		opt(lOpts)
	}
}

// setDefault applies default values configurations for the logger
func (lOpts *logOptions) setDefault() {
	lOpts.skipFrameCount = 1
	lOpts.colors = true
	lOpts.logLevel = ERROR
	lOpts.filePath = false
	lOpts.funcPath = false
	lOpts.writer = os.Stdout
	lOpts.output = TextFormat
	lOpts.timeEncoder = TimeEncoderOfLayout("2006-01-02 15:04:05")
	lOpts.levelEncoder = CapitalLevelEncoder
}

// copy returns a copy of existing configuration values of the logger.
func (lOpts *logOptions) copy() *logOptions {
	return &logOptions{
		name:              lOpts.name,
		colors:            lOpts.colors,
		logLevel:          lOpts.logLevel,
		filePath:          lOpts.filePath,
		funcPath:          lOpts.funcPath,
		skipFrameCount:    lOpts.skipFrameCount,
		writer:            lOpts.writer,
		output:            lOpts.output,
		timeEncoder:       lOpts.timeEncoder,
		levelEncoder:      lOpts.levelEncoder,
		ctxTraceExtractor: lOpts.ctxTraceExtractor,
	}
}

// WithName sets a prefixed value to be logged
func WithName(name string) Option {
	return func(opts *logOptions) {
		opts.name = name
	}
}

// WithStdOut sets the log writer.
func WithColors(color bool) Option {
	return func(opts *logOptions) {
		opts.colors = color
	}
}

/* WithLogLevel sets the log level.
This would determine what types of logs would be logged based on the precedence of the log level*/
func WithLogLevel(lvl Level) Option {
	return func(opts *logOptions) {
		opts.logLevel = lvl
	}
}

func WithLevelEncoder(enc func(Level, zapcore.PrimitiveArrayEncoder)) Option {
	return func(opts *logOptions) {
		opts.levelEncoder = enc
	}
}

func WithTimeEncoder(enc func(time.Time, zapcore.PrimitiveArrayEncoder)) Option {
	return func(opts *logOptions) {
		opts.timeEncoder = enc
	}
}

// WithStdOut sets the log writer.
func WithStdOut(w io.Writer) Option {
	return func(opts *logOptions) {
		opts.writer = w
	}
}

// WithSkipFrameCount sets the frame count to skip when reading filepath, func path.
func WithSkipFrameCount(c int) Option {
	return func(opts *logOptions) {
		opts.skipFrameCount = c
	}
}

// WithOutput sets the output format for log entries. Either JSON or text based.
func WithOutputFormat(o OUTPUT_FORMAT) Option {
	return func(opts *logOptions) {
		opts.output = o
	}
}

// WithFilePath sets whether the file path is logged or not.
func WithFilePath(enabled bool) Option {
	return func(opts *logOptions) {
		opts.filePath = enabled
	}
}

// WithFuncPath sets whether the function name is logged or not.
func WithFuncPath(enabled bool) Option {
	return func(opts *logOptions) {
		opts.funcPath = enabled
	}
}

// WithCtxTraceExtractor allows setting up of a function to extract trace from the context.
// Default value func(_ context.Context) string{return ""}
func WithCtxTraceExtractor(fn func(ctx context.Context) string) Option {
	return func(opts *logOptions) {
		opts.ctxTraceExtractor = fn
	}
}
