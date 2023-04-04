package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	S    *database
)

type IStore interface {
	Resources() ResourceStore
	Menus() MenuStore
	CategoryTags() CategoryTagStore
}

type database struct {
	db *gorm.DB
}

var _ IStore = (*database)(nil)

// NewStore returns a new store.
func NewStore(db *gorm.DB) *database {
	once.Do(func() {
		S = &database{db: db}
	})
	return S
}

func (ds *database) Resources() ResourceStore {
	return newResources(ds.db)
}

func (ds *database) Menus() MenuStore {
	return newMenus(ds.db)
}

func (ds *database) CategoryTags() CategoryTagStore {
	return newCategoryTags(ds.db)
}
