package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type loggerInterface interface {
  Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	logger *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{GetOutput()},
		Level:       zap.NewAtomicLevelAt(GetLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log.logger, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetOutput() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(envLogOutput)))
	if output == "" {
		return "stdout"
	}
	return output
}

func GetLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func GetLogger() loggerInterface {
	return log
}

func (log logger) Print(v ...interface{}) {
  Info(fmt.Sprintf("%v", v))
}

func (log logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func Info(msg string, tags ...zap.Field) {
	log.logger.Info(msg, tags...)
	log.logger.Sync()

}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	log.logger.Error(msg, tags...)
	log.logger.Sync()
}
