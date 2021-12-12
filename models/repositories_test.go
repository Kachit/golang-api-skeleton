package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Models_Repositories_RepositoriesFactory_GetRepositoryAbstract(t *testing.T) {
	factory := NewRepositoriesFactory(nil)
	result := factory.GetRepositoryAbstract(TableUsers)
	assert.IsType(t, (*RepositoryAbstract)(nil), result)
	assert.Equal(t, TableUsers, result.table)
}

func Test_Models_Repositories_RepositoriesFactory_GetUsersRepository(t *testing.T) {
	factory := NewRepositoriesFactory(nil)
	result := factory.GetUsersRepository()
	assert.IsType(t, (*UsersRepository)(nil), result)
	assert.Equal(t, TableUsers, result.table)
}

func Test_Models_Repositories_RepositoryAbstract_Insert(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("INSERT INTO users \\(foo\\) VALUES \\(\\$1\\) RETURNING id").WillReturnRows(rows)
	row := map[string]interface{}{
		"foo": "bar",
	}
	_, err := repository.insert(row)
	assert.NoError(t, err)
}

func Test_Models_Repositories_RepositoryAbstract_Update(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectExec("UPDATE users SET foo = \\$1 WHERE").WillReturnResult(sqlmock.NewResult(0, 1))
	row := map[string]interface{}{
		"foo": "bar",
	}
	result, err := repository.update(row, nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, int(result))
}

func Test_Models_Repositories_RepositoryAbstract_FetchOne(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)
	user := &User{}
	err := repository.fetchOne(user, sq.Eq{"id": 1}, nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, "foo", user.Name)
	assert.Equal(t, "bar", user.Email)
}

func Test_Models_Repositories_RepositoryAbstract_FetchOneWithOrderBy(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")
	orderBy := map[string]string{"id": "ASC"}

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1 ORDER BY id ASC").WithArgs(1).WillReturnRows(rows)
	user := &User{}
	err := repository.fetchOne(user, sq.Eq{"id": 1}, orderBy)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, "foo", user.Name)
	assert.Equal(t, "bar", user.Email)
}

func Test_Models_Repositories_RepositoryAbstract_FetchAll(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT \\* FROM users").WillReturnRows(rows)
	collection := UsersCollection{}
	err := repository.fetchAll(&collection, nil, 0, 0, nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), collection[0].Id)
	assert.Equal(t, "foo", collection[0].Name)
	assert.Equal(t, "bar", collection[0].Email)
}

func Test_Models_Repositories_RepositoryAbstract_FetchAllWithLimitOffset(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT \\* FROM users LIMIT 10 OFFSET 20").WillReturnRows(rows)
	collection := UsersCollection{}
	err := repository.fetchAll(&collection, nil, 10, 20, nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), collection[0].Id)
	assert.Equal(t, "foo", collection[0].Name)
	assert.Equal(t, "bar", collection[0].Email)
}

func Test_Models_Repositories_RepositoryAbstract_FetchAllWithOrderBy(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT \\* FROM users ORDER BY id ASC").WillReturnRows(rows)
	collection := UsersCollection{}
	orderBy := map[string]string{"id": "ASC"}
	err := repository.fetchAll(&collection, nil, 0, 0, orderBy)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), collection[0].Id)
	assert.Equal(t, "foo", collection[0].Name)
	assert.Equal(t, "bar", collection[0].Email)
}

func Test_Models_Repositories_RepositoryAbstract_FetchColumn(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1)

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users").WillReturnRows(rows)
	row := repository.fetchColumn("COUNT(*)", nil)
	var result int
	err := row.Scan(&result)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func Test_Models_Repositories_RepositoryAbstract_Count(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1)

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TableUsers)
	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users").WillReturnRows(rows)
	result, err := repository.count(nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result)
}
