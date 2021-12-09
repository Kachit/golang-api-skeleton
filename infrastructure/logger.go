package infrastructure

import (
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/lajosbencz/glo"
)

type Logger interface {
	// Logger logs an debug line
	Debug(string, ...interface{})
	// Logger logs an info line
	Info(string, ...interface{})
	// Logger logs an info line
	Notice(string, ...interface{})
	// Logger logs a warning line
	Warning(string, ...interface{})
	// Logger logs an error line
	Error(string, ...interface{})
	// Logger logs an critical line
	Critical(string, ...interface{})
	// Logger logs an alert line
	Alert(string, ...interface{})
	// Logger logs an emergency line
	Emergency(string, ...interface{})
}

type LoggerAdapterGlo struct {
	logger glo.Facility
}

func (l *LoggerAdapterGlo) Debug(msg string, params ...interface{}) {
	_ = l.logger.Debug(msg, params)
}

func (l *LoggerAdapterGlo) Info(msg string, params ...interface{}) {
	_ = l.logger.Info(msg, params)
}

func (l *LoggerAdapterGlo) Notice(msg string, params ...interface{}) {
	_ = l.logger.Notice(msg, params)
}

func (l *LoggerAdapterGlo) Warning(msg string, params ...interface{}) {
	_ = l.logger.Warning(msg, params)
}

func (l *LoggerAdapterGlo) Error(msg string, params ...interface{}) {
	_ = l.logger.Error(msg, params)
}

func (l *LoggerAdapterGlo) Critical(msg string, params ...interface{}) {
	_ = l.logger.Critical(msg, params)
}

func (l *LoggerAdapterGlo) Alert(msg string, params ...interface{}) {
	_ = l.logger.Alert(msg, params)
}

func (l *LoggerAdapterGlo) Emergency(msg string, params ...interface{}) {
	_ = l.logger.Emergency(msg, params)
}

func NewLogger(cfg *config.Config) Logger {
	adapterGlo := NewLoggerAdapterGlo()
	log := &LoggerAdapterGlo{logger: adapterGlo}
	return log
}

func NewLoggerAdapterGlo() glo.Facility {
	log := glo.NewStdFacility()
	return log
}
