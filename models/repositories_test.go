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
	result := factory.GetRepositoryAbstract(TABLE_USERS)
	assert.IsType(t, (*RepositoryAbstract)(nil), result)
	assert.Equal(t, TABLE_USERS, result.table)
}

func Test_Models_Repositories_RepositoriesFactory_GetUsersRepository(t *testing.T) {
	factory := NewRepositoriesFactory(nil)
	result := factory.GetUsersRepository()
	assert.IsType(t, (*UsersRepository)(nil), result)
	assert.Equal(t, TABLE_USERS, result.table)
}

func Test_Models_Repositories_RepositoryAbstract_Insert(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
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
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
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
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
	mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)
	user := &User{}
	err := repository.fetchOne(user, sq.Eq{"id": 1})
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, "foo", user.Name)
	assert.Equal(t, "bar", user.Email)
}

func Test_Models_Repositories_RepositoryAbstract_FetchAll(t *testing.T) {
	sqlxDB, mock := testable.GetDatabaseMock()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "foo", "bar")

	factory := NewRepositoriesFactory(sqlxDB)
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
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
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
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
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
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
	repository := factory.GetRepositoryAbstract(TABLE_USERS)
	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users").WillReturnRows(rows)
	row := repository.fetchColumn("COUNT(*)", nil)
	var result int
	err := row.Scan(&result)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
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
