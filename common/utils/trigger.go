package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/duke-git/lancet/v2/netutil"
)

type Headers struct {
	Name string `json:"name"`
	Val  string `json:"val"`
}

type WEBHook struct {
	Type    string                 `json:"type"`
	Method  string                 `json:"method"`
	URL     string                 `json:"url"`
	Headers []Headers              `json:"headers"`
	Body    map[string]interface{} `json:"params"`
}

/*******************函数声明*****************************
 **  函数名称:   TriggerStart
 **  功能说明:   触发器启动
 **  入参    :   1.type WEBHook struct
 **  返回    :   1.string http响应信息
 **	             2.error 错误信息
 ********************************************************/
func TriggerStart(w WEBHook) (string, error) {
	header := StructSliceToMap(w.Headers)
	return DOHttp(w, header)
}

// StructSliceToMapSlice : struct切片转为map
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
func DOHttp(w WEBHook, header map[string]string) (string, error) {
	var respBody string
	body, err := json.Marshal(w.Body)
	if err != nil {
		return "", err
	}
	switch {
	case w.Method == "get":
		resp, err := netutil.HttpGet(w.URL, header, nil, body)
		if err != nil {
			log.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		respBody = string(body)
	case w.Method == "put":
		resp, err := netutil.HttpPut(w.URL, header, nil, body)
		if err != nil {
			log.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		respBody = string(body)
	case w.Method == "delete":
		resp, err := netutil.HttpDelete(w.URL, header, nil, body)
		if err != nil {
			log.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		respBody = string(body)
	default:
		resp, err := netutil.HttpPost(w.URL, header, nil, body)
		if err != nil {
			log.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		respBody = string(body)
	}
	return respBody, nil
}
