package main

import (
	"act/common/act"
	ent "act/common/act"
	"act/common/act/procdef"
	"act/common/cache"
	"act/common/models"
	"act/common/store"
	"act/common/tools/linq"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func query() []*ent.ProcDef {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/wflow?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	defs, err := store.ProcDef.Query().Where(procdef.FormIDEQ("12393481")).Select(procdef.FieldVersion).All(ctx)
	if err != nil {
		log.Println(err)
	}
	return defs
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func linqfor() {
	maxVersion := linq.From(query()).SelectT(func(a *act.ProcDef) int32 {
		return a.Version
	}).Max()
	fmt.Println("maxVersion", maxVersion.(int32))
}

func yxqfor() {

	var maxVersion int32 = 0
	for _, v := range query() {
		if maxVersion < v.Version {
			maxVersion = v.Version
		}
	}
	fmt.Println("maxVersion", maxVersion)
}

// 基准测试
func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		/*
			BenchmarkFib20-16
			      51	  21740778 ns/op	   36516 B/op	     391 allocs/op
			PASS
		*/
		yxqfor()
		/*
			BenchmarkFib20-16
			      73	  16517734 ns/op	   37087 B/op	     415 allocs/op
			PASS
		*/
		// linqfor()
	}
}

// 基准测试

// 0.355s
func TestActDbCreate6(t *testing.T) {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/wflow?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	defs, err := store.ProcDef.Query().Where(procdef.FormIDEQ("12393481")).Select(procdef.FieldVersion).All(ctx)
	if err != nil {
		t.Log(err)
	}

	var maxVersion int32 = 0
	for _, v := range defs {
		if maxVersion < v.Version {
			maxVersion = v.Version
		}
	}
	fmt.Println("maxVersion", maxVersion)

}

// 0.182s
// -------最大版本--------
func TestActDbCreate5(t *testing.T) {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/wflow?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	defs, err := store.ProcDef.Query().Where(procdef.FormIDEQ("12393481")).Select(procdef.FieldVersion).All(ctx)
	if err != nil {
		t.Log(err)
	}
	maxVersion := linq.From(defs).SelectT(func(a *act.ProcDef) int32 {
		return a.Version
	}).Max()
	fmt.Println("maxVersion", maxVersion.(int32))
}

// ---------------

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit       int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset      int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SearchCount bool   `protobuf:"varint,3,opt,name=searchCount,proto3" json:"searchCount,omitempty"`
	Filter      string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}
type CommonRpcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//返回数据JSON
	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

// 查询分页
func PageResult(page *PageRequest, total int64, data interface{}, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	return JsonResult(&models.PageResponse{
		Records:     data,
		Current:     page.Offset/page.Limit + 1,
		Size:        page.Limit,
		SearchCount: page.SearchCount,
		Total:       total,
		Pages:       int64(math.Ceil(float64(total / page.Limit))),
	}, err)
}
func JsonResult(data interface{}, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	if data == nil {
		return &CommonRpcRes{}, errors.New("not found")
	}
	resJson, err := json.Marshal(data)
	return Result(string(resJson), err)
}
func Result(data string, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	return &CommonRpcRes{
		Json: data,
	}, nil
}
func TestActDbCreate4(t *testing.T) {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/wflow?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	var offset int = 0
	var limit int = 5
	var filter string = "假"

	query := store.ProcDef.Query()
	total, err := query.Count(ctx)
	if err != nil {
		log.Println("err", err)
	}
	procdefArr, err := query.
		Where(procdef.NameContains(filter)).
		Order(act.Asc(procdef.FieldID)).Offset(offset).Limit(limit).
		All(ctx)
	if err != nil {
		log.Println("procdefArr", err)
	}

	res, err := PageResult(&PageRequest{Limit: int64(limit), Offset: int64(offset)}, int64(total), procdefArr, err)
	if err != nil {
		log.Println("res", res)
	}

	t.Log("111111111111111")
	t.Log("res.Json", res.Json)
	t.Log("222222222222222")

}

func TestActDbCreate(t *testing.T) {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:ye199169@(127.0.0.1:3306)/act?charset=utf8&parseTime=true&loc=Local",
	})
	defer store.Close()
	ctx := context.Background()
	// // Run the auto migration tool create database.
	// if err := store.Schema.Create(ctx, schemax.WithForeignKeys(true)); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	store.ProcDef.Create().
		SetName("11").SetCode("11").SetVersion(1).SetResource("aaad打324").
		SetCreateUserID(111).SetCreateUserName("ling").SetCreateTime(time.Now()).
		SetTargetID(222).
		Save(ctx)
}

func TestAmain(t *testing.T) {
	client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/wflow?parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	fmt.Println("000")
	CreatProcDef(ctx, client)
	fmt.Println("111")
}

func CreatProcDef(ctx context.Context, client *ent.Client) (*ent.ProcDef, error) {
	p, err := client.ProcDef.Create().
		SetName("11").SetCode("11").SetVersion(1).SetResource("aaad打324").
		SetCreateUserID(111).SetCreateUserName("ling").SetCreateTime(time.Now()).
		SetTargetID(222).
		Save(ctx)
	fmt.Println("err", err)
	fmt.Println("p", p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
