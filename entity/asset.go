package entity

import "gopkg.in/mgo.v2/bson"

type Asset struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Address    string        `json:"address" bson:"address"`
	Amount     float32       `json:"amount" bson:"amount"`
	Name       string        `json:"name" bson:"name"`
	Blockchain string        `json:"blockchain" bson:"blockchain"`
}
