package log

import (
	"go.uber.org/zap/zapcore"
)

var (
	EpochTimeEncoder       = zapcore.EpochTimeEncoder
	EpochMillisTimeEncoder = zapcore.EpochMillisTimeEncoder
	EpochNanosTimeEncoder  = zapcore.EpochNanosTimeEncoder
	ISO8601TimeEncoder     = zapcore.ISO8601TimeEncoder
	RFC3339TimeEncoder     = zapcore.RFC3339TimeEncoder
	RFC3339NanoTimeEncoder = zapcore.RFC3339NanoTimeEncoder
)

// create a time encoder from layouts
var TimeEncoderOfLayout = zapcore.TimeEncoderOfLayout
