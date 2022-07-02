package entity

import "gopkg.in/mgo.v2/bson"

type AssetInterface interface {
	GetId() bson.ObjectId
	GetAddress() string
	GetValue() float32
	GetName() string
	GetBlockchain() string
}

type Asset struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Address    string        `json:"address" bson:"address"`
	Value      float32       `json:"value" bson:"value"`
	Name       string        `json:"name" bson:"name"`
	Blockchain string        `json:"blockchain" bson:"blockchain"`
}

func (a *Asset) GetId() bson.ObjectId {
	return a.Id
}
func (a *Asset) GetAddress() string {
	return a.Address
}

func (a *Asset) GetValue() float32 {
	return a.Value
}
func (a *Asset) GetName() string {
	return a.Name
}

func (a *Asset) GetBlockchain() string {
	return a.Blockchain
}
