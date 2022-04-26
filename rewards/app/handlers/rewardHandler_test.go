package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"qwik.in/rewards/entity"
	"qwik.in/rewards/log"
	"qwik.in/rewards/service"
)

func setupRouter(handler RewardHandler) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	newRouter := r.Group("rewards/api")

	newRouter.GET("/search/:id", handler.Searchreward)
	newRouter.GET("/", handler.Getall)
	newRouter.PUT("/:id", handler.UpdateReward)
	newRouter.DELETE("/:id", handler.DeleteReward)
	newRouter.POST("/", handler.AddReward)
	return r
}

func TestAddRewardSuccess(t *testing.T) {
	//setting up data to be passed
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}

	//setting up mock
	mockService := new(service.MockRewardService)
	mockService.On("CreateReward", p).Return(nil)
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodPost, "/rewards/api/", bytes.NewBuffer(RewardData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Reward Created\"}", w.Body.String())
}

func TestAddRewardFailure(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}

	mockService := new(service.MockRewardService)
	mockService.On("CreateReward", p).Return(errors.New("Error creating Reward"))
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	req, _ := http.NewRequest(http.MethodPost, "/rewards/api/", bytes.NewBuffer(RewardData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"messaege\":\"Something went wrong\"}", w.Body.String())
}

func TestFindRewardSuccess(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	Rewards := []entity.Reward{p}

	mockService := new(service.MockRewardService)
	mockService.On("GetAll").Return(Rewards, nil)
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	req, _ := http.NewRequest(http.MethodGet, "/rewards/api/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedRewards []entity.Reward
	err := json.Unmarshal(w.Body.Bytes(), &returnedRewards)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, returnedRewards)
	assert.Equal(t, 1, len(returnedRewards))
	assert.Equal(t, p, returnedRewards[0])
}

func TestFindRewardFailure(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	Rewards := []entity.Reward{p}

	mockService := new(service.MockRewardService)
	mockService.On("GetAll").Return(Rewards, errors.New("Cannot find Rewards"))
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	req, _ := http.NewRequest(http.MethodGet, "/rewards/api/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedRewards []entity.Reward
	err := json.Unmarshal(w.Body.Bytes(), &returnedRewards)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Nil(t, returnedRewards)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}

func TestUpdateRewardSuccess(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}
	RewardId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockRewardService)
	mockService.On("UpdateReward", RewardId, p).Return(nil)
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodPut, "/rewards/api/"+RewardId, bytes.NewBuffer(RewardData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Reward Updated Successfully\"}", w.Body.String())
}

func TestUpdateRewardFailure(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}
	RewardId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockRewardService)
	mockService.On("UpdateReward", RewardId, p).Return(errors.New("Cannot update Reward"))
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodPut, "/rewards/api/"+RewardId, bytes.NewBuffer(RewardData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}

func TestDeleteRewardSuccess(t *testing.T) {
	RewardId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockRewardService)
	mockService.On("DeleteReward", RewardId).Return(nil)
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodDelete, "/rewards/api/"+RewardId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Reward Deleted Successfully\"}", w.Body.String())
}

func TestDeleteRewardFailure(t *testing.T) {
	RewardId := "c2ae2173-b1e2-4c41-bfc5-10deb7ae3e81"

	mockService := new(service.MockRewardService)
	mockService.On("DeleteReward", RewardId).Return(errors.New("Cannot delete Reward"))
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	//making http call
	req, _ := http.NewRequest(http.MethodDelete, "/rewards/api/"+RewardId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//asserts
	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}

func TestSearchRewardSuccess(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	Rewards := []entity.Reward{p}

	mockService := new(service.MockRewardService)
	mockService.On("SearchReward").Return(Rewards, nil)
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	req, _ := http.NewRequest(http.MethodGet, "/rewards/api/search", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedRewards []entity.Reward
	err := json.Unmarshal(w.Body.Bytes(), &returnedRewards)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, returnedRewards)
	assert.Equal(t, 1, len(returnedRewards))
	assert.Equal(t, p, returnedRewards[0])
}

func TestSearchRewardFailure(t *testing.T) {
	RewardData, _ := os.ReadFile("sampleCreateReward.json")
	var p entity.Reward
	if err := json.Unmarshal(RewardData, &p); err != nil {
		log.Error(err)
	}
	p.SetId()
	Rewards := []entity.Reward{p}

	mockService := new(service.MockRewardService)
	mockService.On("SearchReward").Return(Rewards, errors.New("Error finding Reward"))
	RewardHandler := NewRewardHandler(mockService)
	router := setupRouter(RewardHandler)

	req, _ := http.NewRequest(http.MethodGet, "/rewards/api/search", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var returnedRewards []entity.Reward
	err := json.Unmarshal(w.Body.Bytes(), &returnedRewards)
	if err != nil {
		log.Error(err.Error())
	}

	mockService.AssertExpectations(t)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Something went wrong\"}", w.Body.String())
}
