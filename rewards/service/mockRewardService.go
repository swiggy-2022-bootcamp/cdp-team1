package service

import (
	"github.com/stretchr/testify/mock"
	"qwik.in/rewards/entity"
)

type MockRewardService struct {
	mock.Mock
}

func NewMockRewardService() RewardService {
	return &MockRewardService{}
}

func (m MockRewardService) GetAll() ([]entity.Reward, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]entity.Reward), args.Error(1)
}

func (m MockRewardService) UpdateReward(RewardId string, Reward entity.Reward) error {
	args := m.Called(RewardId, Reward)
	return args.Error(0)
}

func (m MockRewardService) DeleteReward(RewardId string) error {
	args := m.Called(RewardId)
	return args.Error(0)
}

func (m MockRewardService) SearchReward(limit string) (entity.Reward, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(entity.Reward), args.Error(1)
}
func (m MockRewardService) CreateReward(Reward entity.Reward) error {
	args := m.Called(Reward)
	return args.Error(0)
}
