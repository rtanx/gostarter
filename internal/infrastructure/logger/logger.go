package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel
	WarnLevel   Level = zap.WarnLevel
	ErrorLevel  Level = zap.ErrorLevel
	DPaniclevel Level = zap.DPanicLevel
	PanicLevel  Level = zap.PanicLevel
	FatalLevel  Level = zap.FatalLevel
	DebugLevel  Level = zap.DebugLevel
)

type Logger struct {
	l     *zap.Logger
	level Level
}

var std *Logger = New(os.Stderr, InfoLevel, ProdMode, Console, WithCaller(true), AddCallerSkip(1), AddStacktrace(DPaniclevel))

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}
func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func New(writer io.Writer, level Level, confMode ConfigMode, outputEncoder OutputEncoder, opts ...Option) *Logger {
	if writer == nil {
		panic("nil writter given")
	}
	var (
		conf    zap.Config
		encoder zapcore.Encoder
	)
	if !confMode.Valid() {
		panic("invalid config mode")
	}
	switch confMode {
	case DevMode:
		conf = zap.NewDevelopmentConfig()
	case ProdMode:
		fallthrough
	default:
		conf = zap.NewProductionConfig()
	}
	conf.EncoderConfig.ConsoleSeparator = " | "
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	conf.EncoderConfig.EncodeTime = timeEncoder
	if !outputEncoder.Valid() {
		panic("invalid output encode")
	}
	switch outputEncoder {
	case JSON:
		encoder = zapcore.NewJSONEncoder(conf.EncoderConfig)
	case Console:
		fallthrough
	default:
		encoder = zapcore.NewConsoleEncoder(conf.EncoderConfig)
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	return &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
}

func Default() *Logger {
	return std
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006/01/02 - 15:04:05 MST")))
}
