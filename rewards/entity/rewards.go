package entity

type Reward struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}

func NewRewards(id string, name string, description string, points int) (*Reward,error) {
	reward := &Reward{
		ID:          id,
		Name:        name,
		Description: description,
		Points:      points,
	}
	return reward,nil
}
