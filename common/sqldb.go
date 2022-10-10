package common

import (
	"go-wflow/common/cache"
	"go-wflow/common/models"
	"go-wflow/common/utils/snowflake"
	schema "go-wflow/kernel/ent"

	_ "github.com/go-sql-driver/mysql"

	"github.com/qkbyte/ent/dialect/sql"
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
	// TODO cache redis
	// init cache driver
	// drv, redis := cache.NewDriver(dialect.Driver{}, c.Cache)

	entClient := schema.NewClient(schema.Driver(db))
	if c.Mode == models.DBMode {
		entClient = entClient.Debug()
	}
	if c.Mode == models.DevMode {
		// TODO开启软删除
		entClient = entClient.Debug().SoftDelete()
		// entClient = entClient.Debug()
	}
	return &DbStore{
		Client: entClient,
		// Cache:  redis,
	}
}
