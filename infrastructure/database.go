package infrastructure

import (
	"github.com/kachit/golang-api-skeleton/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDatabase(config *config.Config) (*gorm.DB, error) {
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
		Logger: logger.Default.LogMode(logMode),
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
