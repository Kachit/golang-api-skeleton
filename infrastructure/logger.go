package infrastructure

import (
	"encoding/json"
	"github.com/int128/slack"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/lajosbencz/glo"
	"strings"
	"time"
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

// NewFormatter creates a Formatter from a string
func NewFormatter(f string) glo.Formatter {
	return &MattermostFormatter{f}
}

type MattermostFormatter struct {
	format string
}

func (f *MattermostFormatter) Format(time time.Time, level glo.Level, line string, params ...interface{}) string {
	m := ""
	if len(params) > 0 {
		b, err := json.Marshal(params)
		if err == nil {
			m = string(b)
		}
	}

	r := strings.NewReplacer(
		"{T}", time.Format("2006-01-02T15:04:05Z07:00"),
		"{L}", level.String(),
		"{M}", line,
		"{P}", m,
	)
	return r.Replace(f.format)
}

func NewLogger(cfg *config.Config) Logger {
	adapterGlo := NewLoggerAdapterGlo(&cfg.Logger)
	log := &LoggerAdapterGlo{logger: adapterGlo}
	return log
}

func NewLoggerAdapterGlo(cfg *config.LoggerConfig) glo.Facility {
	mh := NewMattermostHandler(&cfg.Mattermost)
	log := glo.NewStdFacility()
	log.PushHandler(mh)
	return log
}

func NewMattermostHandler(cfg *config.LoggerAdapterMattermostConfig) *MattermostHandler {
	c := &MattermostHandlerConfig{Username: cfg.Username}
	client := newMattermostClient(cfg)
	filter := glo.NewFilterLevel(glo.Warning)
	formatter := NewFormatter("{T} [{L}] {M} {P}")
	h := &MattermostHandler{config: c, writer: client, formatter: formatter}
	h.PushFilter(filter)
	return h
}

func newMattermostClient(cfg *config.LoggerAdapterMattermostConfig) *slack.Client {
	client := &slack.Client{WebhookURL: cfg.WebhookUrl}
	return client
}

type MattermostHandlerConfig struct {
	Username string
}

type MattermostHandler struct {
	config    *MattermostHandlerConfig
	writer    *slack.Client
	formatter glo.Formatter
	filters   []glo.Filter
}

// Log logs a line with a specific level
func (h *MattermostHandler) Log(level glo.Level, line string, params ...interface{}) error {
	valid := true
	for _, f := range h.filters {
		if !f.Check(level, line, params) {
			valid = false
			break
		}
	}
	if !valid {
		return nil
	}
	l := h.formatter.Format(time.Now(), level, line, params...) + "\n"
	msg := &slack.Message{
		Username: h.config.Username,
		Text:     l,
		//Attachments: att,
	}
	err := h.writer.Send(msg)
	return err
}

func (h *MattermostHandler) SetFormatter(formatter glo.Formatter) glo.Handler {
	h.formatter = formatter
	return h
}

func (h *MattermostHandler) ClearFilters() glo.Handler {
	h.filters = []glo.Filter{}
	return h
}

func (h *MattermostHandler) PushFilter(filter glo.Filter) glo.Handler {
	h.filters = append(h.filters, filter)
	return h
}
