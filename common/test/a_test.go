package main

import (
	ent "act/common/act"
	"act/common/act/procdef"
	"act/common/models"
	"act/common/store"
	"act/common/tools/linq"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
	client, err := ent.Open("mysql", "root:ye199169@tcp(127.0.0.1:3306)/act?parseTime=True&loc=Local")
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

func TestMax(t *testing.T) {
	client, err := ent.Open("mysql", "root:ye199169@tcp(127.0.0.1:3306)/act?parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	defs, err := client.ProcDef.Query().Where(procdef.FormIDEQ("0a270a2faaea4881bb3792dcb758d7bc")).Select(procdef.FieldVersion).All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	max := linq.From(defs).Max()
	log.Println("max", max)
}
