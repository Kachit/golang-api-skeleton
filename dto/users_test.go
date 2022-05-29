package dto

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DTOUsersTestSuite struct {
	suite.Suite
	gin *gin.Context
}

func (suite *DTOUsersTestSuite) SetupTest() {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	suite.gin = c
}

func (suite *DTOUsersTestSuite) TestBindCreateUserDTOValidFullFilled() {
	jsonStr := `{"name": "Name", "email": "foo@bar.baz", "password": "pwd"}`
	suite.gin.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	obj, err := BindCreateUserDTO(suite.gin)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Name", obj.Name)
	assert.Equal(suite.T(), "foo@bar.baz", obj.Email)
	assert.Equal(suite.T(), "pwd", obj.Password)
}

func (suite *DTOUsersTestSuite) TestBindCreateUserDTOInvalidEmptyName() {
	jsonStr := `{"email": "foo@bar.baz", "password": "pwd"}`
	suite.gin.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	result, err := BindCreateUserDTO(suite.gin)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Key: 'CreateUserDTO.Name' Error:Field validation for 'Name' failed on the 'required' tag", err.Error())
}

func (suite *DTOUsersTestSuite) TestBindCreateUserDTOInvalidEmptyEmail() {
	jsonStr := `{"name": "Name ", "password": "pwd "}`
	suite.gin.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	result, err := BindCreateUserDTO(suite.gin)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Key: 'CreateUserDTO.Email' Error:Field validation for 'Email' failed on the 'required' tag", err.Error())
}

func (suite *DTOUsersTestSuite) TestBindCreateUserDTOInvalidWrongEmail() {
	jsonStr := `{"name": "Name ", "password": "pwd ", "email": "foo"}`
	suite.gin.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	result, err := BindCreateUserDTO(suite.gin)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Key: 'CreateUserDTO.Email' Error:Field validation for 'Email' failed on the 'email' tag", err.Error())
}

func (suite *DTOUsersTestSuite) TestBindCreateUserDTOInvalidEmptyPassword() {
	jsonStr := `{"name": "Name ", "email": "foo@bar.baz"}`
	suite.gin.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	result, err := BindCreateUserDTO(suite.gin)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Key: 'CreateUserDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag", err.Error())
}

func (suite *DTOUsersTestSuite) TestBindEditUserDTOValidFullFilled() {
	jsonStr := `{"name": "Name", "email": "foo@bar.baz"}`
	suite.gin.Request, _ = http.NewRequest("PUT", "/v1/users/1", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	result, err := BindEditUserDTO(suite.gin)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Name", result.Name)
	assert.Equal(suite.T(), "foo@bar.baz", result.Email)
}

func (suite *DTOUsersTestSuite) TestBindEditUserDTOInvalidWrongEmail() {
	jsonStr := `{"name": "Name ", "email": "foo"}`
	suite.gin.Request, _ = http.NewRequest("PUT", "/v1/users/1", bytes.NewBufferString(jsonStr))
	suite.gin.Request.Header.Add("Content-Type", gin.MIMEJSON)
	result, err := BindEditUserDTO(suite.gin)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Key: 'EditUserDTO.Email' Error:Field validation for 'Email' failed on the 'email' tag", err.Error())
}

func TestDTOUsersTestSuite(t *testing.T) {
	suite.Run(t, new(DTOUsersTestSuite))
}
