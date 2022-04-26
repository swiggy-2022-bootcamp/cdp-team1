package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	router := setupRouter(InitCustomerController(nil))

	req, _ := http.NewRequest(http.MethodGet, "/api/customer-admin/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"Server is Up and Running"}`, w.Body.String())

}
