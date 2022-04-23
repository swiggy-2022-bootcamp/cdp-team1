package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"qwik.in/rewards/entity"
	"qwik.in/rewards/repository"
)

func TestCreateRewardSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("SaveReward").Return(nil)
	rewardService := NewRewardService(mockRepo)

	reward := entity.Reward{}
	result := rewardService.CreateReward(reward)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, result)
}

func TestCreateRewardFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("SaveReward").Return(errors.New("Cannot save Reward"))
	RewardService := NewRewardService(mockRepo)

	Reward := entity.Reward{}
	result := RewardService.CreateReward(Reward)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result)
}

func TestGetAllSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	Reward := entity.Reward{}
	mockRepo.On("FindAll").Return([]entity.Reward{Reward}, nil)
	RewardService := NewRewardService(mockRepo)

	Rewards, err := RewardService.GetAll()

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, Rewards)
	assert.Equal(t, len(Rewards), 1)
}

func TestGetAllFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("FindAll").Return([]entity.Reward{}, errors.New("Cannot find Rewards"))
	RewardService := NewRewardService(mockRepo)

	Rewards, err := RewardService.GetAll()

	mockRepo.AssertExpectations(t)
	assert.Nil(t, Rewards)
	assert.NotNil(t, err)
}

func TestUpdateRewardSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	Reward := entity.Reward{}
	mockRepo.On("SaveReward").Return(nil)
	RewardService := NewRewardService(mockRepo)

	err := RewardService.UpdateReward(entity.NewID(), Reward)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestUpdateRewardFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	Reward := entity.Reward{}
	mockRepo.On("SaveReward").Return(errors.New("Cannot save Reward"))
	RewardService := NewRewardService(mockRepo)

	err := RewardService.UpdateReward(entity.NewID(), Reward)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestDeleteRewardSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("DeleteReward").Return(nil)
	RewardService := NewRewardService(mockRepo)

	err := RewardService.DeleteReward(entity.NewID())

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDeleteRewardFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("DeleteReward").Return(errors.New("Cannot delete Reward"))
	RewardService := NewRewardService(mockRepo)

	err := RewardService.DeleteReward(entity.NewID())

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestSearchRewardSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	Reward := entity.Reward{}
	mockRepo.On("FindWithLimit").Return([]entity.Reward{Reward}, nil)
	RewardService := NewRewardService(mockRepo)

	Rewards, err := RewardService.SearchReward(1)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, Rewards)
	assert.Equal(t, len(Rewards), 1)
}

func TestSearchRewardFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("FindWithLimit").Return([]entity.Reward{}, errors.New("Cannot find Rewards"))
	RewardService := NewRewardService(mockRepo)

	Rewards, err := RewardService.SearchReward(1)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, Rewards)
	assert.NotNil(t, err)
}
