package model

import (
	"qwik.in/account-frontstore/protos"
	"time"
)

type Account struct {
	CustomerId      string                `json:"customer_id"`
	Firstname       string                `json:"first_name"`
	Lastname        string                `json:"last_name"`
	Email           string                `json:"email"`
	Password        string                `json:"password"`
	Telephone       string                `json:"telephone"`
	CustomerGroupId int                   `json:"customer_group_id"`
	Cart            []*protos.Product     `json:"cart"`
	RewardsTotal    int32                 `json:"rewards_total"`
	UserBalance     []*protos.PaymentMode `json:"user_balance"`
	Agree           int                   `json:"agree"`
	DateAdded       time.Time             `json:"date_added"`
}
