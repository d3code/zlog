package zcore

import (
    "github.com/d3code/clog/color"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
    "time"
)

type ConsoleEncoder struct {
    zapcore.Encoder
    Pool buffer.Pool
}

func (e *ConsoleEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
    buf := e.Pool.Get()

    switch level := entry.Level; {
    case level == zapcore.DebugLevel:
        entry.Message = color.String(entry.Message, "grey")
    case level == zapcore.WarnLevel:
        entry.Message = color.String(entry.Message, "yellow")
    case level >= zapcore.ErrorLevel:
        entry.Message = color.String(entry.Message, "red")
    }

    //entry.Time = entry.Time.Local()

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

func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    format := t.Format(time.RFC3339)
    greyTime := color.String(format, "grey")
    enc.AppendString(greyTime)
}

func EncodeCallerColor(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
    path := caller.TrimmedPath()
    greyPath := color.String(path, "grey")
    enc.AppendString(greyPath)
}

func EncodeLevelColor() zapcore.LevelEncoder {
    debug := color.String("DEBUG", "grey")
    info := color.String("INFO", "blue")
    warning := color.String("WARNING", "yellow")
    errorLevel := color.String("ERROR", "red")
    critical := color.String("CRITICAL", "red")
    alert := color.String("ALERT", "red")
    emergency := color.String("EMERGENCY", "red")

    return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {

        switch l {
        case zapcore.DebugLevel:
            enc.AppendString(debug)
        case zapcore.InfoLevel:
            enc.AppendString(info)
        case zapcore.WarnLevel:
            enc.AppendString(warning)
        case zapcore.ErrorLevel:
            enc.AppendString(errorLevel)
        case zapcore.DPanicLevel:
            enc.AppendString(critical)
        case zapcore.PanicLevel:
            enc.AppendString(alert)
        case zapcore.FatalLevel:
            enc.AppendString(emergency)
        }
    }
}
