package rest

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RestTestSuite struct {
	suite.Suite
}

func (suite *RestTestSuite) TestNewResponseBody() {
	data := map[string]string{"foo": "bar"}
	result := NewResponseBody(data)
	bytesArr, _ := json.Marshal(result)
	assert.True(suite.T(), result.Result)
	assert.Equal(suite.T(), data, result.Data)
	assert.Nil(suite.T(), result.Meta)
	assert.Empty(suite.T(), result.Error)
	assert.Equal(suite.T(), `{"result":true,"data":{"foo":"bar"},"meta":null,"error":""}`, string(bytesArr))
}

func (suite *RestTestSuite) TestNewResponseBodyError() {
	result := NewResponseBodyError(fmt.Errorf("foo"))
	bytesArr, _ := json.Marshal(result)
	assert.False(suite.T(), result.Result)
	assert.Nil(suite.T(), result.Data)
	assert.Nil(suite.T(), result.Meta)
	assert.Equal(suite.T(), "foo", result.Error)
	assert.Equal(suite.T(), `{"result":false,"data":null,"meta":null,"error":"foo"}`, string(bytesArr))
}

func (suite *RestTestSuite) TestNewResponseBodyWithPagination() {
	data := map[string]string{"foo": "bar"}
	result := NewResponseBodyWithPagination(data, 100, 10)
	bytesArr, _ := json.Marshal(result)
	assert.True(suite.T(), result.Result)
	assert.Equal(suite.T(), data, result.Data)
	assert.Empty(suite.T(), result.Error)
	assert.Equal(suite.T(), `{"result":true,"data":{"foo":"bar"},"meta":{"pagination":{"total":100,"count":10}},"error":""}`, string(bytesArr))
}

func TestRestTestSuite(t *testing.T) {
	suite.Run(t, new(RestTestSuite))
}
