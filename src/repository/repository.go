package repository

import (
	"storage/models"
)

type Repository interface {
	Get(key string) (models.StorageItem, bool, error)
	Set(key string, val string, expire int64) (models.StorageItem, error)
	Delete(key string)
}
