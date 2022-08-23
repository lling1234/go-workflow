package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// import _ "github.com/jinzhu/gorm"

// Task 流程任务表
// ExecutionID 执行流ID
// Name 任务名称，在流程文件中定义
// TaskDefKey 任务定义的ID值
// Assignee 被指派执行该任务的人
// Owner 任务拥有人
type Task struct {
	Identity
	// Company 任务创建人对应的公司
	// Company string `json:"company"`
	// ExecutionID     string `json:"executionID"`
	// 当前执行流所在的节点
	NodeID string `json:"nodeId"`
	Level  int    `json:"level"`
	// 流程实例id
	ProcInstID int64 `json:"procInstID"`
	//Assignee   string `json:"assignee"`
	ClaimTime time.Time `json:"claimTime"`
	// 还未审批的用户数，等于0代表会签已经全部审批结束，默认值为1
	MemberCount   int `json:"memberCount" gorm:"default:1"`
	UnCompleteNum int `json:"unCompleteNum" gorm:"default:1"`
	//审批通过数
	AgreeNum int `json:"agreeNum"`
	// and 为会签，or为或签，默认为or
	ActType    string `json:"actType" gorm:"default:'or'"`
	IsFinished int    `gorm:"default:false" json:"isFinished"`
}

func (Task) TableName() string {
	return "act_task"
}

// NewTask 新建任务
func (t *Task) NewTask() (int64, error) {
	err := db.Create(t).Error
	if err != nil {
		return 0, err
	}
	return t.ID, nil
}

// UpdateTx UpdateTx
func (t *Task) UpdateTx(tx *gorm.DB) error {
	err := tx.Model(&Task{}).Updates(t).Error
	return err
}

// GetTaskByID GetTaskById
func GetTaskByID(id int64) (*Task, error) {
	var t = &Task{}
	err := db.Where("id=?", id).Find(t).Error
	return t, err
}

// GetTaskLastByProInstID GetTaskLastByProInstID
// 根据流程实例id获取上一个任务
func GetTaskLastByProInstID(procInstID int64) (*Task, error) {
	var t = &Task{}
	err := db.Where("proc_inst_id=? and is_finished=1", procInstID).Order("claim_time desc").First(t).Error
	return t, err
}

// NewTaskTx begin tx
// 开启事务
func (t *Task) NewTaskTx(tx *gorm.DB) (int64, error) {
	// str, _ := util.ToJSONStr(t)
	// fmt.Printf("newTask:%s", str)
	t.CreateTime = time.Now()
	t.ClaimTime = time.Now()
	err := tx.Create(t).Error
	if err != nil {
		return 0, err
	}
	return t.ID, nil
}

// DeleteTask 删除任务
//func DeleteTask(id int64) error {
//	err := db.Where("id=?", id).Delete(&Task{}).Error
//	return err
//}
