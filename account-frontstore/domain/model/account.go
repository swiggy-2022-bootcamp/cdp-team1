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
	/*
		//TBC:
		Cart         Cart `json:"cart"`
		RewardsTotal float64
		UserBalance  PaymentsModeDTO
	*/
	Agree     int       `json:"agree"`
	DateAdded time.Time `json:"date_added"`
}
