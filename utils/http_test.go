package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type UtilsHttpTestSuite struct {
	suite.Suite
}

func (suite *UtilsHttpTestSuite) TestDumpHttpRequestWithoutBody() {
	req, _ := http.NewRequest("GET", "https://foo.bar/v1/users?foo=bar", nil)
	req.Header.Add("Content-Type", gin.MIMEJSON)
	req.Header.Add("X-Auth-Token", "qwerty")
	result := DumpHttpRequest(req)
	assert.Equal(suite.T(), req.Header, result["headers"])
	assert.Equal(suite.T(), req.URL.String(), result["url"])
	assert.Equal(suite.T(), req.Method, result["method"])
}

func (suite *UtilsHttpTestSuite) TestDumpHttpRequestWithBody() {
	jsonStr := `{"user_id": 123, "mode": "demo", "box_price": 100}`
	req, _ := http.NewRequest("POST", "https://foo.bar/v1/users", bytes.NewBufferString(jsonStr))
	req.Header.Add("Content-Type", gin.MIMEJSON)
	req.Header.Add("X-Auth-Token", "qwerty")
	result := DumpHttpRequest(req)
	buf := &bytes.Buffer{}
	buf.ReadFrom(req.Body)
	assert.Equal(suite.T(), req.Header, result["headers"])
	assert.Equal(suite.T(), req.URL.String(), result["url"])
	assert.Equal(suite.T(), req.Method, result["method"])
	assert.Equal(suite.T(), jsonStr, result["body"])
	assert.Equal(suite.T(), jsonStr, string(buf.Bytes()))
}

func TestUtilsHttpTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsHttpTestSuite))
}
