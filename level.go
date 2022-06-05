package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DEBUG  Level = zap.DebugLevel
	INFO   Level = zap.InfoLevel
	WARN   Level = zap.WarnLevel
	ERROR  Level = zap.ErrorLevel
	DPANIC Level = zap.DPanicLevel
	PANIC  Level = zap.PanicLevel
	FATAL  Level = zap.FatalLevel
)

// Level encoders to serialize levels
var (
	LowercaseLevelEncoder      = zapcore.LowercaseLevelEncoder
	LowercaseColorLevelEncoder = zapcore.LowercaseColorLevelEncoder
	CapitalLevelEncoder        = zapcore.CapitalLevelEncoder
	CapitalColorLevelEncoder   = zapcore.CapitalColorLevelEncoder
)
