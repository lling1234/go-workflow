package store

import (
	"act/common"
	"act/common/models"
)

type ActStore struct {
	*common.DbStore
}

func NewActrStore(c models.DbStoreConfig) *ActStore {
	return &ActStore{
		DbStore: common.NewDbStore(c),
	}
}
