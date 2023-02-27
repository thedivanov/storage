package internalstorage

import (
	"storage/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert := assert.New(t)
	internalstorage := NewInternalStorage()

	internalstorage.Set("test-key-1", "test-data-1", 0)
	internalstorage.Set("test-key-2", "test-data-2", 0)
	internalstorage.Set("test-key-3", "test-data-3", 1)

	item1, found1, err := internalstorage.Get("test-key-1")
	assert.Equal(nil, err)
	assert.Equal(true, found1)
	assert.Equal(models.StorageItem{Key: "test-key-1", Data: "test-data-1"}, item1)

	item2, found2, err := internalstorage.Get("test-key-2")
	assert.Equal(nil, err)
	assert.Equal(true, found2)
	assert.Equal(models.StorageItem{Key: "test-key-2", Data: "test-data-2"}, item2)

	item3, found3, err := internalstorage.Get("test-key-3")
	assert.Error(err, "Variable is expired")
	assert.Equal(found3, false)
	assert.Empty(item3)

	item4, found4, err := internalstorage.Get("test-key-4")
	assert.Equal(nil, err)
	assert.Equal(false, found4)
	assert.Empty(item4)
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	internalstorage := NewInternalStorage()

	internalstorage.Set("test-key-1", "test-data-1", 0)
	internalstorage.Set("test-key-2", "test-data-2", 0)

	item1, found1, err := internalstorage.Get("test-key-1")
	assert.Equal(nil, err)
	assert.Equal(true, found1)
	assert.Equal(models.StorageItem{Key: "test-key-1", Data: "test-data-1"}, item1)

	item2, found2, err := internalstorage.Get("test-key-2")
	assert.Equal(nil, err)
	assert.Equal(true, found2)
	assert.Equal(models.StorageItem{Key: "test-key-2", Data: "test-data-2"}, item2)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	internalstorage := NewInternalStorage()

	internalstorage.Set("test-key-1", "test-data-1", 0)

	internalstorage.Delete("test-key-1")

	item, found, err := internalstorage.Get("test-key-1")
	assert.Equal(nil, err)
	assert.Equal(false, found)
	assert.Empty(item)
}
