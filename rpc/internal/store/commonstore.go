package store

import (
	"act/common"
	"act/common/models"
)

type CommonStore struct {
	*common.DbStore
}

func NewCommonStore(c models.DbStoreConfig) *CommonStore {
	return &CommonStore{
		DbStore: common.NewDbStore(c),
	}
}
