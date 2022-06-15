package log

// import (
// 	"context"

// 	"go.uber.org/zap"
// )

// type logger struct {
// 	log     *zap.SugaredLogger
// 	options *logOptions
// 	traceID string
// }

// func (l *logger) Debug(message string, fields ...Field) {
// 	l.log.Debugw(format(l.traceID, message), fields...)
// }

// func (l *logger) Info(message string, fields ...Field) {
// 	l.log.Info(format(l.traceID, message), fields...)
// }

// func (l *logger) Warn(message string, fields ...Field) {
// 	l.log.Warn(format(l.traceID, message), fields...)
// }

// func (l *logger) Error(message string, fields ...Field) {
// 	l.log.Error(format(l.traceID, message), fields...)
// }

// func (l *logger) Panic(message string, fields ...Field) {
// 	l.log.Panic(format(l.traceID, message), fields...)
// }

// func (l *logger) Fatal(message string, fields ...Field) {
// 	l.log.Fatal(format(l.traceID, message), fields...)
// }

// func (l *logger) Sync() error {
// 	return l.log.Sync()
// }

// func (l *logger) Named(name string) *logger {
// 	return &logger{
// 		log:     l.log.Named(name),
// 		options: l.options,
// 	}
// }

// func (l *logger) WithContext(ctx context.Context) *logger {
// 	if l.options.ctxTraceExtractor != nil {
// 		trace := l.withExtractedTrace(ctx)
// 		return &logger{
// 			log:     l.log,
// 			options: l.options.copy(),
// 			traceID: trace,
// 		}

// 	}
// 	return l
// }

// func (l *logger) NewLog(opt ...Option) *fieldLogger {
// 	opts := l.options.copy()
// 	opts.apply(opt...)
// 	return initLogger(opts)
// }

// // withExtractedTrace adds the extacted trace value to the event.
// func (l *logger) withExtractedTrace(ctx context.Context) string {
// 	if l.options.ctxTraceExtractor != nil {
// 		if trace := l.options.ctxTraceExtractor(ctx); trace != "" {
// 			return trace
// 		}
// 	}
// 	return ""
// }

// New create a new field logger
// func NewLog(opt ...Option) *logger {
// 	opts := &logOptions{}
// 	opts.setDefault()
// 	opts.apply(opt...)
// 	return initLogger(opts)
// }
