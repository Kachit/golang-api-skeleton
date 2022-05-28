package services

import (
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
)

func Test_Services_UsersService_GetListByFilterSuccess(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := service.GetListByFilter()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result[0].Id)
}

func Test_Services_UsersService_GetListByFilterError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	_, err := service.GetListByFilter()
	assert.Error(t, err)
	assert.Equal(t, `UsersService.GetListByFilter: UsersRepository.GetListByFilter: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func Test_Services_UsersService_CountByFilterSuccess(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	result, err := service.CountByFilter()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func Test_Services_UsersService_CountByFilterError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	_, err := service.CountByFilter()
	assert.Error(t, err)
	assert.Equal(t, `UsersService.CountByFilter: UsersRepository.CountByFilter: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func Test_Services_UsersService_GetByIdFound(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	bs := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := bs.GetById(123)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result.Id)
}

func Test_Services_UsersService_GetByIdNotFound(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	_, err := service.GetById(123)
	assert.Error(t, err)
	assert.Equal(t, "UsersService.GetById: UsersRepository.GetById: record not found", err.Error())
}

func Test_Services_UsersService_CheckIsUniqueEmailValid(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	err := service.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.NoError(t, err)
}

func Test_Services_UsersService_CreateSuccess(t *testing.T) {
	cfg, _ := testable.NewConfigMock()
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{
		RF: rf,
		PG: infrastructure.NewPasswordGenerator(cfg),
	})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectCommit()

	ud := &dto.CreateUserDTO{Name: "Name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := service.Create(ud)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, ud.Name, user.Name)
	assert.Equal(t, ud.Email, user.Email)
	assert.NotEqual(t, ud.Password, user.Password)
	assert.NotEmpty(t, user.CreatedAt)
}

func Test_Services_UsersService_CreateError(t *testing.T) {
	cfg, _ := testable.NewConfigMock()
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{
		RF: rf,
		PG: infrastructure.NewPasswordGenerator(cfg),
	})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","password","created_at","modified_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"foo"}).AddRow(1))

	mock.ExpectRollback()

	ud := &dto.CreateUserDTO{Name: "Name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := service.Create(ud)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, `UsersService.Create: UsersRepository.Create: sql: Scan error on column index 0, name "foo": unsupported Scan, storing driver.Value type int into type *models.User`, err.Error())
}

func Test_Services_UsersService_CreateNotUniqueEmail(t *testing.T) {
	cfg, _ := testable.NewConfigMock()
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{
		RF: rf,
		PG: infrastructure.NewPasswordGenerator(cfg),
	})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	ud := &dto.CreateUserDTO{Name: "Name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := service.Create(ud)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, `UsersService.Create: not unique user email`, err.Error())
}

func Test_Services_UsersService_EditSuccess(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{
		RF: rf,
	})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	ud := &dto.EditUserDTO{Name: "Name", Email: "foo@bar.baz"}
	user, err := service.Edit(1, ud)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, ud.Name, user.Name)
	assert.Equal(t, ud.Email, user.Email)
	assert.NotEmpty(t, user.ModifiedAt)
}

func Test_Services_UsersService_EditError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{
		RF: rf,
	})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"modified_at"=$5,"deleted_at"=$6 WHERE "id" = $7`)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectCommit()

	ud := &dto.EditUserDTO{Name: "Name", Email: "foo@bar.baz"}
	user, err := service.Edit(1, ud)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, `UsersService.Edit: UsersRepository.Edit: all expectations were already fulfilled, call to database transaction Begin was not expected; invalid transaction`, err.Error())
}

func Test_Services_UsersService_CheckIsUniqueEmailInvalid(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	err := service.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.Error(t, err)
	assert.Equal(t, "not unique user email", err.Error())
}

func Test_Services_UsersService_CheckIsUniqueEmailDatabaseError(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow("foo"))

	err := service.checkIsUniqueEmail("foo@bar.baz", nil)
	assert.Error(t, err)
	assert.Equal(t, `UsersRepository.CountByEmail: sql: Scan error on column index 0, name "count(*)": converting driver.Value type string ("foo") to a int64: invalid syntax`, err.Error())
}

func Test_Services_UsersService_CheckIsUniqueEmailSameUserValid(t *testing.T) {
	rf := models.NewRepositoriesFactory(nil)
	service := NewUsersService(&infrastructure.Container{RF: rf})

	err := service.checkIsUniqueEmail("foo@bar.baz", &models.User{Email: "foo@bar.baz"})
	assert.NoError(t, err)
}

func Test_Services_UsersService_CheckIsUniqueEmailNotSameUserInvalid(t *testing.T) {
	db, mock := testable.GetDatabaseMock()
	rf := models.NewRepositoriesFactory(db)
	service := NewUsersService(&infrastructure.Container{RF: rf})
	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "users" WHERE email = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	err := service.checkIsUniqueEmail("foo@bar.baz", &models.User{Email: "foo1@bar.baz"})
	assert.Error(t, err)
	assert.Equal(t, "not unique user email", err.Error())
}

func Test_Services_UsersService_BuildUserFromCreateUserDTO(t *testing.T) {
	cfg, _ := testable.NewConfigMock()
	service := NewUsersService(&infrastructure.Container{PG: infrastructure.NewPasswordGenerator(cfg)})
	userDto := &dto.CreateUserDTO{Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	user, err := service.buildUserFromCreateUserDTO(userDto)
	assert.NoError(t, err)
	assert.Equal(t, userDto.Name, user.Name)
	assert.Equal(t, userDto.Email, user.Email)
	assert.Equal(t, userDto.Password, user.Password)

	assert.Empty(t, user.Id)
	assert.Empty(t, user.CreatedAt)
}

func Test_Services_UsersService_FillUserFromEditUserDTO(t *testing.T) {
	service := NewUsersService(&infrastructure.Container{})
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	userDto := &dto.EditUserDTO{Name: "name1", Email: "foo1@bar.baz"}
	user, err := service.fillUserFromEditUserDTO(user, userDto)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.Id)
	assert.Equal(t, userDto.Name, user.Name)
	assert.Equal(t, userDto.Email, user.Email)
	assert.Equal(t, "pwd", user.Password)
	assert.Empty(t, user.CreatedAt)
}
