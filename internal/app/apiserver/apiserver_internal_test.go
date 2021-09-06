package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Body.String(), "Hello")
}
