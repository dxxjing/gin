package Logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"runtime"
)

func InitZapLogger() {
	hook := getLoggerWriter()
	p, _ := os.Getwd()
	fmt.Println("test-dir")
	fmt.Println(path.Dir(p))
	encoderConfig := encoderConfig()
	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置 json格式
		//zapcore.NewConsoleEncoder(encoderConfig),		//普通文本格式
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.AddSync(&hook),	//打印到文件
		atomicLevel,                                                                     // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	/*filed := zap.Fields(
		zap.String("serviceName", "hippo-server"),
		//zap.Int("uid", 3688),
	)*/

	// 构造日志
	logger := zap.New(core, caller, development)
	//替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	//多个进程间 互不影响  印证了 进程间共享全局变量(仅限于有血缘关系的进程间)，但全局变量的值在进程间是不共享的，
	//因为生成子进程 会将父进程的代码拷贝一份（全局变量在每个子进程中都是初始值）
	zap.ReplaceGlobals(logger)
}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		//EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			//自定义
			encoder.AppendString(zapcore.NewEntryCaller(runtime.Caller(6)).String())
		},// 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
		ConsoleSeparator: "|",
	}
}

func getLoggerWriter() lumberjack.Logger {
	return lumberjack.Logger{
		Filename:   "./logs/server.log", // 日志文件路径
		MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}
}