package Logger

import (
	"go.uber.org/zap"
	"time"
)

//固定字段
type CommFields struct {
	Uid int
	TraceID string
	UserAgent string
	Url string
	Params string
	Cost time.Duration
	ClientIP string
}

var defaultFields []zap.Field


func Debug(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Debug(msg, parseParams(commFields, fields...)...)
}

func Info(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Info(msg, parseParams(commFields, fields...)...)
}

func Warn(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Warn(msg, parseParams(commFields, fields...)...)
}

func Error(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Error(msg, parseParams(commFields, fields...)...)
}

func Fatal(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Fatal(msg, parseParams(commFields, fields...)...)
}


//仅用于打印msg
func DebugMsg(msg string) {
	zap.L().Debug(msg, parseParams(CommFields{}, defaultFields...)...)
}

func InfoMsg(msg string) {
	zap.L().Info(msg, parseParams(CommFields{}, defaultFields...)...)
}

func WarnMsg(msg string) {
	zap.L().Warn(msg, parseParams(CommFields{}, defaultFields...)...)
}

func ErrorMsg(msg string) {
	zap.L().Error(msg, parseParams(CommFields{}, defaultFields...)...)
}

func FatalMsg(msg string) {
	zap.L().Fatal(msg, parseParams(CommFields{}, defaultFields...)...)
}


func parseParams(commFields CommFields, fields ...zap.Field) []zap.Field {
	s := []zap.Field{
		zap.Int("uid", commFields.Uid),
		zap.String("trace_id", commFields.TraceID),
		zap.String("user_agent", commFields.UserAgent),
		zap.String("url", commFields.Url),
		zap.String("params", commFields.Params),
		zap.Int64("cost", int64(commFields.Cost)),
		zap.String("client_ip", commFields.ClientIP),
		//以上为固定字段
		zap.Namespace("data"),
	}
	return append(s, fields...)
}