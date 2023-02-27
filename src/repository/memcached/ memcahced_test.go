package memcached

import (
	"storage/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert := assert.New(t)
	memcahced, err := NewMemcahced("localhost:11211")
	if err != nil {
		t.Error(err)
	}

	memcahced.Set("test-key-1", "test-data-1", 0)
	memcahced.Set("test-key-2", "test-data-2", 0)
	memcahced.Set("test-key-3", "test-data-3", 1)

	item1, found1, err := memcahced.Get("test-key-1")
	assert.Equal(nil, err)
	assert.Equal(true, found1)
	assert.Equal(models.StorageItem{Key: "test-key-1", Data: "test-data-1"}, item1)

	item2, found2, err := memcahced.Get("test-key-2")
	assert.Equal(nil, err)
	assert.Equal(true, found2)
	assert.Equal(models.StorageItem{Key: "test-key-2", Data: "test-data-2"}, item2)

	item3, found3, err := memcahced.Get("test-key-4")
	assert.Error(err, "Memcached response corrupt")
	assert.Equal(false, found3)
	assert.Empty(item3)
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	memcahced, err := NewMemcahced("localhost:11211")
	if err != nil {
		t.Error(err)
	}

	memcahced.Set("test-key-1", "test-data-1", 0)
	memcahced.Set("test-key-2", "test-data-2", 0)

	item1, found1, err := memcahced.Get("test-key-1")
	assert.Equal(nil, err)
	assert.Equal(true, found1)
	assert.Equal(models.StorageItem{Key: "test-key-1", Data: "test-data-1"}, item1)

	item2, found2, err := memcahced.Get("test-key-2")
	assert.Equal(nil, err)
	assert.Equal(true, found2)
	assert.Equal(models.StorageItem{Key: "test-key-2", Data: "test-data-2"}, item2)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	memcahced, err := NewMemcahced("localhost:11211")
	if err != nil {
		t.Error(err)
	}

	memcahced.Set("test-key-1", "test-data-1", 0)

	memcahced.Delete("test-key-1")

	item, found, err := memcahced.Get("test-key-1")
	assert.Error(err, "Not found")
	assert.Equal(false, found)
	assert.Empty(item)
}
