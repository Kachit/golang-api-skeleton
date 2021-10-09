package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Models_Repositories_UsersRepository_Count(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1)

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetUsersRepository()
	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users").WillReturnRows(rows)
	result, err := repository.Count(nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func Test_Models_Repositories_UsersRepository_GetByID(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetUsersRepository()
	mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)
	user, err := repository.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, "foo", user.Name)
	assert.Equal(t, "bar", user.Email)
}

func Test_Models_Repositories_UsersRepository_GetByCode(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetUsersRepository()
	mock.ExpectQuery("SELECT \\* FROM users WHERE code = \\$1").WithArgs("foo").WillReturnRows(rows)
	user, err := repository.GetByEmail("foo")
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, "foo", user.Name)
	assert.Equal(t, "bar", user.Email)
}

func Test_Models_Repositories_UsersRepository_GetList(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetUsersRepository()
	mock.ExpectQuery("SELECT \\* FROM users").WillReturnRows(rows)
	collection, err := repository.GetList(nil, 0, 0, nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), (*collection)[0].Id)
	assert.Equal(t, "foo", (*collection)[0].Name)
	assert.Equal(t, "bar", (*collection)[0].Email)
}

//func Test_Models_Repositories_UsersRepository_Insert(t *testing.T) {
//	sqlxDB, mock := testable.GetDatabaseMock()
//
//	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
//
//	factory := NewRepositoriesFactory(sqlxDB)
//	repository := factory.GetUsersRepository()
//	promocode := &User{}
//	promocode.Status = PROMOCODE_STATUS_ACTIVE
//
//	mock.ExpectQuery("INSERT INTO promocodes \\(code,created_at,discount,limits,link,locale,options,status,subscription_id,title,usage_limits,usage_type\\) VALUES \\(\\$1,\\$2,\\$3,\\$4,\\$5,\\$6,\\$7,\\$8,\\$9,\\$10,\\$11,\\$12\\) RETURNING id").WillReturnRows(rows)
//	_, err := repository.Insert(promocode)
//	assert.NoError(t, err)
//	assert.Equal(t, uint64(1), promocode.Id)
//}
//
//func Test_Models_Repositories_UsersRepository_Update(t *testing.T) {
//	sqlxDB, mock := testable.GetDatabaseMock()
//
//	factory := NewRepositoriesFactory(sqlxDB)
//	repository := factory.GetUsersRepository()
//	mock.ExpectExec("UPDATE promocodes SET code = \\$1, discount = \\$2, limits = \\$3, link = \\$4, locale = \\$5, options = \\$6, status = \\$7, subscription_id = \\$8, title = \\$9, usage_limits = \\$10, usage_type = \\$11 WHERE id = \\$12").WillReturnResult(sqlmock.NewResult(0, 1))
//	promocode := &User{}
//	promocode.Id = 1
//	promocode.Status = PROMOCODE_STATUS_ACTIVE
//	_, err := repository.Update(promocode)
//	assert.NoError(t, err)
//}
