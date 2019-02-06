package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()                         // HL
	w := httptest.NewRecorder()                     // HL
	req, _ := http.NewRequest("GET", "/hello", nil) // HL

	router.ServeHTTP(w, req) // HL

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hi.", w.Body.String())
}
