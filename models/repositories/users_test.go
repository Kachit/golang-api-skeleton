package repositories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
	db, mock := infrastructure.NewDatabaseMock()
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
	assert.Equal(suite.T(), "UsersRepository.GetById: User not found", err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByIdError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.GetById(123)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.GetById: SQLSTATE 01000`, err.Error())
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
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.GetListByFilter()
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.GetListByFilter: SQLSTATE 01000`, err.Error())
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
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.CountByFilter()
	assert.Empty(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.CountByFilter: SQLSTATE 01000`, err.Error())
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
	assert.Equal(suite.T(), "UsersRepository.GetByEmail: User not found", err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestGetByEmailError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.GetByEmail("foo@bar.baz")
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.GetByEmail: SQLSTATE 01000`, err.Error())
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
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.CountByEmail("foo@bar.baz")
	assert.Empty(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.CountByEmail: SQLSTATE 01000`, err.Error())
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCreateSuccess() {
	user := entities.NewUserEntityStub(nil)
	user.Id = 0

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err := suite.testable.Create(user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.NotEmpty(suite.T(), user.CreatedAt)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestCreateError() {
	user := entities.NewUserEntityStub(nil)
	user.Id = 0

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	err := suite.testable.Create(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.Create: SQLSTATE 01000`, err.Error())
	assert.NotEmpty(suite.T(), user.CreatedAt)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestEditSuccess() {
	user := entities.NewUserEntityStub(nil)

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.testable.Edit(user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ModifiedAt)
}

func (suite *ModelsRepositoriesUsersRepositoryTestSuite) TestEditError() {
	user := entities.NewUserEntityStub(nil)

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	err := suite.testable.Edit(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.Edit: SQLSTATE 01000`, err.Error())
	assert.NotEmpty(suite.T(), user.ModifiedAt)
}

func TestModelsRepositoriesUsersRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsRepositoriesUsersRepositoryTestSuite))
}
