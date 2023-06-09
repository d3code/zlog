package zlog

import (
    "github.com/d3code/zlog/zcore"
    "go.uber.org/zap"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
    "os"
)

var Log *zap.SugaredLogger

func init() {
    config := zcore.Config
    if env, found := os.LookupEnv("ENVIRONMENT"); !found || env != "production" {
        config.EncodeTime = zcore.EncodeTime
        config.EncodeCaller = zcore.EncodeCallerColor
        config.EncodeLevel = zcore.EncodeLevelColor()

        encoder := &zcore.ConsoleEncoder{
            Encoder: zapcore.NewConsoleEncoder(config),
            Pool:    buffer.NewPool(),
        }
        Log = zap.New(
            zapcore.NewCore(
                encoder,
                os.Stdout,
                zapcore.DebugLevel,
            ),
            zap.ErrorOutput(os.Stderr),
            zap.AddStacktrace(zapcore.FatalLevel),
            zap.AddCaller(),
        ).Sugar()
    } else {
        config.EncodeCaller = zcore.EncodeCallerFull

        encoder := &zcore.JsonEncoder{
            Encoder: zapcore.NewJSONEncoder(config),
            Pool:    buffer.NewPool(),
        }
        Log = zap.New(
            zapcore.NewCore(
                encoder,
                os.Stdout,
                zapcore.DebugLevel,
            ),
            zap.ErrorOutput(os.Stderr),
            zap.AddStacktrace(zapcore.FatalLevel),
            zap.AddCaller(),
        ).Sugar()
    }
}
