package common

import (
	schema "act/common/act"
	"act/common/cache"
	"act/common/models"
	"act/common/tools/snowflake"
	_ "github.com/go-sql-driver/mysql"

	"entgo.io/ent/dialect/sql"
)

type DbStore struct {
	*schema.Client
	Cache *cache.Redis
}

func NewDbStore(c models.DbStoreConfig) *DbStore {
	_, err := snowflake.NewSnowflake(c.DataCenterId, c.WorkerId)
	if err != nil {
		panic("connect database error")
	}
	db, err := sql.Open(c.Driver, c.DataSource)
	if err != nil {
		panic("connect database error")
	}
	// init cache driver
	drv, redis := cache.NewDriver(db, c.Cache)

	entClient := schema.NewClient(schema.Driver(drv))
	if c.Mode == models.DevMode {
		entClient = entClient.Debug()
	}
	return &DbStore{
		Client: entClient,
		Cache:  redis,
	}
}
