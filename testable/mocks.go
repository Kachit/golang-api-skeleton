package testable

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func GetDatabaseMock() (*sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	//defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	return sqlxDB, mock
}
