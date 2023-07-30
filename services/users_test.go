package services

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type ServicesUsersServiceTestSuite struct {
	suite.Suite
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	testable *UsersService
}

func (suite *ServicesUsersServiceTestSuite) SetupTest() {
	cfg, _ := infrastructure.NewConfigMock()
	db, mock := infrastructure.NewDatabaseMock()
	suite.db = db
	suite.mock = mock
	suite.testable = NewUsersService(&infrastructure.Container{
		DB: db,
		PG: infrastructure.NewPasswordGenerator(cfg),
	})
	mock.MatchExpectationsInOrder(false)
}

func (suite *ServicesUsersServiceTestSuite) TestGetListByFilterSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetListByFilter()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result[0].Id)
}

func (suite *ServicesUsersServiceTestSuite) TestGetListByFilterError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.GetListByFilter()
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersService.GetListByFilter: UsersRepository.GetListByFilter: SQLSTATE 01000`, err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestCountByFilterSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := suite.testable.CountByFilter()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(1), result)
}

func (suite *ServicesUsersServiceTestSuite) TestCountByFilterError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	result, err := suite.testable.CountByFilter()
	assert.Empty(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersService.CountByFilter: UsersRepository.CountByFilter: SQLSTATE 01000`, err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestGetByIdFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := suite.testable.GetById(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), result.Id)
}

func (suite *ServicesUsersServiceTestSuite) TestGetByIdNotFound() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	result, err := suite.testable.GetById(123)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "UsersService.GetById: UsersRepository.GetById: User not found", err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestCheckIsUniqueEmailValid() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.NoError(suite.T(), err)
}

func (suite *ServicesUsersServiceTestSuite) TestCreateSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	ud := dto.NewCreateUserDTOStub(nil)
	user, err := suite.testable.Create(ud)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.Equal(suite.T(), ud.Name, user.Name)
	assert.Equal(suite.T(), ud.Email, user.Email)
	assert.NotEqual(suite.T(), ud.Password, user.Password)
	assert.NotEmpty(suite.T(), user.CreatedAt)
}

func (suite *ServicesUsersServiceTestSuite) TestCreateError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	ud := dto.NewCreateUserDTOStub(nil)
	user, err := suite.testable.Create(ud)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), `UsersService.Create: UsersRepository.Create: SQLSTATE 01000`, err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestCreateNotUniqueEmail() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	ud := dto.NewCreateUserDTOStub(nil)
	user, err := suite.testable.Create(ud)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), `UsersService.Create: not unique user email`, err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestEditSuccess() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ud := dto.NewEditUserDTOStub(nil)
	user, err := suite.testable.Edit(1, ud)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint64(1), user.Id)
	assert.Equal(suite.T(), ud.Name, user.Name)
	assert.Equal(suite.T(), ud.Email, user.Email)
	assert.NotEmpty(suite.T(), user.ModifiedAt)
}

func (suite *ServicesUsersServiceTestSuite) TestEditError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	ud := dto.NewEditUserDTOStub(nil)
	user, err := suite.testable.Edit(1, ud)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), `UsersService.Edit: UsersRepository.Edit: SQLSTATE 01000`, err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestCheckIsUniqueEmailInvalid() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not unique user email", err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestCheckIsUniqueEmailError() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnError(fmt.Errorf("SQLSTATE 01000"))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `UsersRepository.CountByEmail: SQLSTATE 01000`, err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestCheckIsUniqueEmailSameUserValid() {
	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", &entities.User{Email: "foo@bar.baz"})
	assert.NoError(suite.T(), err)
}

func (suite *ServicesUsersServiceTestSuite) TestCheckIsUniqueEmailNotSameUserInvalid() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	err := suite.testable.checkIsUniqueEmail("foo@bar.baz", &entities.User{Email: "foo1@bar.baz"})
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not unique user email", err.Error())
}

func (suite *ServicesUsersServiceTestSuite) TestBuildUserFromCreateUserDTO() {
	userDto := dto.NewCreateUserDTOStub(nil)
	user, err := suite.testable.buildUserFromCreateUserDTO(userDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), userDto.Name, user.Name)
	assert.Equal(suite.T(), userDto.Email, user.Email)
	assert.Equal(suite.T(), userDto.Password, user.Password)

	assert.Empty(suite.T(), user.Id)
	assert.Empty(suite.T(), user.CreatedAt)
}

func (suite *ServicesUsersServiceTestSuite) TestFillUserFromEditUserDTO() {
	user := entities.NewUserEntityStub(nil)
	userId := user.Id
	userPwd := user.Password
	userDto := dto.NewEditUserDTOStub(nil)
	user, err := suite.testable.fillUserFromEditUserDTO(user, userDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), userId, user.Id)
	assert.Equal(suite.T(), userDto.Name, user.Name)
	assert.Equal(suite.T(), userDto.Email, user.Email)
	assert.Equal(suite.T(), userPwd, user.Password)
}

func TestServicesUsersServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServicesUsersServiceTestSuite))
}
