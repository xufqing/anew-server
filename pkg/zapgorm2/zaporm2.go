package zapgorm2

import (
	"anew-server/pkg/common"
	"context"
	"fmt"
	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type Logger struct {
	ZapLogger     *zap.SugaredLogger
	LogLevel      gormlogger.LogLevel
	SlowThreshold time.Duration
}

func New(zapLogger *zap.SugaredLogger) Logger {
	logLevel := gormlogger.Info
	if common.Conf.Logs.Level == 1 {
		logLevel = gormlogger.Warn
	} else if common.Conf.Logs.Level >= 2 {
		logLevel = gormlogger.Error
	}
	return Logger{
		ZapLogger:     zapLogger,
		LogLevel:      logLevel,
		SlowThreshold: 100 * time.Millisecond,
	}
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
		LogLevel:      level,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.ZapLogger.Infof(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.ZapLogger.Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.ZapLogger.Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		sql, rows := fc()
		l.ZapLogger.Error(fmt.Sprintf("[rows]:%d\t[sql]:%s", rows, sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		l.ZapLogger.Warn(fmt.Sprintf("[rows]:%d\t[sql]:%s", rows, sql))
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.ZapLogger.Debug(fmt.Sprintf("[rows]:%d\t[sql]:%s", rows, sql))
	}
}
