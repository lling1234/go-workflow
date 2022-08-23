package util2

import (
	"encoding/json"
	"errors"
	"log"

	idworker "github.com/gitstliu/go-id-worker"
)

// 生成雪花ID
func NewSnowflakeId() int64 {
	worker := &idworker.IdWorker{}
	worker.InitIdWorker(1000, 1)
	id, err := worker.NextId()
	if err != nil {
		log.Println("输出雪花id错误", err)
		return -1
	}
	return id
}

func StructToJson(T any) (string, error) {
	data, err := json.Marshal(T)
	if err != nil {
		return "", errors.New("json转换失败")
	}
	return string(data), nil
}
