package zlog

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
    Log.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
    Log.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
    Log.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
    Log.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
    Log.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
    Log.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
    Log.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
    Log.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
    Log.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
    Log.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
    Log.Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
    Log.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
    Log.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
    Log.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
    Log.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
    Log.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
    Log.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
    Log.Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
    Log.DPanicw(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
    Log.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
    Log.Fatalw(msg, keysAndValues)
}

// Debugln uses fmt.Sprintln to construct and log a message.
func Debugln(args ...interface{}) {
    Log.Debugln(args)
}

// Infoln uses fmt.Sprintln to construct and log a message.
func Infoln(args ...interface{}) {
    Log.Infoln(args)
}

// Warnln uses fmt.Sprintln to construct and log a message.
func Warnln(args ...interface{}) {
    Log.Warnln(args)
}

// Errorln uses fmt.Sprintln to construct and log a message.
func Errorln(args ...interface{}) {
    Log.Errorln(args)
}

// DPanicln uses fmt.Sprintln to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicln(args ...interface{}) {
    Log.DPanicln(args)
}

// Panicln uses fmt.Sprintln to construct and log a message, then panics.
func Panicln(args ...interface{}) {
    Log.Panicln(args)
}

// Fatalln uses fmt.Sprintln to construct and log a message, then calls os.Exit.
func Fatalln(args ...interface{}) {
    Log.Fatalln(args)
}
