package memcached

import (
	"storage/models"

	"github.com/thedivanov/memcached"
)

type Memcahced struct {
	conn *memcached.Conn
}

func NewMemcahced(MemcachedAddr string) (*Memcahced, error) {
	conn, err := memcached.NewConn(MemcachedAddr)
	if err != nil {
		return nil, err
	}
	return &Memcahced{
		conn: conn,
	}, nil
}

func (m *Memcahced) Set(key string, val string, expire int64) (models.StorageItem, error) {
	err := m.conn.Set(key, val, expire)
	if err != nil {
		return models.StorageItem{}, err
	}

	return models.StorageItem{
		Key:        key,
		Data:       val,
		Expiration: expire,
	}, nil
}

func (m *Memcahced) Get(key string) (models.StorageItem, bool, error) {
	item, err := m.conn.Get(key)
	if err != nil {
		return models.StorageItem{}, false, err
	}

	return models.StorageItem{
		Key:        item.Key,
		Data:       item.Value,
		Expiration: int64(item.Expiration),
	}, true, nil
}

func (m *Memcahced) Delete(key string) {
	m.conn.Delete(key)
}
