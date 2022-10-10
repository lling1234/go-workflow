package test

import (
	"context"
	"encoding/json"
	"fmt"
	"go-wflow/common/models"
	"go-wflow/common/store"
	"go-wflow/common/utils/snowflake"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow"
	"io/ioutil"
	"log"
	"reflect"
	"testing"

	"github.com/duke-git/lancet/v2/netutil"
	_ "github.com/go-sql-driver/mysql"
)

// StructSliceToMapSlice : struct切片转为map切片
func StructSliceToMap(source interface{}) map[string]string {
	// 判断，interface转为[]interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		panic("ERROR: Unknown type, slice expected.")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	fmt.Println("ret", ret)
	fmt.Println("ret[0]", ret[0])
	a := ret[0].(Headers)
	fmt.Println("a", a)
	// 转换之后的结果变量
	data := make(map[string]string)

	// 通过遍历，每次迭代将struct转为map
	for i := 0; i < len(ret); i++ {
		data[ret[i].(Headers).Name] = ret[i].(Headers).Val
	}

	fmt.Printf("==== Convert Result ====\n%+v\n", data)
	return data
}

type WEBHook struct {
	Type    string                 `json:"type"`
	Method  string                 `json:"method"`
	URL     string                 `json:"url"`
	Headers []Headers              `json:"headers"`
	Body    map[string]interface{} `json:"params"`
}
type Headers struct {
	Name string `json:"name"`
	Val  string `json:"val"`
}
type Body struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestTriggerGet(t *testing.T) {
	// 1.slice转map

	headerSlice := make([]Headers, 0)
	headerSlice = append(headerSlice, Headers{Name: "Content-Type", Val: "application/json; charset=utf-8"})
	headerSlice = append(headerSlice, Headers{Name: "cookie", Val: "uuid=d1f0ecf1-26df-41e6-9e38-1e2f83a6fb79"})
	headerSlice = append(headerSlice, Headers{Name: "Authorization", Val: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXV"})

	header := StructSliceToMap(headerSlice)
	t.Log("header", header)

	reqBody := make(map[string]interface{})
	reqBody["id"] = 11
	reqBody["name"] = "ling"
	reqBody["age"] = 22
	bodyInfo, _ := json.Marshal(reqBody)
	// 2.发送http请求
	url := "http://httpbin.org/get"

	resp, err := netutil.HttpGet(url, header, nil, bodyInfo)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}
func TestTriggerPost(t *testing.T) {
	// 1.slice转map

	headerSlice := make([]Headers, 0)
	headerSlice = append(headerSlice, Headers{Name: "Content-Type", Val: "application/json"})
	headerSlice = append(headerSlice, Headers{Name: "cookie", Val: "uuid=d1f0ecf1-26df-41e6-9e38-1e2f83a6fb79"})
	headerSlice = append(headerSlice, Headers{Name: "Authorization", Val: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXV"})

	header := StructSliceToMap(headerSlice)
	t.Log("header", header)

	reqBody := make(map[string]interface{})
	reqBody["id"] = 11
	reqBody["name"] = "ling"
	reqBody["age"] = 22
	bodyInfo, _ := json.Marshal(reqBody)
	// 2.发送http请求
	url := "http://httpbin.org/post"

	resp, err := netutil.HttpPost(url, header, nil, bodyInfo)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func getW() wflow.WflowInterface {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:ye199169@(127.0.0.1:3306)/wflow?charset=utf8&parseTime=true&loc=Local",
	})
	ctx := context.Background()
	var w wflow.WflowInterface
	w = wflow.New(11, store.Client, ctx)
	return w
}

func Test(t *testing.T) {
	t.Log("snowflake.SnowGen.NextVal()", snowflake.SnowGen.NextVal())
	t.Log("snowflake.SnowGen.NextVal()", snowflake.SnowGen.NextVal())
	t.Log("snowflake.SnowGen.NextVal()", snowflake.SnowGen.NextVal())
	t.Log("snowflake.SnowGen.NextVal()", snowflake.SnowGen.NextVal())
}

func TestDbCreateTable(t *testing.T) {
	store := store.NewCommonStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "db",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(127.0.0.1:3306)/wflow?charset=utf8&parseTime=true&loc=Local",
	})
	ctx := context.Background()
	if err := store.Schema.Create(ctx); err != nil {
		t.Logf("failed creating schema resources: %v", err)
	}
}

func TestCreateProcDef(t *testing.T) {
	w := getW()
	in := new(kernel.SaveProcDefReq)
	in.FormId = 47729803
	in.AppId = 335609
	in.Name = "手机维修流程"
	in.Code = "MobileRepair"
	in.FormName = "维修"
	in.AppName = "手机售后"
	in.Resource = "{\"name\":\"发起人\",\"type\":\"ROOT\",\"nodeId\":\"ROOT\",\"parentId\":\"\",\"children\":{\"name\":\"抄送人\",\"type\":\"NOTIFIER\",\"nodeId\":\"node_470518668595\",\"parentId\":\"ROOT\",\"children\":{\"name\":\"审批人\",\"type\":\"APPROVAL\",\"nodeId\":\"node_146074821718\",\"parentId\":\"node_470518668595\",\"children\":{\"name\":\"抄送人\",\"type\":\"NOTIFIER\",\"nodeId\":\"node_146164582104\",\"parentId\":\"node_146074821718\",\"children\":{\"name\":\"审批人\",\"type\":\"APPROVAL\",\"nodeId\":\"node_146439246668\",\"parentId\":\"node_146164582104\",\"children\":{\"name\":\"抄送人\",\"type\":\"NOTIFIER\",\"nodeId\":\"node_146719095357\",\"parentId\":\"node_146439246668\",\"children\":{\"name\":\"\",\"type\":\"\",\"nodeId\":\"\",\"parentId\":\"\",\"children\":null,\"branches\":null,\"props\":null,\"conditionProps\":null},\"branches\":null,\"props\":{\"mode\":\"\",\"station\":\"\",\"assignedUser\":[{\"id\":61769798,\"name\":\"李四\",\"type\":\"user\"}],\"refuse\":null},\"conditionProps\":null},\"branches\":null,\"props\":{\"mode\":\"OR\",\"station\":\"\",\"assignedUser\":[{\"id\":381496,\"name\":\"旅人\",\"type\":\"user\"},{\"id\":568898,\"name\":\"王翠花\",\"type\":\"user\"}],\"refuse\":{\"type\":\"TO_NODE\",\"target\":\"node_146074821718\"}},\"conditionProps\":null},\"branches\":null,\"props\":{\"assignedUser\":[{\"id\":61769798,\"name\":\"李四\",\"type\":\"user\"},{\"id\":489564,\"name\":\"李秋香\",\"type\":\"user\"}],\"refuse\":null},\"conditionProps\":null},\"branches\":null,\"props\":{\"mode\":\"AND\",\"station\":\"\",\"assignedUser\":[{\"id\":381496,\"name\":\"旅人\",\"type\":\"user\"}],\"refuse\":{\"type\":\"TO_BEFORE\",\"target\":\"\"}},\"conditionProps\":null},\"branches\":null,\"props\":{\"mode\":\"\",\"station\":\"\",\"assignedUser\":[{\"id\":381496,\"name\":\"旅人\",\"type\":\"user\"}],\"refuse\":null},\"conditionProps\":null},\"branches\":null,\"props\":{\"mode\":\"\",\"station\":\"\",\"assignedUser\":null,\"refuse\":null},\"conditionProps\":null}"

	w.CreateProcDef(in)
}

func TestSetActive(t *testing.T) {
	w := getW()
	in := new(kernel.QueryProcDefReq)
	in.FormId = 47729803
	in.AppId = 335609
	in.Version = 1
	w.SetIsActive(in)
}

//func TestFindDefByVersion(t *testing.T) {
//	w := getW()
//	in := new(kernel.QueryProcDefReq)
//	in.FormId = 47729803
//	in.AppId = 335609
//	in.Version = 1
//	data := w.FindDefByVersion(in)
//	log.Println(data)
//}

func TestStartView(t *testing.T) {
	w := getW()
	in := new(kernel.StartViewReq)
	in.FormId = 47729803
	in.AppId = 335609
	in.DataId = 33334
	data, _ := w.StartView(in)
	log.Printf("data:%v", data.Resource)
}

func TestStartProcess(t *testing.T) {
	w := getW()
	in := new(kernel.StartViewReq)
	in.FormId = 47729803
	in.AppId = 335609
	in.DataId = 33334
	data, _ := w.StartView(in)
	in2 := new(kernel.StartProcInstReq)
	in2.FormId = 47729803
	in2.AppId = 335609
	in2.DataId = 33334
	in2.Resource = data.Resource
	in2.Notifiers = data.Notifiers
	// TODO
	in2.DefId = data.DefId
	w.StartProcess(in2)
	log.Println("启动成功")
	//log.Printf("data:%v", data.Resource)
}

//func TestFindProcInstByDataId(t *testing.T) {
//	w := getW()
//	inst, _ := w.FindProcInstByDataId(33334)
//	log.Printf("inst:%v", inst)
//}

func TestCompleteProcess(t *testing.T) {
	w := getW()
	in := new(kernel.CompleteProcInstReq)
	in.Result = 6
	in.Comment = "PASS"
	in.DataId = 33334
	w.CompleteProcess(in)
	log.Println("审批成功")
}
