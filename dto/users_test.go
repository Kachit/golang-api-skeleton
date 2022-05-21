package dto

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_DTO_BindCreateUserDTOValidFullFilled(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"name": "Name", "email": "foo@bar.baz", "password": "pwd"}`
	c.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	obj, err := BindCreateUserDTO(c)
	assert.NoError(t, err)
	assert.Equal(t, "Name", obj.Name)
	assert.Equal(t, "foo@bar.baz", obj.Email)
	assert.Equal(t, "pwd", obj.Password)
}

func Test_DTO_BindCreateUserDTOInvalidEmptyName(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"email": "foo@bar.baz", "password": "pwd"}`
	c.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	_, err := BindCreateUserDTO(c)
	assert.Error(t, err)
	assert.Equal(t, "Key: 'CreateUserDTO.Name' Error:Field validation for 'Name' failed on the 'required' tag", err.Error())
}

func Test_DTO_BindCreateUserDTOInvalidEmptyEmail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"name": "Name ", "password": "pwd "}`
	c.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	_, err := BindCreateUserDTO(c)
	assert.Error(t, err)
	assert.Equal(t, "Key: 'CreateUserDTO.Email' Error:Field validation for 'Email' failed on the 'required' tag", err.Error())
}

func Test_DTO_BindCreateUserDTOInvalidWrongEmail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"name": "Name ", "password": "pwd ", "email": "foo"}`
	c.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	_, err := BindCreateUserDTO(c)
	assert.Error(t, err)
	assert.Equal(t, "Key: 'CreateUserDTO.Email' Error:Field validation for 'Email' failed on the 'email' tag", err.Error())
}

func Test_DTO_BindCreateUserDTOInvalidEmptyPassword(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"name": "Name ", "email": "foo@bar.baz"}`
	c.Request, _ = http.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	_, err := BindCreateUserDTO(c)
	assert.Error(t, err)
	assert.Equal(t, "Key: 'CreateUserDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag", err.Error())
}

func Test_DTO_BindEditUserDTOValidFullFilled(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"name": "Name", "email": "foo@bar.baz", "password": "pwd"}`
	c.Request, _ = http.NewRequest("PUT", "/v1/users/1", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	obj, err := BindEditUserDTO(c)
	assert.NoError(t, err)
	assert.Equal(t, "Name", obj.Name)
	assert.Equal(t, "foo@bar.baz", obj.Email)
	assert.Equal(t, "pwd", obj.Password)
}

func Test_DTO_BindEditUserDTOInvalidWrongEmail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonStr := `{"name": "Name ", "password": "pwd ", "email": "foo"}`
	c.Request, _ = http.NewRequest("PUT", "/v1/users/1", bytes.NewBufferString(jsonStr))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
	_, err := BindEditUserDTO(c)
	assert.Error(t, err)
	assert.Equal(t, "Key: 'EditUserDTO.Email' Error:Field validation for 'Email' failed on the 'email' tag", err.Error())
}
