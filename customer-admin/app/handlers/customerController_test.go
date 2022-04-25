package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/service/mocks"
	"qwik.in/customers-admin/internal/errors"
	"testing"
)

func setupRouter(customerController CustomerControllerInterface) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	newRouter := r.Group("api/customer-admin")
	newRouter.POST("/customer", customerController.CreateCustomer)
	newRouter.GET("/customer/:customerId", customerController.GetCustomerById)
	newRouter.GET("/customer/email/:customerEmail", customerController.GetCustomerByEmail)
	newRouter.PUT("/customer/:customerId", customerController.UpdateCustomer)
	newRouter.DELETE("/customer/:customerId", customerController.DeleteCustomer)

	newRouter.GET("/", HealthCheck)

	newRouter.GET("/user", GetAdminUsers)
	return r
}

func TestCreateCustomerForSuccess(t *testing.T) {
	sampleCustomerBody := model.Customer{
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("CreateCustomer", sampleCustomerBody).Return(&sampleCustomerBody, nil)

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodPost, "/api/customer-admin/customer", bytes.NewBuffer(marshalledSampleCustomerBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateCustomerForFailure(t *testing.T) {
	sampleCustomerBody := model.Customer{
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("CreateCustomer", sampleCustomerBody).Return(nil, errors.NewEmailAlreadyRegisteredError())

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodPost, "/api/customer-admin/customer", bytes.NewBuffer(marshalledSampleCustomerBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"User with given email already exists\"", w.Body.String())
}

func TestGetCustomerByIdForSuccess(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	//marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("GetCustomerById", sampleCustomerBody.CustomerId).Return(&sampleCustomerBody, nil)

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodGet, "/api/customer-admin/customer/abcdefg", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetCustomerByIdForFailure(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	//marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("GetCustomerById", sampleCustomerBody.CustomerId).Return(nil, errors.NewUserNotFoundError())

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodGet, "/api/customer-admin/customer/abcdefg", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"User not found\"", w.Body.String())
}

func TestGetCustomerByEmailForSuccess(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	//marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("GetCustomerByEmail", sampleCustomerBody.Email).Return(&sampleCustomerBody, nil)

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodGet, "/api/customer-admin/customer/email/ashwin@gmail.com", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetCustomerByEmailForFailure(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	//marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("GetCustomerByEmail", sampleCustomerBody.Email).Return(nil, errors.NewUserNotFoundError())

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodGet, "/api/customer-admin/customer/email/ashwin@gmail.com", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"User not found\"", w.Body.String())
}

func TestUpdateCustomerForSuccess(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("UpdateCustomer", sampleCustomerBody.CustomerId, sampleCustomerBody).Return(&sampleCustomerBody, nil)

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodPut, "/api/customer-admin/customer/abcdefg", bytes.NewBuffer(marshalledSampleCustomerBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateCustomerForFailure(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("UpdateCustomer", sampleCustomerBody.CustomerId, sampleCustomerBody).Return(nil, errors.NewUserNotFoundError())

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodPut, "/api/customer-admin/customer/abcdefg", bytes.NewBuffer(marshalledSampleCustomerBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"User not found\"", w.Body.String())
}

func TestUpdateCustomerForEmailFailure(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("UpdateCustomer", sampleCustomerBody.CustomerId, sampleCustomerBody).Return(nil, errors.NewEmailAlreadyRegisteredError())

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodPut, "/api/customer-admin/customer/abcdefg", bytes.NewBuffer(marshalledSampleCustomerBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"User with given email already exists\"", w.Body.String())
}

func TestDeleteCustomerForSuccess(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	//marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	str := "deletion successful"
	mockCustomerService.On("DeleteCustomer", sampleCustomerBody.CustomerId).Return(&str, nil)

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodDelete, "/api/customer-admin/customer/abcdefg", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"deletion successful\"", w.Body.String())
}

func TestDeleteCustomerForFailure(t *testing.T) {
	sampleCustomerBody := model.Customer{
		CustomerId:      "abcdefg",
		Firstname:       "Ashwin",
		Lastname:        "Nayar",
		Email:           "ashwin@gmail.com",
		Password:        "password",
		Telephone:       "1-541-754-3010",
		CustomerGroupId: 1,
		Address:         nil,
		Agree:           1,
	}
	//marshalledSampleCustomerBody, _ := json.Marshal(sampleCustomerBody)

	mockCustomerService := &mocks.CustomerServiceInterface{}
	mockCustomerService.On("DeleteCustomer", sampleCustomerBody.CustomerId).Return(nil, errors.NewUserNotFoundError())

	customerController := InitCustomerController(mockCustomerService)
	router := setupRouter(customerController)

	req, _ := http.NewRequest(http.MethodDelete, "/api/customer-admin/customer/abcdefg", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"User not found\"", w.Body.String())
}
