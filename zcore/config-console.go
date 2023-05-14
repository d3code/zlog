package zcore

import (
    "github.com/d3code/clog/pkg"
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

    if entry.Level == zapcore.DebugLevel {
        entry.Message = pkg.ColorString(entry.Message, "grey")
    }

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

func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    format := t.Format(time.RFC3339)
    greyTime := pkg.ColorString(format, "grey")
    enc.AppendString(greyTime)
}

func EncodeCallerColor(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
    path := caller.TrimmedPath()
    greyPath := pkg.ColorString(path, "grey")
    enc.AppendString(greyPath)
}

func EncodeLevelColor() zapcore.LevelEncoder {
    debug := pkg.ColorString("DEBUG", "grey")
    info := pkg.ColorString("INFO", "blue")
    warning := pkg.ColorString("WARNING", "yellow")
    errorLevel := pkg.ColorString("ERROR", "red")
    critical := pkg.ColorString("CRITICAL", "red")
    alert := pkg.ColorString("ALERT", "red")
    emergency := pkg.ColorString("EMERGENCY", "red")

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
