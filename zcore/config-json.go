package zcore

import (
    "github.com/d3code/clog/color"
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

    if entry.Level == zapcore.DebugLevel {
        entry.Message = color.String(entry.Message, "grey")
    }

    fields = append(fields, zap.String("environment", os.Getenv("environment")))
    entry.Time = entry.Time.Local()

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
