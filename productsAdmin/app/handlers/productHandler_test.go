package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
	"qwik.in/productsAdmin/service"
	"testing"
)

func setupRouter(handler ProductHandler) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	newRouter := r.Group("products/api")

	newRouter.POST("/", handler.AddProduct)
	newRouter.GET("/", handler.GetProduct)
	newRouter.PUT("/:id", handler.UpdateProduct)
	newRouter.DELETE("/:id", handler.DeleteProduct)
	newRouter.GET("/search", handler.SearchProduct)
	return r
}

func TestAddProductSuccess(t *testing.T) {
	//setting up data to be passed
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}

	//setting up mock
	mockService := new(service.MockProductService)
	mockService.On("CreateProduct", p).Return(nil)
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodPost, "/products/api/", bytes.NewBuffer(productData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Product Created\"}", w.Body.String())
}

func TestAddProductFailure(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}

	mockService := new(service.MockProductService)
	mockService.On("CreateProduct", p).Return(errors.New("Error creating product"))
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	req, _ := http.NewRequest(http.MethodPost, "/products/api/", bytes.NewBuffer(productData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"messaege\":\"Something went wrong\"}", w.Body.String())
}

func TestFindProductSuccess(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	products := []entity.Product{p}

	mockService := new(service.MockProductService)
	mockService.On("GetAll").Return(products, nil)
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	req, _ := http.NewRequest(http.MethodGet, "/products/api/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedProducts []entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &returnedProducts)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, returnedProducts)
	assert.Equal(t, 1, len(returnedProducts))
	assert.Equal(t, p, returnedProducts[0])
}

func TestFindProductFailure(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	products := []entity.Product{p}

	mockService := new(service.MockProductService)
	mockService.On("GetAll").Return(products, errors.New("Cannot find products"))
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	req, _ := http.NewRequest(http.MethodGet, "/products/api/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedProducts []entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &returnedProducts)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Nil(t, returnedProducts)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}

func TestUpdateProductSuccess(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}
	productId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockProductService)
	mockService.On("UpdateProduct", productId, p).Return(nil)
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodPut, "/products/api/"+productId, bytes.NewBuffer(productData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Product Updated Successfully\"}", w.Body.String())
}

func TestUpdateProductFailure(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}
	productId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockProductService)
	mockService.On("UpdateProduct", productId, p).Return(errors.New("Cannot update product"))
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodPut, "/products/api/"+productId, bytes.NewBuffer(productData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}

func TestDeleteProductSuccess(t *testing.T) {
	productId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockProductService)
	mockService.On("DeleteProduct", productId).Return(nil)
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodDelete, "/products/api/"+productId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Product Deleted Successfully\"}", w.Body.String())
}

func TestDeleteProductFailure(t *testing.T) {
	productId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockProductService)
	mockService.On("DeleteProduct", productId).Return(errors.New("Cannot delete product"))
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodDelete, "/products/api/"+productId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}

func TestSearchProductSuccess(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	products := []entity.Product{p}

	mockService := new(service.MockProductService)
	mockService.On("SearchProduct").Return(products, nil)
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	req, _ := http.NewRequest(http.MethodGet, "/products/api/search", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedProducts []entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &returnedProducts)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, returnedProducts)
	assert.Equal(t, 1, len(returnedProducts))
	assert.Equal(t, p, returnedProducts[0])
}

func TestSearchProductFailure(t *testing.T) {
	productData, _ := os.ReadFile("sampleCreateProduct.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	products := []entity.Product{p}

	mockService := new(service.MockProductService)
	mockService.On("SearchProduct").Return(products, errors.New("Error finding product"))
	productHandler := NewProductHandler(mockService)
	router := setupRouter(productHandler)

	req, _ := http.NewRequest(http.MethodGet, "/products/api/search", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedProducts []entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &returnedProducts)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}
