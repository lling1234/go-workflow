package main

import (
	"act/common/act"
	ent "act/common/act"
	"act/common/act/procdef"
	"act/common/act/procinst"
	"act/common/cache"
	"act/common/models"
	"act/common/store"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/random"
	_ "github.com/go-sql-driver/mysql"
)

// 查询版本最大 select max(version) from act_proc_def where form_id=''
func TestActDbCreate6(t *testing.T) {
	store := store.NewActrStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/workflowdemo?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	procdefList, err := store.ProcDef.Query().Where(procdef.CreateUserName("ling")).All(ctx)
	if err != nil {
		t.Log("procdefList err")
	}
	log.Println("procdefList", procdefList)
}
func TestActDbCreate5(t *testing.T) {
	store := store.NewActrStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/workflowdemo?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	procdefList, err := store.ProcInst.Query().Where(procinst.ID(826790988067274752), procinst.ProcDefID(800650655701041152)).First(ctx)
	if err != nil {
		t.Log("procdefList err")
	}
	log.Println("procdefList", procdefList)
	// 6.2 更新
	procdefUpdate, err := procdefList.Update().SetIsFinished(2).SetNodeID("1659320194369").SetTaskID(9212524636351135751).Save(ctx)
	if err != nil {
		t.Log("procdefUpdate err")
	}
	log.Println("procdefUpdate", procdefUpdate)
}
func TestActDbCreate4(t *testing.T) {
	store := store.NewActrStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/workflowdemo?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()
	p, err := store.Execution.Create().SetCreateTime(time.Now()).
		SetProcInstID(int64(random.RandInt(1, 100000000000))).
		SetProcDefID(8725662109012164608).SetIsActive(1).
		SetStartTime(time.Now()).SetNodeInfos("resourceStr").Save(ctx)
	if err != nil {
		t.Log("Execution.Create() err")
	}
	t.Log("Execution", p)
}
func TestActDbCreate3(t *testing.T) {
	store := store.NewActrStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/workflowdemo?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	p, err := store.ProcInst.Create().SetCreateTime(time.Now()).SetProcDefID(11).SetTitle("111").
		SetNodeID("111").SetTaskID(11).SetStartTime(time.Now()).SetEndTime(time.Now().AddDate(0, 0, 1)).SetStartUserID(12025).SetStartUserName("ling").
		SetTargetID(1727882).SetDataID(0).SetState(1).Save(ctx)
	if err != nil {
		t.Log("ProcInst.Create() err")
	}
	t.Log("p", p)
}
func TestActDbCreate2(t *testing.T) {
	store := store.NewActrStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/workflowdemo?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
	})
	defer store.Close()
	ctx := context.Background()

	procDefResult := store.ProcDef.Query().
		Where(procdef.YewuFormID("1239348"), procdef.TargetIDEQ(1727882)).
		Order(act.Asc(procdef.FieldVersion))

	procdefIdArr, err := procDefResult.IDs(ctx)
	if err != nil {
		t.Log("procdefIdArr")
	}
	t.Log("procdefIdArr", procdefIdArr[0])
}

func TestActDbCreate(t *testing.T) {
	store := store.NewActrStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/workflowdemo?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr:     "qkbyte.orginone.cn:6002",
			Password: "orginone",
		},
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
	client, err := ent.Open("mysql", "root:123456@tcp(localhost:3306)/workflowDemo?parseTime=True")
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
