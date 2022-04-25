package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldGetAllAdmins(t *testing.T) {
	customerController := &CustomerController{}
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodGet, "/api/customer-admin/user", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
