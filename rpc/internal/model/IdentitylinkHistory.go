package model

import (
	"github.com/jinzhu/gorm"
)

type IdentityLinkHistory struct {
	IdentityLink
}

func CopyIdentitylinkToHistoryByProcInstID(procInstID int64, tx *gorm.DB) error {
	return tx.Exec("insert into identitylink_history select * from identitylink where proc_inst_id=?", procInstID).Error
}

func FindParticipantHistoryByProcInstID(procInstID int) ([]*IdentityLinkHistory, error) {
	var datas []*IdentityLinkHistory
	err := db.Select("id,user_id,step,comment").Where("proc_inst_id=? and type=?", procInstID, IdentityTypes[PARTICIPANT]).Order("id asc").Find(&datas).Error
	return datas, err
}
