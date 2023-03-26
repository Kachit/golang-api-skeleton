package infrastructure

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDatabase(config *Config) (*gorm.DB, error) {
	//options
	logMode := logger.Silent
	//debug options override
	if config.App.Debug {
		logMode = logger.Info
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.GetDatabaseDsn(),
		//PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logMode),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.Database.MaxConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute)
	return db, nil
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
