package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Identitylink 用户组同任务的关系
type IdentityLink struct {
	Identity
	//Group      string `json:"group,omitempty"`
	Type       string `json:"type,omitempty"`
	UserID     int64  `json:"userid,omitempty"`
	UserName   string `json:"username,omitempty"`
	TaskID     int64  `json:"taskID,omitempty"`
	TargetId   int64  `json:"targetId,omitempty"`
	Step       int    `json:"step"`
	ProcInstID int64  `json:"procInstID,omitempty"`
	//DataId     int64  `json:"dataId,omitempty"`
	Comment string `json:"comment,omitempty"`
	Result  int    `json:"result"` // 3驳回至发起人、4驳回到上一级、6未通过、7已通过
}

func (IdentityLink) TableName() string {
	return "act_identity_link"
}

// IdentityType 类型
type IdentityType int

const (
	// CANDIDATE 候选
	CANDIDATE IdentityType = iota + 1
	// PARTICIPANT 参与人
	PARTICIPANT
	// MANAGER 上级领导
	MANAGER
	// NOTIFIER 抄送人
	NOTIFIER
)

// IdentityTypes IdentityTypes
var IdentityTypes = [...]string{CANDIDATE: "candidate", PARTICIPANT: "participant", MANAGER: "主管", NOTIFIER: "notifier"}

// SaveTx SaveTx
func (i *IdentityLink) SaveTx(tx *gorm.DB) error {
	i.CreateTime = time.Now()
	// if len(i.Company) == 0 {
	// 	return errors.New("Identitylink表的company字段不能为空！！")
	// }
	err := tx.Create(i).Error
	return err
}

// DelCandidateByProcInstID DelCandidateByProcInstID
// 删除历史候选人
func DelCandidateByProcInstID(procInstID int64, tx *gorm.DB) error {
	return tx.Where("proc_inst_id=? and type=?", procInstID, IdentityTypes[CANDIDATE]).Delete(&IdentityLink{}).Error
}

// ExistsNotifierByProcInstIDAndGroup 抄送人是否已经存在
func ExistsNotifierByProcInstID(procInstID int64) (bool, error) {
	var count int
	err := db.Model(&IdentityLink{}).Where("identitylink.proc_inst_id=?  and identitylink.type=?", procInstID, IdentityTypes[NOTIFIER]).Count(&count).Error
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

func ExistsIdentityByTaskId(taskId int64, userId int64) (bool, error) {
	var count int
	err := db.Model(&IdentityLink{}).Where(" task_id=?  and user_id=?", taskId, userId).Count(&count).Error
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

// IfParticipantByTaskID IfParticipantByTaskID
// 针对指定任务判断用户是否已经审批过了
func IfParticipantByTaskID(userID, taskID int64) (bool, error) {
	var count int
	err := db.Model(&IdentityLink{}).Where("user_id=?  and task_id=?", userID, taskID).Count(&count).Error
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

// FindParticipantByProcInstID 查询参与审批的人
func FindParticipantByProcInstID(procInstID int64) ([]*IdentityLink, error) {
	var datas []*IdentityLink
	err := db.Select("id,user_id,user_name,step,comment").Where("proc_inst_id=? and type=?", procInstID, IdentityTypes[PARTICIPANT]).Order("id asc").Find(&datas).Error
	return datas, err
}
