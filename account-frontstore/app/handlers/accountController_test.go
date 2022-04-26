package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/service/mocks"
	"qwik.in/account-frontstore/internal/errors"
	"testing"
)

func setupRouter(accountController AccountControllerInterface) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	newRouter := r.Group("api/account-frontstore")
	newRouter.POST("/register", accountController.RegisterAccount)
	newRouter.GET("/account/", accountController.GetAccountById)
	newRouter.PUT("/account/", accountController.UpdateAccount)

	newRouter.GET("/", HealthCheck)
	return r
}

func TestRegisterAccountForSuccess(t *testing.T) {
	sampleAccountBody := model.Account{
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("CreateAccount", sampleAccountBody).Return(&sampleAccountBody, nil)

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodPost, "/api/account-frontstore/register", bytes.NewBuffer(marshalledSampleAccountBody))
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestRegisterAccountForFailure(t *testing.T) {
	sampleAccountBody := model.Account{
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("CreateAccount", sampleAccountBody).Return(nil, errors.NewEmailAlreadyRegisteredError())

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodPost, "/api/account-frontstore/register", bytes.NewBuffer(marshalledSampleAccountBody))
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"User with given email already exists\"", w.Body.String())
}

func TestGetAccountByIdForSuccess(t *testing.T) {
	sampleAccountBody := model.Account{
		CustomerId:      "12345",
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	//marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("GetAccountById", sampleAccountBody.CustomerId).Return(&sampleAccountBody, nil)

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodGet, "/api/account-frontstore/account/", nil)
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetAccountByIdForFailure(t *testing.T) {
	sampleAccountBody := model.Account{
		CustomerId:      "12345",
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	//marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("GetAccountById", sampleAccountBody.CustomerId).Return(nil, errors.NewUserNotFoundError())

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodGet, "/api/account-frontstore/account/", nil)
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"User not found\"", w.Body.String())
}

func TestUpdateAccountForSuccess(t *testing.T) {
	sampleAccountBody := model.Account{
		CustomerId:      "12345",
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("UpdateAccount", sampleAccountBody.CustomerId, sampleAccountBody).Return(&sampleAccountBody, nil)

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodPut, "/api/account-frontstore/account/", bytes.NewBuffer(marshalledSampleAccountBody))
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateAccountForFailure(t *testing.T) {
	sampleAccountBody := model.Account{
		CustomerId:      "12345",
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("UpdateAccount", sampleAccountBody.CustomerId, sampleAccountBody).Return(nil, errors.NewUserNotFoundError())

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodPut, "/api/account-frontstore/account/", bytes.NewBuffer(marshalledSampleAccountBody))
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"User not found\"", w.Body.String())
}

func TestUpdateAccountForEmailFailure(t *testing.T) {
	sampleAccountBody := model.Account{
		CustomerId:      "12345",
		Firstname:       "Ravikumar",
		Lastname:        "S",
		Email:           "ravi@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Agree:           1,
	}
	marshalledSampleAccountBody, _ := json.Marshal(sampleAccountBody)

	mockAccountService := &mocks.AccountServiceInterface{}
	mockAccountService.On("UpdateAccount", sampleAccountBody.CustomerId, sampleAccountBody).Return(nil, errors.NewEmailAlreadyRegisteredError())

	accountController := InitAccountController(mockAccountService)
	router := setupRouter(accountController)

	req, _ := http.NewRequest(http.MethodPut, "/api/account-frontstore/account/", bytes.NewBuffer(marshalledSampleAccountBody))
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTIxNTU2Njd9.nT0WTSMOxfFIGLC5ABkJ8Nk-n2CuoZN-cW0FSK5W1CU")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"User with given email already exists\"", w.Body.String())
}
