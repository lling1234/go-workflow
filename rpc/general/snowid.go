package general

import (
	"github.com/gitstliu/go-id-worker"
	"log"
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
