package logger

import "go.uber.org/zap"

type Option = zap.Option

var (
	WithCaller    = zap.WithCaller
	AddCallerSkip = zap.AddCallerSkip
	AddStacktrace = zap.AddStacktrace
)

type ConfigMode string

const (
	DevMode  ConfigMode = "DEV"
	ProdMode ConfigMode = "PROD"
)

type OutputEncoder string

func (cm ConfigMode) Valid() bool {
	switch cm {
	case DevMode, ProdMode:
		return true
	default:
		return false
	}
}

const (
	Console OutputEncoder = "CONSOLE"
	JSON    OutputEncoder = "JSON"
)

func (oe OutputEncoder) Valid() bool {
	switch oe {
	case Console, JSON:
		return true
	default:
		return false
	}
}
