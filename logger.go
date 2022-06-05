package log

// type logger struct {
// 	log     *zap.SugaredLogger
// 	options *logOptions
// }

// func (l *logger) Debug(args ...interface{}) {
// 	l.log.Debug(args)
// }

// func (l *logger) Info(message string, fields ...Field) {
// 	l.log.Info(message, fields...)
// }

// func (l *logger) Warn(message string, fields ...Field) {
// 	l.log.Warn(message, fields...)
// }

// func (l *logger) Error(message string, fields ...Field) {
// 	l.log.Error(message, fields...)
// }

// func (l *logger) Panic(message string, fields ...Field) {
// 	l.log.Panic(message, fields...)
// }

// func (l *logger) Fatal(message string, fields ...Field) {
// 	l.log.Fatal(message, fields...)
// }

// func (l *logger) Sync() error {
// 	return l.log.Sync()
// }

// func (l *logger) Named(name string) *fieldLogger {
// 	return &fieldLogger{
// 		log: l.log.Named(name),
// 	}
// }

// func (l *logger) NewLog(opt ...Option) *fieldLogger {
// 	opts := l.options.copy()
// 	opts.apply(opt...)
// 	return initLogger(opts)
// }

// // New create a new field logger
// func NewLog(opt ...Option) *fieldLogger {
// 	opts := &logOptions{}
// 	opts.setDefault()
// 	opts.apply(opt...)
// 	return initLogger(opts)
// }

// func initSimpleLogger(opts *logOptions) *fieldLogger {
// 	config := zap.NewProductionEncoderConfig()

// 	var zapOptions []zap.Option
// 	var outputEncoder zapcore.Encoder

// 	if opts.funcPath {
// 		config.FunctionKey = "Function"
// 	}

// 	if opts.filePath {
// 		config.CallerKey = "Caller"
// 		zapOptions = append(zapOptions, zap.AddCaller())
// 	}

// 	if opts.skipFrameCount != 0 {
// 		zapOptions = append(zapOptions, zap.AddCallerSkip(opts.skipFrameCount))
// 	}

// 	config.EncodeLevel = opts.levelEncoder
// 	config.EncodeTime = opts.timeEncoder

// 	switch opts.output {
// 	case JSONFormat:
// 		config.EncodeLevel = CapitalLevelEncoder
// 		outputEncoder = zapcore.NewJSONEncoder(config)

// 	case TextFormat:
// 		outputEncoder = zapcore.NewConsoleEncoder(config)
// 	}

// 	zapOptions = append(zapOptions, zap.AddStacktrace(PANIC))

// 	// fileEncoder := zapcore.NewConsoleEncoder(config)
// 	// consoleEncoder := zapcore.NewConsoleEncoder(config)
// 	// logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	// writer := zapcore.AddSync(logFile)

// 	core := zapcore.NewTee(
// 		// zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
// 		zapcore.NewCore(outputEncoder, zapcore.AddSync(os.Stderr), opts.logLevel),
// 	)

// 	// core := zapcore.NewCore(outputEncoder, zapcore.AddSync(os.Stdout), opts.logLevel)
// 	var log *zap.Logger
// 	log = zap.New(core, zapOptions...)

// 	if opts.name != "" {
// 		log = log.Named(opts.name)
// 	}
// 	return &fieldLogger{
// 		log:     log,
// 		options: opts,
// 	}
// }
