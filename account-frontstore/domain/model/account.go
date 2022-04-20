package model

import (
	"time"
)

type Account struct {
	CustomerId      string `json:"customer_id"`
	Firstname       string `json:"first_name"`
	Lastname        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Telephone       string `json:"telephone"`
	CustomerGroupId int    `json:"customer_group_id"`
	//Cart         Cart `json:"cart"`
	RewardsTotal int32
	//UserBalance  PaymentsModeDTO
	Agree     int       `json:"agree"`
	DateAdded time.Time `json:"date_added"`
}
