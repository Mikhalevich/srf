package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

type Logger struct {
	l *logrus.Entry
}

func New() *Logger {
	l := logrus.New()
	l.SetFormatter(&ecslogrus.Formatter{})

	// return &Logger{
	// 	l: l.WithField("service_name", name),
	// }
	return &Logger{
		l: logrus.NewEntry(l),
	}
}

func (lw *Logger) Debugf(format string, args ...interface{}) {
	lw.l.Debugf(format, args...)
}

func (lw *Logger) Infof(format string, args ...interface{}) {
	lw.l.Infof(format, args...)
}

func (lw *Logger) Warnf(format string, args ...interface{}) {
	lw.l.Warnf(format, args...)
}

func (lw *Logger) Errorf(format string, args ...interface{}) {
	lw.l.Errorf(format, args...)
}

func (lw *Logger) Debug(args ...interface{}) {
	lw.l.Debug(args...)
}

func (lw *Logger) Info(args ...interface{}) {
	lw.l.Info(args...)
}

func (lw *Logger) Warn(args ...interface{}) {
	lw.l.Warn(args...)
}

func (lw *Logger) Error(args ...interface{}) {
	lw.l.Error(args...)
}

func (lw *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{
		l: lw.l.WithContext(ctx),
	}
}

func (lw *Logger) WithError(err error) *Logger {
	return &Logger{
		l: lw.l.WithError(err),
	}
}

func (lw *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		l: lw.l.WithField(key, value),
	}
}

func (lw *Logger) WithFields(fields map[string]interface{}) *Logger {
	return &Logger{
		l: lw.l.WithFields(fields),
	}
}
