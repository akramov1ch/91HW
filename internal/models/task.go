package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    Status      string             `bson:"status" json:"status"`
    CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
