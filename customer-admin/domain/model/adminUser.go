package model

import (
	"time"
)

type AdminUser struct {
	UserId      string    `json:"user_id"`
	Firstname   string    `json:"first_name"`
	Lastname    string    `json:"last_name"`
	Username    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Telephone   string    `json:"telephone"`
	UserGroupId int       `json:"user_group_id"`
	UserGroup   string    `json:"user_group"`
	Status      int       `json:"status"`
	DateAdded   time.Time `json:"date_added"`
}
