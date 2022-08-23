package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Candidate 候选人 当流程走到该节点时，会获取到该节点的所有的候选人，并存入候选人表中
// 审批完成后，将数据存入到identitylink中。在整个流程完成后，删除改流程实例对应的候选人数据。
type Candidate struct {
	Identity
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName"`
	ProcInstID int64  `json:"procInstID"`
	TargetId   int64  `json:"targetId"`
	TaskID     int64  `json:"taskID,omitempty"`
	IsFinish   int    `json:"isFinish"`
}

func (Candidate) TableName() string {
	return "act_candidate"
}

// SaveTx SaveTx
func (c *Candidate) SaveTx(tx *gorm.DB) error {
	c.CreateTime = time.Now()
	if c.UserId == 0 {
		return errors.New("Candidate表的UserId字段不能为空！！")
	}
	err := tx.Create(c).Error
	return err
}

// DelByProcInstID 删除历史候选人
func (c *Candidate) DelByProcInstID(procInstID int64, tx *gorm.DB) error {
	return tx.Where("proc_inst_id=? ", procInstID).Delete(&Candidate{}).Error
}

// ExistsCandidateByTaskId 查询 用户是否包含在审批人中
func ExistsCandidateByTaskId(taskId int64, userId int64) (bool, error) {
	var count int
	err := db.Model(&Candidate{}).Where(" task_id=?  and user_id=?", taskId, userId).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
