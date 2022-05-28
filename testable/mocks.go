package testable

import (
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/lajosbencz/glo"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type LoggerMock struct {
	Level  glo.Level
	Msg    string
	Params interface{}
}

func (l *LoggerMock) write(level glo.Level, msg string, params ...interface{}) {
	l.Level = level
	l.Msg = msg
	l.Params = params
}

func (l *LoggerMock) Debug(msg string, params ...interface{}) {
	l.write(glo.Debug, msg, params)
}

func (l *LoggerMock) Info(msg string, params ...interface{}) {
	l.write(glo.Info, msg, params)
}

func (l *LoggerMock) Notice(msg string, params ...interface{}) {
	l.write(glo.Notice, msg, params)
}

func (l *LoggerMock) Warning(msg string, params ...interface{}) {
	l.write(glo.Warning, msg, params)
}

func (l *LoggerMock) Error(msg string, params ...interface{}) {
	l.write(glo.Error, msg, params)
}

func (l *LoggerMock) Critical(msg string, params ...interface{}) {
	l.write(glo.Critical, msg, params)
}

func (l *LoggerMock) Alert(msg string, params ...interface{}) {
	l.write(glo.Alert, msg, params)
}

func (l *LoggerMock) Emergency(msg string, params ...interface{}) {
	l.write(glo.Emergency, msg, params)
}

func GetDatabaseMock() (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	logMode := logger.Silent

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db",
		DriverName:           "postgres",
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	})

	db, _ := gorm.Open(dialector, &gorm.Config{Logger: logger.Default.LogMode(logMode)})
	return db, mock
}

func NewConfigMock() (*config.Config, error) {
	return config.NewConfig("../config.yml")
}

func GetLoggerMock() *LoggerMock {
	return &LoggerMock{}
}
