package logger

import (
	"go.uber.org/zap"
)

type LoggerService struct {
	logger      *zap.Logger
	loggerLevel *zap.AtomicLevel
}

var LogService *LoggerService

func init() {
	logger, loggerLevel, err := NewLogger()
	if err != nil {
		panic(err)
	}

	LogService = &LoggerService{
		logger:      logger,
		loggerLevel: loggerLevel,
	}
}

func NewLogger() (*zap.Logger, *zap.AtomicLevel, error) {
	return NewLoggerWithConfig(zap.NewProductionConfig())
}

func NewLoggerWithConfig(config zap.Config) (*zap.Logger, *zap.AtomicLevel, error) {
	logger, err := config.Build()
	if err != nil {
		return nil, nil, err
	}

	return logger, &config.Level, nil
}

func (l *LoggerService) Info(args ...interface{}) {
	l.logger.Sugar().Info(args...)
}

func (l *LoggerService) RawInfo(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *LoggerService) Debug(args ...interface{}) {
	l.logger.Sugar().Debug(args...)
}

func (l *LoggerService) RawDebug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *LoggerService) Warn(args ...interface{}) {
	l.logger.Sugar().Warn(args...)
}

func (l *LoggerService) RawWarn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *LoggerService) Error(args ...interface{}) {
	l.logger.Sugar().Error(args...)
}

func (l *LoggerService) RawError(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *LoggerService) Fatal(args ...interface{}) {
	l.logger.Sugar().Fatal(args...)
}

func (l *LoggerService) RawFatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *LoggerService) Panic(args ...interface{}) {
	l.logger.Sugar().Panic(args...)
}

func (l *LoggerService) RawPanic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *LoggerService) Sync() error {
	return l.logger.Sync()
}

func (l *LoggerService) With(fields ...zap.Field) *zap.Logger {
	return l.logger.With(fields...)
}

func (l *LoggerService) Named(name string) *zap.Logger {
	return l.logger.Named(name)
}

func (l *LoggerService) SetLevel(level zap.AtomicLevel) {
	l.loggerLevel = &level
}

func (l *LoggerService) GetLevel() zap.AtomicLevel {
	return *l.loggerLevel
}
