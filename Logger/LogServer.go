package Logger

import (
	"go.uber.org/zap"
)
//固定字段
type CommFields struct {
	Uid int
	TraceID string
	UserAgent string
	Url string
	Params string
}

func Debug(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Debug(msg, parseParams(commFields, fields)...)
}

func Info(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Info(msg, parseParams(commFields, fields)...)
}

func Error(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Warn(msg, parseParams(commFields, fields)...)
}

func Fatal(msg string, commFields CommFields, fields ...zap.Field) {
	zap.L().Fatal(msg, parseParams(commFields, fields)...)
}


func parseParams(commFields CommFields, fields []zap.Field) []zap.Field {
	s := []zap.Field{
		zap.Int("uid", commFields.Uid),
		zap.String("trace_id", commFields.TraceID),
		zap.String("user_agent", commFields.UserAgent),
		zap.String("url", commFields.Url),
		zap.String("params", commFields.Params),
		//以上为固定字段
		zap.Namespace("data"),
	}
	return append(s, fields...)
}
