package repository

import "cryptocurrencies-service/entity"

type AssetRepositoryInterface interface {
	Insert(asset *entity.Asset) error
	Read(id string) (asset *entity.Asset, err error)
	Delete(id string) error
	Update(asset *entity.Asset) error
	GetAll() ([]*entity.Asset, error)
}

type AssetRepositoryAdapter struct {
	Persister AssetRepositoryInterface
}

func NewAssetRepositoryAdapter(persister AssetRepositoryInterface) AssetRepositoryInterface {
	return &AssetRepositoryAdapter{persister}
}

func (assetAdaper *AssetRepositoryAdapter) Insert(asset *entity.Asset) error {
	return assetAdaper.Persister.Insert(asset)
}

func (assetAdaper *AssetRepositoryAdapter) Read(id string) (asset *entity.Asset, err error) {
	return assetAdaper.Persister.Read(id)
}

func (assetAdaper *AssetRepositoryAdapter) Delete(id string) error {
	return assetAdaper.Persister.Delete(id)
}
func (assetAdaper *AssetRepositoryAdapter) Update(asset *entity.Asset) error {
	return assetAdaper.Persister.Update(asset)
}

func (assetAdaper *AssetRepositoryAdapter) GetAll() ([]*entity.Asset, error) {
	return assetAdaper.Persister.GetAll()
}
