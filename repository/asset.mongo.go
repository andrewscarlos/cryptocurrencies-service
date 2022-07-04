package repository

import (
	"cryptocurrencies-service/db"
	"cryptocurrencies-service/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const AssetCollection = "assets"

type AssetRepository struct {
	c *mgo.Collection
}

func NewAssetRepository(conn db.Connection) AssetRepositoryInterface {
	return &AssetRepository{conn.DB().C(AssetCollection)}
}

func (r *AssetRepository) Insert(asset *entity.Asset) error {
	return r.c.Insert(asset)
}

func (r *AssetRepository) Read(id string) (asset *entity.Asset, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&asset)
	return asset, err
}

func (r *AssetRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

func (r *AssetRepository) Update(asset *entity.Asset) error {
	return r.c.UpdateId(asset.Id, asset)
}
