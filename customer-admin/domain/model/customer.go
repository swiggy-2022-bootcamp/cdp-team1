package model

import (
	"time"
)

type Customer struct {
	CustomerId      string `json:"customer_id"`
	Firstname       string `json:"first_name"`
	Lastname        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Telephone       string `json:"telephone"`
	CustomerGroupId int    `json:"customer_group_id"`
	//TBC
	Address []struct {
		HouseNumber string `json:"house_number"`
		Street      string `json:"street"`
		Area        string `json:"area"`
		City        string `json:"city"`
		Country     string `json:"country"`
		Pincode     string `json:"pincode"`
		Default     int    `json:"default"`
	} `json:"address"`
	Affiliate struct {
		Company           string `json:"company"`
		Website           string `json:"website"`
		Tracking          string `json:"tracking"`
		Commission        string `json:"commission"`
		Tax               string `json:"tax"`
		BankName          string `json:"bank_name"`
		BankBranchNumber  string `json:"bank_branch_number"`
		BankSwiftCode     string `json:"bank_swift_code"`
		BankAccountName   string `json:"bank_account_name"`
		BankAccountNumber string `json:"bank_account_number"`
		Status            int    `json:"status"`
	} `json:"affiliate"`
	Agree      int       `json:"agree"`
	DateAdded  time.Time `json:"date_added"`
	Approved   int       `json:"approved"`
	Safe       int       `json:"safe"`
	Newsletter int       `json:"newsletter"`
	Status     int       `json:"status"`
}
