package repository

import "qwik.in/rewards/entity"

type RewardRepository interface {
	Connect() error
	FindOne(rewardId string) (entity.Reward, error)
	FindAll() ([]entity.Reward, error)
	DeleteReward(rewardId string) error
	SaveReward(Reward entity.Reward) error
	// FindAndUpdate(rewardId string, reward entity.Reward) error
}
