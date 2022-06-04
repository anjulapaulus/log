package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	log *zap.Logger
}

func (l *logger) Debug(message string, fields ...Field) {
	l.log.Debug(message, fields...)
}

func (l *logger) Info(message string, fields ...Field) {
	l.log.Info(message, fields...)
}

func (l *logger) Warn(message string, fields ...Field) {
	l.log.Warn(message, fields...)
}

func (l *logger) Error(message string, fields ...Field) {
	l.log.Error(message, fields...)
}

func (l *logger) Panic(message string, fields ...Field) {
	l.log.Panic(message, fields...)
}

func (l *logger) Fatal(message string, fields ...Field) {
	l.log.Fatal(message, fields...)
}

func (l *logger) Sync() error {
	return l.log.Sync()
}

func (l *logger) Named(name string) *logger {
	return &logger{
		log: l.log.Named(name),
	}
}

// func format(message interface{}) string {
// 	return fmt.Sprintf("%v", message)
// }

// New create a new logger (not support log rotating).
func NewLog(opt ...Option) *logger {
	opts := logOptions{}
	opts.setDefault()
	opts.apply(opt...)

	config := zap.NewProductionEncoderConfig()

	var zapOptions []zap.Option
	var outputEncoder zapcore.Encoder

	switch {
	case opts.funcPath:
		config.FunctionKey = "Function"

	case opts.filePath:
		config.FunctionKey = "Caller"
		zapOptions = append(zapOptions, zap.AddCaller())

	case opts.skipFrameCount != 0:
		zapOptions = append(zapOptions, zap.AddCallerSkip(opts.skipFrameCount))
	}

	switch opts.output {
	case JSONFormat:
		outputEncoder = zapcore.NewJSONEncoder(config)

	case TextFormat:
		outputEncoder = zapcore.NewConsoleEncoder(config)
	}

	zapOptions = append(zapOptions, zap.AddStacktrace(PANIC))

	config.EncodeLevel = zapcore.LevelEncoder(opts.levelEncoder)
	config.EncodeTime = zapcore.TimeEncoder(opts.timeEncoder)

	// fileEncoder := zapcore.NewConsoleEncoder(config)
	// consoleEncoder := zapcore.NewConsoleEncoder(config)
	// logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// writer := zapcore.AddSync(logFile)

	// core := zapcore.NewTee(
	// 	// zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	// 	zapcore.NewCore(outputEncoder, zapcore.AddSync(os.Stdout), opts.logLevel),
	// )

	core := zapcore.NewCore(outputEncoder, zapcore.AddSync(os.Stdout), opts.logLevel)
	log := zap.New(core, zapOptions...)
	return &logger{
		log: log,
	}
}
