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
	TX() ITXStore
}

type ITXStore interface {
	IStore
	Commit() error
	Rollback() error
}

type database struct {
	db *gorm.DB
}

// NewStore returns a new store.
func NewStore(db *gorm.DB) IStore {
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

func (ds *database) TX() ITXStore {
	ds.db.Begin()
	return &database{db: ds.db.Begin()}
}

func (ds *database) Commit() error {
	return ds.db.Commit().Error
}

func (ds *database) Rollback() error {
	return ds.db.Rollback().Error
}
