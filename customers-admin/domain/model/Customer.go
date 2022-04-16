package model

import (
	"time"
)

type Customer struct {
	CustomerId      string `json:"customerId"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Telephone       string `json:"telephone"`
	CustomerGroupId int    `json:"customer_group_id"`
	//TBC
	Address []struct {
		HouseNumber string `json:"houseNumber"`
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
		BankName          string `json:"bankName"`
		BankBranchNumber  string `json:"bankBranchNumber"`
		BankSwiftCode     string `json:"bankSwiftCode"`
		BankAccountName   string `json:"bankAccountName"`
		BankAccountNumber string `json:"bankAccountNumber"`
		Status            int    `json:"status"`
	} `json:"affiliate"`
	Agree      int       `json:"agree"`
	DateAdded  time.Time `json:"dateAdded"`
	Approved   int       `json:"approved"`
	Safe       int       `json:"safe"`
	Newsletter int       `json:"newsletter"`
	Status     int       `json:"status"`
}
