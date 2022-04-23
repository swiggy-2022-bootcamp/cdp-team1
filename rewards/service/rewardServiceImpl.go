package service

import (
	"qwik.in/rewards/entity"
	"qwik.in/rewards/repository"
)

type RewardServiceImpl struct {
	rewardRepository repository.RewardRepository
}

func NewRewardService(rewardRepository repository.RewardRepository) RewardService {
	return &RewardServiceImpl{rewardRepository: rewardRepository}
}

func (r RewardServiceImpl) GetAll() ([]entity.Reward, error) {
	all, err := r.rewardRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (r RewardServiceImpl) DeleteReward(rewardId string) error {
	err := r.rewardRepository.DeleteReward(rewardId)
	if err != nil {
		return err
	}
	return nil
}

func (p RewardServiceImpl) SearchReward(query string) (entity.Reward, error) {
	one, err := p.rewardRepository.FindOne(query)
	r := entity.Reward{}
	if err != nil {
		return r, err
	}
	return one, nil
}
func (r RewardServiceImpl) UpdateReward(RewardId string, Reward entity.Reward) error {
	Reward.ID = RewardId
	if err := r.rewardRepository.SaveReward(Reward); err != nil {
		return err
	}
	return nil
}
func (r RewardServiceImpl) CreateReward(Reward entity.Reward) error {
	Reward.SetId()
	if err := r.rewardRepository.SaveReward(Reward); err != nil {
		return err
	}
	return nil
}
