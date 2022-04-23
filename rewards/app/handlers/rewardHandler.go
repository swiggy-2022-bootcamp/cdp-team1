package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"qwik.in/rewards/entity"
	"qwik.in/rewards/log"
	"qwik.in/rewards/service"
)

type RewardHandler struct {
	rewardService service.RewardService
}

func NewRewardHandler(rewardService service.RewardService) RewardHandler {
	return RewardHandler{rewardService: rewardService}
}

// Searchreward godoc
// @Summary Search rewards
// @Description Search reward with given query
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p RewardHandler) Searchreward(c *gin.Context) {

	id := c.Param("id")
	log.Info("Find rewards with id : ", id)

	rewards, err := p.rewardService.SearchReward(id)
	if err == nil {
		c.JSON(http.StatusOK, rewards)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// Getall godoc
// @Summary Get Rewards
// @Description Get a list of all Rewards
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p RewardHandler) Getall(c *gin.Context) {
	rewards, err := p.rewardService.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, rewards)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// UpdateReward godoc
// @Summary Update Rewards
// @Description Update Reward with given id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [PUT]
func (p RewardHandler) UpdateReward(c *gin.Context) {

	RewardId := c.Param("id")

	var Reward entity.Reward
	if err := c.BindJSON(&Reward); err != nil {
		log.Error(err)
	}

	log.Info("Update Reward having id : ", RewardId, " with values: ", Reward)

	err := p.rewardService.UpdateReward(RewardId, Reward)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Reward Updated Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// DeleteReward godoc
// @Summary Delete Rewards
// @Description Delete Reward with given id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [DELETE]
func (p RewardHandler) DeleteReward(c *gin.Context) {

	RewardId := c.Param("id")
	log.Info("Delete Reward with id : ", RewardId)

	err := p.rewardService.DeleteReward(RewardId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Reward Deleted Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// AddReward godoc
// @Summary AddReward
// @Description Create a new Reward object, generate id and save in DB
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Reward Created
// @Router / [POST]
func (p RewardHandler) AddReward(c *gin.Context) {
	var Reward entity.Reward
	if err := c.BindJSON(&Reward); err != nil {
		log.Error(err)
	}

	log.Info("Add Reward with values ", Reward)

	err := p.rewardService.CreateReward(Reward)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"messaege": "Something went wrong"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Reward Created"})
	}
}
