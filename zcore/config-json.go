package zcore

import (
    "go.uber.org/zap"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
    "os"
)

type JsonEncoder struct {
    zapcore.Encoder
    Pool buffer.Pool
}

func (e *JsonEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
    buf := e.Pool.Get()

    if env, found := os.LookupEnv("ENVIRONMENT"); found {
        fields = append(fields, zap.String("environment", env))
    }

    consoleBuffer, err := e.Encoder.EncodeEntry(entry, fields)
    if err != nil {
        return nil, err
    }

    _, err = buf.Write(consoleBuffer.Bytes())
    if err != nil {
        return nil, err
    }

    return buf, nil
}

func EncodeCallerFull(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
    path := caller.FullPath()
    enc.AppendString(path)
}
