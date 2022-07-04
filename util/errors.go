package util

import "errors"

var (
	ErrNotFound        = errors.New("asset not found")
	ErrCreateFailed    = errors.New("asset created failed")
	ErrUpdateFailed    = errors.New("asset updated failed")
	ErrDeleteFailed    = errors.New("asset deleted failed")
	ErrUnauthorized    = errors.New("Unauthorized")
	ErrInvalidObjectId = errors.New("invalid objectId")
)
