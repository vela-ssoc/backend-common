package logback

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志接口定义
type Logger interface {
	Trace(...any)
	Debug(...any)
	Info(...any)
	Warn(...any)
	Error(...any)
	Tracef(string, ...any)
	Debugf(string, ...any)
	Infof(string, ...any)
	Warnf(string, ...any)
	Errorf(string, ...any)
	Level() zapcore.Level
}

func Sugar(l *zap.Logger) Logger {
	return &sugarLog{sugar: l.Sugar()}
}

type sugarLog struct {
	sugar *zap.SugaredLogger
}

func (sg *sugarLog) Trace(v ...any)            { sg.sugar.Debug(v...) }
func (sg *sugarLog) Debug(v ...any)            { sg.sugar.Debug(v...) }
func (sg *sugarLog) Info(v ...any)             { sg.sugar.Info(v...) }
func (sg *sugarLog) Warn(v ...any)             { sg.sugar.Warn(v...) }
func (sg *sugarLog) Error(v ...any)            { sg.sugar.Error(v...) }
func (sg *sugarLog) Tracef(s string, v ...any) { sg.sugar.Debugf(s, v...) }
func (sg *sugarLog) Debugf(s string, v ...any) { sg.sugar.Debugf(s, v...) }
func (sg *sugarLog) Infof(s string, v ...any)  { sg.sugar.Infof(s, v...) }
func (sg *sugarLog) Warnf(s string, v ...any)  { sg.sugar.Warnf(s, v...) }
func (sg *sugarLog) Errorf(s string, v ...any) { sg.sugar.Errorf(s, v...) }
func (sg *sugarLog) Level() zapcore.Level      { return sg.sugar.Level() }
