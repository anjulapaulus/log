package log

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type fieldLogger struct {
	log     *zap.Logger
	options *logOptions
	traceID string
}

func (l *fieldLogger) Debug(message string, fields ...Field) {
	l.log.Debug(format(l.traceID, message), fields...)
}

func (l *fieldLogger) Info(message string, fields ...Field) {
	l.log.Info(format(l.traceID, message), fields...)
}

func (l *fieldLogger) Warn(message string, fields ...Field) {
	l.log.Warn(format(l.traceID, message), fields...)
}

func (l *fieldLogger) Error(message string, fields ...Field) {
	l.log.Error(format(l.traceID, message), fields...)
}

func (l *fieldLogger) Panic(message string, fields ...Field) {
	l.log.Panic(format(l.traceID, message), fields...)
}

func (l *fieldLogger) Fatal(message string, fields ...Field) {
	l.log.Fatal(format(l.traceID, message), fields...)
}

func (l *fieldLogger) Sync() error {
	return l.log.Sync()
}

func (l *fieldLogger) Named(name string) *fieldLogger {
	return &fieldLogger{
		log:     l.log.Named(name),
		options: l.options,
	}
}

func (l *fieldLogger) WithContext(ctx context.Context) *fieldLogger {
	if l.options.ctxTraceExtractor != nil {
		trace := l.withExtractedTrace(ctx)
		return &fieldLogger{
			log:     l.log,
			options: l.options.copy(),
			traceID: trace,
		}

	}
	return l
}

func format(trace, message string) string {
	return fmt.Sprintf("%s	%s", trace, message)
}

func (l *fieldLogger) NewFieldLog(opt ...Option) *fieldLogger {
	opts := l.options.copy()
	opts.apply(opt...)
	return initLogger(opts)
}

// withExtractedTrace adds the extacted trace value to the event.
func (l *fieldLogger) withExtractedTrace(ctx context.Context) string {
	if l.options.ctxTraceExtractor != nil {
		if trace := l.options.ctxTraceExtractor(ctx); trace != "" {
			return trace
		}
	}
	return ""
}

// New create a new field logger
func NewFieldLog(opt ...Option) *fieldLogger {
	opts := &logOptions{}
	opts.setDefault()
	opts.apply(opt...)
	return initLogger(opts)
}

func initLogger(opts *logOptions) *fieldLogger {
	config := zap.NewProductionEncoderConfig()

	var zapOptions []zap.Option
	var outputEncoder zapcore.Encoder

	if opts.funcPath {
		config.FunctionKey = "Function"
	}

	if opts.filePath {
		config.CallerKey = "Caller"
		zapOptions = append(zapOptions, zap.AddCaller())
	}

	if opts.skipFrameCount != 0 {
		zapOptions = append(zapOptions, zap.AddCallerSkip(opts.skipFrameCount))
	}

	config.EncodeLevel = opts.levelEncoder
	config.EncodeTime = opts.timeEncoder

	switch opts.output {
	case JSONFormat:
		config.EncodeLevel = CapitalLevelEncoder
		outputEncoder = zapcore.NewJSONEncoder(config)

	case TextFormat:
		outputEncoder = zapcore.NewConsoleEncoder(config)
	}

	zapOptions = append(zapOptions, zap.AddStacktrace(PANIC))

	// fileEncoder := zapcore.NewConsoleEncoder(config)
	// consoleEncoder := zapcore.NewConsoleEncoder(config)
	// logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// writer := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		// zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(outputEncoder, zapcore.AddSync(os.Stderr), opts.logLevel),
	)

	// core := zapcore.NewCore(outputEncoder, zapcore.AddSync(os.Stdout), opts.logLevel)
	var log *zap.Logger
	log = zap.New(core, zapOptions...)

	if opts.name != "" {
		log = log.Named(opts.name)
	}
	return &fieldLogger{
		log:     log,
		options: opts,
	}
}
