package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	fmt.Println("Reward Id", rewards.ID)
	fmt.Println("Reward Name", rewards.Name)
	fmt.Println("Reward Description", rewards.Description)
	fmt.Println("Reward Points", rewards.Points)
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
