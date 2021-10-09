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
	mock.ExpectQuery("SELECT \\* FROM users WHERE email = \\$1").WithArgs("foo").WillReturnRows(rows)
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

func Test_Models_Repositories_UsersRepository_Insert(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetUsersRepository()
	user := &User{}
	user.Email = "foo@bar.baz"

	mock.ExpectQuery("INSERT INTO users \\(created_at,description,email,name,password\\) VALUES \\(\\$1,\\$2,\\$3,\\$4,\\$5\\) RETURNING id").WillReturnRows(rows)
	_, err := repository.Insert(user)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
}

func Test_Models_Repositories_UsersRepository_Update(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetUsersRepository()
	mock.ExpectExec("UPDATE users SET created_at = \\$1, description = \\$2, email = \\$3, name = \\$4, password = \\$5 WHERE id = \\$6").WillReturnResult(sqlmock.NewResult(0, 1))
	user := &User{}
	user.Id = 1
	user.Email = "foo@bar.baz"
	_, err := repository.Update(user)
	assert.NoError(t, err)
}
