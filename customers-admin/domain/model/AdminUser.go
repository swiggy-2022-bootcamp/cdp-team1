package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AdminUser struct {
	UserId      primitive.ObjectID `json:"userId"`
	Firstname   string             `json:"firstname"`
	Lastname    string             `json:"lastname"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	Telephone   string             `json:"telephone"`
	UserGroupId int                `json:"user_group_id"`
	Status      int                `json:"status"`
	DateAdded   time.Time          `json:"dateAdded"`
}
