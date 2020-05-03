package utils

import (

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/RainrainWu/probe/config"
)

var (
	Logger *zap.Logger
	hook lumberjack.Logger = lumberjack.Logger{
        Filename:   config.LOG_FILE,
        MaxSize:    128,
        MaxBackups: 30, 
        MaxAge:     7,
        Compress:   true,
	}
	encoderConfig zapcore.EncoderConfig = zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "linenum",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,
        EncodeName:     zapcore.FullNameEncoder,
    }
)

func init() {

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
        atomicLevel,
    )

	Logger = zap.New(core)
	defer Logger.Sync()
}