package log

import "github.com/go-zhouxun/xlog"

var logger xlog.XLog

func InitLog(_logger xlog.XLog) {
	logger = _logger
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

func Crit(msg string, args ...interface{}) {
	logger.Crit(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}
