package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestEchoServer(t *testing.T) {
	e := echo.New()
	e.GET("/api/v1/health", echoHealthCheck)
	requestBody := `{"name":"Jon Snow","age":17}`

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/1", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.Validator = &CustomValidator{validator: validator.New()}
	c := e.NewContext(req, rec)

	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, echoHealthCheck(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	want := `{"id":1,"name":"Jon Snow","age":17}`
	assert.JSONEq(t, want, rec.Body.String())
}
