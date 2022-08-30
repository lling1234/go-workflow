package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mumushuiding/util"
	"act/rpc/internal/model"
	util2 "act/rpc/internal/workflow-util"
)

// SaveIdentitylinkTx SaveIdentitylinkTx
func SaveIdentityLinkTx(i *model.IdentityLink, tx *gorm.DB) error {
	return i.SaveTx(tx)
}

// AddNotifierTx 添加抄送人候选用户组
func AddNotifierTx(dataId int64, step int, procInstID int64, tx *gorm.DB) error {
	yes, err := ExistsNotifierByProcInstID(procInstID)
	if err != nil {
		return err
	}
	if yes {
		return nil
	}
	i := &model.IdentityLink{
		//Group:      group,
		Type:       model.IdentityTypes[model.NOTIFIER],
		Step:       step,
		ProcInstID: procInstID,
	}
	return SaveIdentityLinkTx(i, tx)
}

// AddCandidateGroupTx AddCandidateGroupTx
// 添加候选用户组
func AddIdentityLinkCandidateTx(targetId int64, step int, taskID int64, procInstID int64, tx *gorm.DB) error {
	err := DelCandidateByProcInstID(procInstID, tx)
	if err != nil {
		return err
	}
	i := &model.IdentityLink{
		//Group:      group,
		Type:       model.IdentityTypes[model.CANDIDATE],
		TaskID:     taskID,
		Step:       step,
		ProcInstID: procInstID,
		TargetId:   targetId,
		//DataId:     dataId,
	}
	return SaveIdentityLinkTx(i, tx)
}

// AddCandidateUserTx AddCandidateUserTx
// 添加候选用户
func AddCandidateUserTx(userID int64, dataId int64, step int, taskID int64, procInstID int64, tx *gorm.DB) error {
	err := DelCandidateByProcInstID(procInstID, tx)
	if err != nil {
		return err
	}
	i := &model.IdentityLink{
		UserID:     userID,
		Type:       model.IdentityTypes[model.CANDIDATE],
		TaskID:     taskID,
		Step:       step,
		ProcInstID: procInstID,
		//DataId:     dataId,
	}
	return SaveIdentityLinkTx(i, tx)
	// var wg sync.WaitGroup
	// var err1, err2 error
	// numberOfRoutine := 2
	// wg.Add(numberOfRoutine)
	// go func() {
	// 	defer wg.Done()
	// 	err1 = DelCandidateByProcInstID(procInstID, tx)
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	i := &model.Identitylink{
	// 		UserID:     userID,
	// 		Type:       model.IdentityTypes[model.CANDIDATE],
	// 		TaskID:     taskID,
	// 		Step:       step,
	// 		ProcInstID: procInstID,
	// 		Company:    company,
	// 	}
	// 	err2 = SaveIdentitylinkTx(i, tx)
	// }()
	// wg.Wait()
	// fmt.Println("保存identyilink结束")
	// if err1 != nil {
	// 	return err1
	// }
	// return err2
}

//AddParticipantTx AddParticipantTx
// 添加任务参与人
func AddParticipantTx(userID int64, username string, targetId int64, comment string, taskID int64, procInstID int64, step int, result int, tx *gorm.DB) error {
	i := &model.IdentityLink{
		Identity: model.Identity{
			ID: util2.NewSnowflakeId(),
		},
		Type:       model.IdentityTypes[model.PARTICIPANT],
		UserID:     userID,
		UserName:   username,
		ProcInstID: procInstID,
		Step:       step,
		TaskID:     taskID,
		TargetId:   targetId,
		Comment:    comment,
		Result:     result,
	}
	return SaveIdentityLinkTx(i, tx)
}

// IfParticipantByTaskID IfParticipantByTaskID
// 针对指定任务判断用户是否已经审批过了
func IfParticipantByTaskID(userID int64, taskID int64) (bool, error) {
	return model.IfParticipantByTaskID(userID, taskID)
}

// DelCandidateByProcInstID DelCandidateByProcInstID
// 删除历史候选人
func DelCandidateByProcInstID(procInstID int64, tx *gorm.DB) error {
	return model.DelCandidateByProcInstID(procInstID, tx)
}

// ExistsNotifierByProcInstIDAndGroup 抄送人是否已经存在
func ExistsNotifierByProcInstID(procInstID int64) (bool, error) {
	return model.ExistsNotifierByProcInstID(procInstID)
}

// FindParticipantByProcInstID 查询参与审批的人
func FindParticipantByProcInstID(procInstID int64) (string, error) {
	datas, err := model.FindParticipantByProcInstID(procInstID)
	if err != nil {
		return "", err
	}
	str, err := util.ToJSONStr(datas)
	if err != nil {
		return "", err
	}
	return str, nil
}
