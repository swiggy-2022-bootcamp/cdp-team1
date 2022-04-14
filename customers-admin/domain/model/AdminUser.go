package model

import (
	"time"
)

type AdminUser struct {
	UserId      string    `json:"userId"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"userName"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Telephone   string    `json:"telephone"`
	UserGroupId int       `json:"user_group_id"`
	Status      int       `json:"status"`
	DateAdded   time.Time `json:"dateAdded"`
}
