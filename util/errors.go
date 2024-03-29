package util

import "errors"

var (
	ErrNotFound            = errors.New("asset not found")
	ErrCreateFailed        = errors.New("asset created failed")
	ErrUpdateFailed        = errors.New("asset updated failed")
	ErrDeleteFailed        = errors.New("asset deleted failed")
	ErrInvalidObjectId     = errors.New("invalid objectId")
	ErrEmptyInput          = errors.New("empty input name or blockchain or amount or address")
	ErrNotConnecInDatabase = errors.New("not connected to database")
	ErrEmptyAssetList      = errors.New("asset list is empty")
)
