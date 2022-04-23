package service

import "qwik.in/rewards/entity"

type RewardService interface {
	GetAll() ([]entity.Reward, error)
	DeleteReward(RewardId string) error
	SearchReward(query string) (entity.Reward, error)
	UpdateReward(RewardId string, Reward entity.Reward) error
	CreateReward(Reward entity.Reward) error
}
