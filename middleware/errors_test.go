package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/lajosbencz/glo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Middleware_HttpErrorHandlerMiddlewareSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)

	ct := c.Writer.Header().Get("Content-Type")
	logger := testable.GetLoggerMock()
	mdl := HttpErrorHandlerMiddleware(logger)
	mdl(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.NotEqual(t, ct, c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, "application/json", c.Writer.Header().Get("Content-Type"))
	assert.Empty(t, logger.Level)
	assert.Empty(t, logger.Msg)
	assert.Empty(t, logger.Params)
}

func Test_Middleware_HttpErrorHandlerMiddlewareError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)
	ct := c.Writer.Header().Get("Content-Type")
	c.AbortWithError(http.StatusForbidden, errors.New("foo bar"))
	logger := testable.GetLoggerMock()
	mdl := HttpErrorHandlerMiddleware(logger)
	mdl(c)
	assert.Equal(t, http.StatusForbidden, c.Writer.Status())
	assert.NotEqual(t, ct, c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, "application/json", c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, glo.Error, logger.Level)
	assert.Equal(t, "foo bar", logger.Msg)
	assert.NotEmpty(t, logger.Params)
}

func Test_Middleware_ConvertErrorToHttpError(t *testing.T) {
	err := errors.New("BoxesResource.Open: BoxesService.Open: box opening price is not equal")
	result := convertErrorToHttpError(err)
	assert.Equal(t, "box opening price is not equal", result.Error())
}

func Test_Middleware_ConvertErrorMessageWithColon(t *testing.T) {
	msg := "BoxesResource.Open: BoxesService.Open: box opening price is not equal"
	result := convertErrorMessage(msg)
	assert.Equal(t, "box opening price is not equal", result)
}

func Test_Middleware_ConvertErrorMessageWithoutColon(t *testing.T) {
	msg := "box opening price is not equal"
	result := convertErrorMessage(msg)
	assert.Equal(t, "box opening price is not equal", result)
}

func Test_Middleware_ConvertErrorMessageSqlError(t *testing.T) {
	msg := `BoxesResource.Open: BoxesService.Open: duplicate key value violates unique constraint "ux_boo_spin_id" (SQLSTATE 23505)`
	result := convertErrorMessage(msg)
	assert.Equal(t, "Database error", result)
}
