package reposiroty

import (
	"cryptocurrencies-service/db"
	"cryptocurrencies-service/model"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const AssetCollection = "assets"

type AssetRepositoryInterface interface {
	Insert(asset *model.Asset) error
	Read(id string) (asset *model.Asset, err error)
	Delete(id string) error
	Update(asset *model.Asset) error
}

type assetRepository struct {
	c *mgo.Collection
}

func NewAssetRepository(conn db.Connection) AssetRepositoryInterface {
	return &assetRepository{conn.DB().C(AssetCollection)}
}

func (r *assetRepository) Insert(asset *model.Asset) error {
	fmt.Println("ASSET", asset)
	return r.c.Insert(asset)
}

func (r *assetRepository) Read(id string) (asset *model.Asset, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&asset)
	return asset, err
}

func (r *assetRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

func (r *assetRepository) Update(asset *model.Asset) error {
	return r.c.UpdateId(asset.Id, asset)
}
