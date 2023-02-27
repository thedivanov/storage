package internalstorage

import (
	"fmt"
	"storage/models"
	"sync"
	"time"
)

type InternalStorage struct {
	mtx   sync.Mutex
	items map[string]models.StorageItem
}

func NewInternalStorage() *InternalStorage {
	return &InternalStorage{
		items: make(map[string]models.StorageItem),
	}
}

func (is *InternalStorage) Get(key string) (models.StorageItem, bool, error) {
	is.mtx.Lock()
	item, found := is.items[key]
	is.mtx.Unlock()
	if !found {
		return models.StorageItem{}, false, nil
	}
	if item.Expiration > 0 {
		if time.Now().Unix() > item.Expiration {
			is.mtx.Lock()
			delete(is.items, key)
			is.mtx.Unlock()
			return models.StorageItem{}, false, fmt.Errorf("Variable is expired")
		}
	}
	item.Key = key
	return item, true, nil
}

func (is *InternalStorage) Set(key string, val string, expire int64) (models.StorageItem, error) {
	var exp int64

	if expire > 0 {
		exp = expire
	}
	is.mtx.Lock()
	is.items[key] = models.StorageItem{
		Data:       val,
		Expiration: exp,
	}
	is.mtx.Unlock()

	return models.StorageItem{Key: key, Data: val, Expiration: exp}, nil
}

func (is *InternalStorage) Delete(key string) {
	is.mtx.Lock()
	delete(is.items, key)
	is.mtx.Unlock()
}
