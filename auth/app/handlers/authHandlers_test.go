package handlers

import (
	"authService/dto"
	"authService/errs"
	"authService/mocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandlers_HealthCheckHandler(t *testing.T) {

	t.Run("All Services Healthy", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		mockAuthSvc.On("HealthCheck").Return(nil)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.GET("/api/auth/", authHandlers.HealthCheckHandler)

		req, _ := http.NewRequest(http.MethodGet, "/api/auth/", nil)

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "HealthCheck", 1)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Any Service Unhealthy", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		mockErr := errs.NewUnexpectedError("mock error")
		mockAuthSvc.On("HealthCheck").Return(mockErr)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.GET("/api/auth/", authHandlers.HealthCheckHandler)

		req, _ := http.NewRequest(http.MethodGet, "/api/auth/", nil)

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "HealthCheck", 1)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestAuthHandlers_LoginHandler(t *testing.T) {

	t.Run("Successful Admin Login", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		username := "ayan59dutta"
		password := "pass@123!"
		mockAuthSvc.On("Login", username, "", password).Return("mock-token", nil)

		reqBody, _ := json.Marshal(&dto.LoginRequestDTO{
			Username: username,
			Password: password,
		})

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/login/", authHandlers.LoginHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login/", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "Login", 1)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Successful Customer Login", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		email := "ayan59dutta@gmail.com"
		password := "pass@123!"
		mockAuthSvc.On("Login", "", email, password).Return("mock-token", nil)

		reqBody, _ := json.Marshal(&dto.LoginRequestDTO{
			Email:    email,
			Password: password,
		})

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/login/", authHandlers.LoginHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login/", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "Login", 1)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Unsuccessful Login", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		mockErr := errs.NewUnexpectedError("mock error")
		username := "ayan59dutta"
		password := "wrongPassword"
		mockAuthSvc.On("Login", username, "", password).Return("", mockErr)

		reqBody, _ := json.Marshal(&dto.LoginRequestDTO{
			Username: username,
			Password: password,
		})

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/login/", authHandlers.LoginHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login/", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "Login", 1)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Invalid Login Request", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		mockAuthSvc.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("mock-token", nil)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/login/", authHandlers.LoginHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login/", nil)

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "Login", 0)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestAuthHandlers_LogoutHandler(t *testing.T) {

	t.Run("Successful Logout", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		token := "mock-token"
		mockAuthSvc.On("VerifyAuthToken", token, "", "").Return(true, nil)
		mockAuthSvc.On("Logout", token).Return(nil)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/logout/", authHandlers.LogoutHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/logout/", nil)
		req.Header.Add("Authorization", token)

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "VerifyAuthToken", 1)
		mockAuthSvc.AssertNumberOfCalls(t, "Logout", 1)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Erroneous Logout", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		token := "mock-token"
		mockErr := errs.NewUnexpectedError("mock-error")
		mockAuthSvc.On("VerifyAuthToken", token, "", "").Return(true, nil)
		mockAuthSvc.On("Logout", token).Return(mockErr)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/logout/", authHandlers.LogoutHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/logout/", nil)
		req.Header.Add("Authorization", token)

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "VerifyAuthToken", 1)
		mockAuthSvc.AssertNumberOfCalls(t, "Logout", 1)
		assert.Equal(t, mockErr.Code, w.Code)
	})
}

func TestAuthHandlers_VerificationHandler(t *testing.T) {

	t.Run("No Body", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		token := "mock-token"
		mockAuthSvc.On("VerifyAuthToken", token, mock.Anything, mock.Anything).Return(true, nil)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/verify/", authHandlers.VerificationHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/verify/", nil)
		req.Header.Add("Authorization", token)

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "VerifyAuthToken", 0)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Valid Token", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		token := "mock-token"
		mockAuthSvc.On("VerifyAuthToken", token, mock.Anything, mock.Anything).Return(true, nil)

		reqBody, _ := json.Marshal(&dto.VerificationRequestDTO{
			UserID: "123",
			Role:   "admin",
		})

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/verify/", authHandlers.VerificationHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/verify/", bytes.NewReader(reqBody))
		req.Header.Add("Authorization", token)
		req.Header.Add("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "VerifyAuthToken", 1)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Invalid Token", func(t *testing.T) {

		mockAuthSvc := mocks.AuthService{}
		authHandlers := AuthHandlers{AuthSvc: &mockAuthSvc}
		token := "mock-token"
		mockErr := errs.NewUnexpectedError("mock-error")
		mockAuthSvc.On("VerifyAuthToken", token, mock.Anything, mock.Anything).Return(false, mockErr)

		reqBody, _ := json.Marshal(&dto.VerificationRequestDTO{
			UserID: "123",
			Role:   "admin",
		})

		w := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/api/auth/verify/", authHandlers.VerificationHandler)

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/verify/", bytes.NewReader(reqBody))
		req.Header.Add("Authorization", token)
		req.Header.Add("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		mockAuthSvc.AssertNumberOfCalls(t, "VerifyAuthToken", 1)
		assert.Equal(t, mockErr.Code, w.Code)
	})
}
