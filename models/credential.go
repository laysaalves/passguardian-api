package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Credential struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    ServiceName string    `json:"service-name"`
    User        string    `json:"user"`
    Password    string    `json:"password"`
    CreatedAt   time.Time `json:"created_at"`
}
