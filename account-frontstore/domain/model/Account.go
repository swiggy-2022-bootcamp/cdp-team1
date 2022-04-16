package model

import (
	"time"
)

type Account struct {
	CustomerId      string `json:"customerId"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
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
	DateAdded time.Time `json:"dateAdded"`
}
