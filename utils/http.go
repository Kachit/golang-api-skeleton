package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func DumpHttpRequest(req *http.Request) map[string]interface{} {
	dump := make(map[string]interface{})
	dump["method"] = req.Method
	dump["url"] = req.URL.String()
	dump["headers"] = req.Header
	if req.Body != nil {
		buf := &bytes.Buffer{}
		buf.ReadFrom(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(buf.Bytes())) // Write body back
		reqBody := string(buf.Bytes())
		dump["body"] = strings.ReplaceAll(strings.ReplaceAll(reqBody, "\n", ""), "  ", "")
	}
	return dump
}
