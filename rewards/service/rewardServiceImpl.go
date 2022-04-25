package service

import (
	"qwik.in/rewards/entity"
	"qwik.in/rewards/log"

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
	if Reward.ID == "" {
		Reward.SetId()
	}
	log.Info(Reward)
	if err := r.rewardRepository.SaveReward(Reward); err != nil {
		return err
	}
	return nil
}

// func (r RewardServiceImpl) GetRewardPoints(userId string) (*proto.Response, error) {
// 	log.Info("Connecting with gRPC server")
// 	// Set up a connection to the server.
// 	serverAddress := fmt.Sprintf("%s:%s", config.GRPC_SERVER_IP, config.GRPC_SERVER_PORT)
// 	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Error("did not connect: ", err)
// 		return nil, err
// 	}

// 	//close
// 	defer func(conn *grpc.ClientConn) {
// 		err := conn.Close()
// 		if err != nil {
// 			log.Error("Connection closed with error", err.Error())
// 		}
// 	}(conn)
// 	c := proto.NewRewardPointServiceClient(conn)

// 	// Disconnect gRPC call upon
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	// Send gRPC request to server the
// 	r2, err := c.GetRewardPoints(ctx, &proto.Request{Id: userId})
// 	if err != nil {
// 		log.Error("could not get response: ", err)
// 		return nil, err
// 	}
// 	log.Info("Response from gRPC server", r)
// 	one, err := r.rewardRepository.FindOne(r2.id)
// 	r1 := entity.Reward{ID: r2.id, Points: r2.transaction_amount}
// 	if err != nil {
// 		r.rewardRepository.SaveReward(r1)
// 	} else {
// 		r1.Points = r1.Points + r2.transaction_amount
// 		r.rewardRepository.SaveReward(r1)
// 	}
// 	return one, nil

// 	log.Info("gRPC received id: ", r.id)
// 	return r, nil
// }
