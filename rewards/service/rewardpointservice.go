package service

import (
	"context"

	"qwik.in/rewards/entity"
	"qwik.in/rewards/log"
	"qwik.in/rewards/proto"
	"qwik.in/rewards/repository"
)

type RewardPointService struct {
	proto.UnimplementedRewardPointServiceServer
}

var (
	Repository repository.RewardRepository
	Service    RewardService
)

func NewRewardPointService() *RewardPointService {
	Repository = repository.NewDynamoRepository()
	err := Repository.Connect()
	if err != nil {
		log.Error("Error while Connecting to DB: ", err)
		return nil
	}
	Service = NewRewardService(Repository)
	return &RewardPointService{}
}

func (s *RewardPointService) GetRewardPoints(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	log.Info("gRPC received message: ", in.GetId())
	reward, err1 := Service.SearchReward(in.GetId())
	if err1 != nil {
		log.Error("Error while searching reward: ", err1)
		return nil, err1
	}
	if reward.Points == 0 {
		reward.Points = int(in.GetTransactionAmount())
		var r1 entity.Reward
		r1.ID = in.GetId()
		r1.Points = int(in.GetTransactionAmount())
		Service.CreateReward(r1)
	} else {
		reward.Points = reward.Points + int(in.GetTransactionAmount())
		Service.UpdateReward(in.GetId(), reward)
	}
	return &proto.Response{
		Id:           in.GetId(),
		RewardPoints: int32(reward.Points),
	}, nil
}
