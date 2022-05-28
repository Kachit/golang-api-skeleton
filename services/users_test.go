package services

import (
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type Services_UsersService_TestSuite struct {
	suite.Suite
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	testable *UsersService
}

func (suite *Services_UsersService_TestSuite) SetupTest() {
	cfg, _ := testable.NewConfigMock()
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	suite.db = db
	suite.mock = mock
	suite.testable = NewUsersService(&infrastructure.Container{
		RF: rf,
		PG: infrastructure.NewPasswordGenerator(cfg),
	})
	mock.MatchExpectationsInOrder(false)
}

func (suite *Services_UsersService_TestSuite) TestGetListByFilterSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetListByFilter()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result[0].Id)
}

func (suite *Services_UsersService_TestSuite) TestGetListByFilterError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	result, err := suite.testable.GetListByFilter()
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersService.GetListByFilter: UsersRepository.GetListByFilter: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func (suite *Services_UsersService_TestSuite) TestCountByFilterSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := suite.testable.CountByFilter()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(1), result)
}

func (suite *Services_UsersService_TestSuite) TestCountByFilterError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	result, err := suite.testable.CountByFilter()
	assert.Empty(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersService.CountByFilter: UsersRepository.CountByFilter: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func (suite *Services_UsersService_TestSuite) TestGetByIdFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetById(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result.Id)
}

func (suite *Services_UsersService_TestSuite) TestGetByIdNotFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	result, err := suite.testable.GetById(123)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "UsersService.GetById: UsersRepository.GetById: record not found", err.Error())
}

func (suite *Services_UsersService_TestSuite) TestCheckIsUniqueEmailValid() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.NoError(suite.T(), err)
}

func (suite *Services_UsersService_TestSuite) TestCreateSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectBegin()

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	suite.mock.ExpectCommit()

	ud := &dto.CreateUserDTO{Name: "Name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := suite.testable.Create(ud)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.Equal(suite.T(), ud.Name, user.Name)
	assert.Equal(suite.T(), ud.Email, user.Email)
	assert.NotEqual(suite.T(), ud.Password, user.Password)
	assert.NotEmpty(suite.T(), user.CreatedAt)
}

func (suite *Services_UsersService_TestSuite) TestCreateError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectBegin()

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	suite.mock.ExpectRollback()

	ud := &dto.CreateUserDTO{Name: "Name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := suite.testable.Create(ud)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), `UsersService.Create: UsersRepository.Create: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func (suite *Services_UsersService_TestSuite) TestCreateNotUniqueEmail() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	ud := &dto.CreateUserDTO{Name: "Name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := suite.testable.Create(ud)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), `UsersService.Create: not unique user email`, err.Error())
}

func (suite *Services_UsersService_TestSuite) TestEditSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	suite.mock.ExpectBegin()

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	suite.mock.ExpectCommit()

	ud := &dto.EditUserDTO{Name: "Name", Email: "foo@bar.baz"}
	user, err := suite.testable.Edit(1, ud)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.Equal(suite.T(), ud.Name, user.Name)
	assert.Equal(suite.T(), ud.Email, user.Email)
	assert.NotEmpty(suite.T(), user.ModifiedAt)
}

func (suite *Services_UsersService_TestSuite) TestEditError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	suite.mock.ExpectBegin()

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	suite.mock.ExpectCommit()

	ud := &dto.EditUserDTO{Name: "Name", Email: "foo@bar.baz"}
	user, err := suite.testable.Edit(1, ud)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), `UsersService.Edit: UsersRepository.Edit: all expectations were already fulfilled, call to database transaction Begin was not expected; invalid transaction`, err.Error())
}

func (suite *Services_UsersService_TestSuite) TestCheckIsUniqueEmailInvalid() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not unique user email", err.Error())
}

func (suite *Services_UsersService_TestSuite) TestCheckIsUniqueEmailError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.CountByEmail: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func (suite *Services_UsersService_TestSuite) TestCheckIsUniqueEmailSameUserValid() {
	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", &models.User{Email: "foo@bar.baz"})
	assert.NoError(suite.T(), err)
}

func (suite *Services_UsersService_TestSuite) TestCheckIsUniqueEmailNotSameUserInvalid() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", &models.User{Email: "foo1@bar.baz"})
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not unique user email", err.Error())
}

func (suite *Services_UsersService_TestSuite) TestBuildUserFromCreateUserDTO() {
	userDto := &dto.CreateUserDTO{Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := suite.testable.buildUserFromCreateUserDTO(userDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), userDto.Name, user.Name)
	assert.Equal(suite.T(), userDto.Email, user.Email)
	assert.Equal(suite.T(), userDto.Password, user.Password)

	assert.Empty(suite.T(), user.Id)
	assert.Empty(suite.T(), user.CreatedAt)
}

func (suite *Services_UsersService_TestSuite) TestFillUserFromEditUserDTO() {
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	userDto := &dto.EditUserDTO{Name: "name1", Email: "foo1@bar.baz"}
	user, err := suite.testable.fillUserFromEditUserDTO(user, userDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.Equal(suite.T(), userDto.Name, user.Name)
	assert.Equal(suite.T(), userDto.Email, user.Email)
	assert.Equal(suite.T(), "pwd", user.Password)
	assert.Empty(suite.T(), user.CreatedAt)
}

func Test_Services_UsersService_TestSuite(t *testing.T) {
	suite.Run(t, new(Services_UsersService_TestSuite))
}
