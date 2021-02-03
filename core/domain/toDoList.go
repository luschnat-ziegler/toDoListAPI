package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description *string            `json:"description" bson:"description"`
	Tasks       []Task             `json:"tasks,omitempty" bson:"tasks,omitempty" validate:"required,dive,required"`
}

type Task struct {
	Id          string  `json:"id" bson:"id"`
	Name        string  `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description *string `json:"description" bson:"description"`
}
