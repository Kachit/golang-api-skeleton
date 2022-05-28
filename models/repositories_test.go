package models

import (
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func Test_Models_Repositories_RepositoriesFactory_GetUsersRepository(t *testing.T) {
	db, _ := testable.GetDatabaseMock()
	rf := NewRepositoriesFactory(db)
	rep := rf.GetUsersRepository()
	assert.NotEmpty(t, rep)
}

type Models_Repositories_UsersRepository_TestSuite struct {
	suite.Suite
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	testable *UsersRepository
}

func (suite *Models_Repositories_UsersRepository_TestSuite) SetupTest() {
	db, mock := testable.GetDatabaseMock()
	suite.db = db
	suite.mock = mock
	suite.testable = &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)
}

func (suite *Models_Repositories_UsersRepository_TestSuite) TestGetByIdFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetById(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result.Id)
}

func (suite *Models_Repositories_UsersRepository_TestSuite) TestGetByIdNotFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	_, err := suite.testable.GetById(123)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "UsersRepository.GetById: record not found", err.Error())
}

func Test_Models_Repositories_UsersRepository_TestSuite(t *testing.T) {
	suite.Run(t, new(Models_Repositories_UsersRepository_TestSuite))
}

func Test_Models_Repositories_UsersRepository_GetByIdError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	_, err := rep.GetById(123)
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.GetById: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func Test_Models_Repositories_UsersRepository_GetListByFilter(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := rep.GetListByFilter()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result[0].Id)
}

func Test_Models_Repositories_UsersRepository_GetListByFilterError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	_, err := rep.GetListByFilter()
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.GetListByFilter: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func Test_Models_Repositories_UsersRepository_CountByFilterSuccess(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := rep.CountByFilter()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func Test_Models_Repositories_UsersRepository_CountByFilterError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	_, err := rep.CountByFilter()
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.CountByFilter: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func Test_Models_Repositories_UsersRepository_GetByEmailFound(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := rep.GetByEmail("foo@bar.baz")
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result.Id)
}

func Test_Models_Repositories_UsersRepository_GetByEmailNotFound(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	_, err := rep.GetByEmail("foo@bar.baz")
	assert.Error(t, err)
	assert.Equal(t, "UsersRepository.GetByEmail: record not found", err.Error())
}

func Test_Models_Repositories_UsersRepository_GetByEmailError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	_, err := rep.GetByEmail("foo@bar.baz")
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.GetByEmail: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func Test_Models_Repositories_UsersRepository_CountByEmail(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := rep.CountByEmail("foo@bar.baz")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func Test_Models_Repositories_UsersRepository_CountByEmailError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	_, err := rep.CountByEmail("foo@bar.baz")
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.CountByEmail: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func Test_Models_Repositories_UsersRepository_CreateSuccess(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	user := &User{Name: "foo", Email: "foo@bar.baz", Password: "pwd"}

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectCommit()

	err := rep.Create(user)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.NotEmpty(t, user.CreatedAt)
}

func Test_Models_Repositories_UsersRepository_CreateError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	user := &User{Name: "foo", Email: "foo@bar.baz", Password: "pwd"}

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	mock.ExpectRollback()

	err := rep.Create(user)
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.Create: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
	assert.NotEmpty(t, user.CreatedAt)
}

func Test_Models_Repositories_UsersRepository_EditSuccess(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	user := &User{Name: "foo", Email: "foo@bar.baz", Password: "pwd", Id: 1}

	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	err := rep.Edit(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ModifiedAt)
}

func Test_Models_Repositories_UsersRepository_EditError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rep := &UsersRepository{db: db}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	user := &User{Name: "foo", Email: "foo@bar.baz", Password: "pwd", Id: 1}

	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectCommit()

	err := rep.Edit(user)
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.Edit: all expectations were already fulfilled, call to database transaction Begin was not expected; invalid transaction`, err.Error())
	assert.NotEmpty(t, user.ModifiedAt)
}
