package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Utils_DumpHttpRequestWithoutBody(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://foo.bar/v1/boxes?foo=bar", nil)
	req.Header.Add("Content-Type", gin.MIMEJSON)
	req.Header.Add("X-Auth-Token", "qwerty")
	result := DumpHttpRequest(req)
	assert.Equal(t, req.Header, result["headers"])
	assert.Equal(t, req.URL.String(), result["url"])
	assert.Equal(t, req.Method, result["method"])
}

func Test_Utils_DumpHttpRequestWithBody(t *testing.T) {
	jsonStr := `{"user_id": 123, "mode": "demo", "box_price": 100}`
	req, _ := http.NewRequest("POST", "https://foo.bar/v1/boxes", bytes.NewBufferString(jsonStr))
	req.Header.Add("Content-Type", gin.MIMEJSON)
	req.Header.Add("X-Auth-Token", "qwerty")
	result := DumpHttpRequest(req)
	buf := &bytes.Buffer{}
	buf.ReadFrom(req.Body)
	assert.Equal(t, req.Header, result["headers"])
	assert.Equal(t, req.URL.String(), result["url"])
	assert.Equal(t, req.Method, result["method"])
	assert.Equal(t, jsonStr, result["body"])
	assert.Equal(t, jsonStr, string(buf.Bytes()))
}
