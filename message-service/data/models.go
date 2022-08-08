package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var client *mongo.Client

// func New(mongo *mongo.Client) Models {
// 	client = mongo

// 	return Models{
// 		Msg: Msg{},
// 	}
// }

type Models struct {
	Msg Msg
}

type Msg struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Data      string             `bson:"data" json:"data"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
