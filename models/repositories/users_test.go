package repositories

import (
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type ModelsRepositoriesUsersRepositoryTestSuite struct {
	suite.Suite
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	testable *UsersRepository
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) SetupTest() {
	db, mock := infrastructure.GetDatabaseMock()
	suite.db = db
	suite.mock = mock
	suite.testable = NewUsersRepository(db)
	mock.MatchExpectationsInOrder(false)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByIdFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetById(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result.Id)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByIdNotFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	result, err := suite.testable.GetById(123)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "UsersRepository.GetById: record not found", err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByIdError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	result, err := suite.testable.GetById(123)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.GetById: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *entities.User`, err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetListByFilterSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetListByFilter()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result[0].Id)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetListByFilterError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	result, err := suite.testable.GetListByFilter()
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.GetListByFilter: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *entities.User`, err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCountByFilterSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := suite.testable.CountByFilter()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(1), result)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCountByFilterError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	result, err := suite.testable.CountByFilter()
	assert.Empty(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.CountByFilter: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByEmailFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetByEmail("foo@bar.baz")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result.Id)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByEmailNotFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	result, err := suite.testable.GetByEmail("foo@bar.baz")
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "UsersRepository.GetByEmail: record not found", err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByEmailError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	result, err := suite.testable.GetByEmail("foo@bar.baz")
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.GetByEmail: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *entities.User`, err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCountByEmailSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := suite.testable.CountByEmail("foo@bar.baz")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(1), result)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCountByEmailError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	result, err := suite.testable.CountByEmail("foo@bar.baz")
	assert.Empty(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.CountByEmail: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCreateSuccess() {
	suite.mock.ExpectBegin()

	user := &entities.User{Name: "foo", Email: "foo@bar.baz", Password: "pwd"}

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	suite.mock.ExpectCommit()

	err := suite.testable.Create(user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.NotEmpty(suite.T(), user.CreatedAt)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCreateError() {
	suite.mock.ExpectBegin()

	user := &entities.User{Name: "foo", Email: "foo@bar.baz", Password: "pwd"}

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	suite.mock.ExpectRollback()

	err := suite.testable.Create(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.Create: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *entities.User`, err.Error())
	assert.NotEmpty(suite.T(), user.CreatedAt)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestEditSuccess() {
	suite.mock.ExpectBegin()

	user := &entities.User{Name: "foo", Email: "foo@bar.baz", Password: "pwd", Id: 1}

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	suite.mock.ExpectCommit()

	err := suite.testable.Edit(user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ModifiedAt)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestEditError() {
	suite.mock.ExpectBegin()

	user := &entities.User{Name: "foo", Email: "foo@bar.baz", Password: "pwd", Id: 1}

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	suite.mock.ExpectCommit()

	err := suite.testable.Edit(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.Edit: all expectations were already fulfilled, call to database transaction Begin was not expected`, err.Error())
	assert.NotEmpty(suite.T(), user.ModifiedAt)
}

func TestModelsRepositoriesUsersRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsRepositoriesUsersRepositoryTestSuite))
}
