package repository

import (
	"github.com/stretchr/testify/mock"
	"qwik.in/rewards/entity"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Connect() error {
	args := mock.Called()
	return args.Error(0)

}

func (mock *MockRepository) FindOne(RewardId string) (entity.Reward, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Reward), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Reward, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Reward), args.Error(1)
}

func (mock *MockRepository) SaveReward(Reward entity.Reward) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) DeleteReward(RewardId string) error {
	args := mock.Called()
	return args.Error(0)
}

