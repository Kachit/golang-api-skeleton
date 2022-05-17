package rest

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Rest_NewResponseBody(t *testing.T) {
	data := map[string]string{"foo": "bar"}
	result := NewResponseBody(data)
	bytesArr, _ := json.Marshal(result)
	assert.True(t, result.Result)
	assert.Equal(t, data, result.Data)
	assert.Nil(t, result.Meta)
	assert.Empty(t, result.Error)
	assert.Equal(t, `{"result":true,"data":{"foo":"bar"},"meta":null,"error":""}`, string(bytesArr))
}

func Test_Rest_NewResponseBodyError(t *testing.T) {
	result := NewResponseBodyError(fmt.Errorf("foo"))
	bytesArr, _ := json.Marshal(result)
	assert.False(t, result.Result)
	assert.Nil(t, result.Data)
	assert.Nil(t, result.Meta)
	assert.Equal(t, "foo", result.Error)
	assert.Equal(t, `{"result":false,"data":null,"meta":null,"error":"foo"}`, string(bytesArr))
}

func Test_Rest_NewResponseBodyWithPagination(t *testing.T) {
	data := map[string]string{"foo": "bar"}
	result := NewResponseBodyWithPagination(data, 100, 10)
	bytesArr, _ := json.Marshal(result)
	assert.True(t, result.Result)
	assert.Equal(t, data, result.Data)
	assert.Empty(t, result.Error)
	assert.Equal(t, `{"result":true,"data":{"foo":"bar"},"meta":{"pagination":{"total":100,"count":10}},"error":""}`, string(bytesArr))
}
