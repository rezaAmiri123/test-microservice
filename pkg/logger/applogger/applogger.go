package applogger

import (
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func NewLoggerConfig(logLevel string, devMode bool, encoder string) *Config {
	return &Config{LogLevel: logLevel, DevMode: devMode, Encoder: encoder}
}

// Application logger
type appLogger struct {
	level       string
	devMode     bool
	encoding    string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

// NewAppLogger App Logger constructor
func NewAppLogger(cfg Config) logger.Logger {
	return &appLogger{level: cfg.LogLevel, devMode: cfg.DevMode, encoding: cfg.Encoder}
}


// For mapping config logger to email_service logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *appLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *appLogger) InitLogger() {
	logLevel := l.getLoggerLevel()

	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if l.devMode {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.NameKey = "[SERVICE]"
	encoderCfg.TimeKey = "[TIME]"
	encoderCfg.LevelKey = "[LEVEL]"
	encoderCfg.FunctionKey = "[CALLER]"
	encoderCfg.CallerKey = "[LINE]"
	encoderCfg.MessageKey = "[MESSAGE]"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeName = zapcore.FullNameEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	if l.encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	mewLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = mewLogger
	l.sugarLogger = mewLogger.Sugar()
}

// Logger methods

// WithName add logger microservice name
func (l *appLogger) WithName(name string) {
	l.logger = l.logger.Named(name)
	l.sugarLogger = l.sugarLogger.Named(name)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *appLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *appLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *appLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *appLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Printf uses fmt.Sprintf to log a templated message
func (l *appLogger) Printf(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *appLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// WarnMsg log error message with warn level.
func (l *appLogger) WarnMsg(msg string, err error) {
	l.logger.Warn(msg, zap.String("error", err.Error()))
}
// Warnf uses fmt.Sprintf to log a templated message.
func (l *appLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *appLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *appLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// Err uses error to log a message.
func (l *appLogger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *appLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *appLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *appLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (l *appLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *appLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Sync flushes any buffered log entries
func (l *appLogger) Sync() error {
	go l.logger.Sync() // nolint: errcheck
	return l.sugarLogger.Sync()
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *appLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

func (l *appLogger)KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, time time.Time){
	l.logger.Debug(
		"Processing Kafka message",
		zap.String(logger.Topic, topic),
		zap.Int(logger.Partition, partition),
		zap.String(logger.Message, message),
		zap.Int(logger.WorkerID, workerID),
		zap.Int64(logger.Offset, offset),
		zap.Time(logger.Time, time),
	)
}
func (l *appLogger)KafkaLogCommittedMessage(topic string, partition int, offset int64){
	l.logger.Info(
		"Committed Kafka message",
		zap.String(logger.Topic, topic),
		zap.Int(logger.Partition, partition),
		zap.Int64(logger.Offset, offset),
	)
}