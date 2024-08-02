package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-gonic/pkg/hello"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Use the router setup from main.go
	r := setupRouter()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	expected := `{"message":"` + hello.GetMessage() + `"}`
	assert.JSONEq(t, expected, rr.Body.String())
}
