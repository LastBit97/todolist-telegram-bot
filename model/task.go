package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"content" bson:"content" binding:"required"`
	ChatId int64              `json:"chat_id" bson:"chat_id" binding:"required"`
}
