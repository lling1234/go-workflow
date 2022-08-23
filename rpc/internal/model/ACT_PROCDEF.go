package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Procdef 流程定义表
type Procdef struct {
	Identity
	Name    string `json:"name,omitempty"`
	Version int    `json:"version,omitempty"`
	// 流程定义json字符串
	Resource string `gorm:"size:10000" json:"resource,omitempty"`
	// 用户id
	//Userid   string `json:"userid,omitempty"`
	//Username string `json:"username,omitempty"`
	// 用户所在公司
	//Company    string `json:"company,omitempty"`
	//DeployTime string `json:"deployTime,omitempty"`

	Code        string `json:"code,omitempty"`
	TargetId    int64  `json:"targetId,omitempty"`
	YewuFormId  string `json:"yewuFormId,omitempty"`
	YewuName    string `json:"yewuName,omitempty"`
	RemainHours int    `json:"remainHours,omitempty"`
	//RemainUnit     string `json:"remainUnit,omitempty"`
	CreateUserId   int64  `json:"createUserId,omitempty"`
	CreateUserName string `json:"createUserName,omitempty"`
	IsDel          int    `json:"isDel,omitempty"`
}

func (Procdef) TableName() string {
	return "act_proc_def"
}

// Save save and return id
// 保存并返回ID
func (p *Procdef) Save() (ID int64, err error) {
	p.CreateTime = time.Now()
	err = db.Create(p).Error
	if err != nil {
		return 0, err
	}
	return p.ID, nil
}

// SaveTx SaveTx
func (p *Procdef) SaveTx(tx *gorm.DB) error {
	p.CreateTime = time.Now()
	err := tx.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

// GetProcdefLatestByNameAndCompany :get latest procdef by name and company
// 根据名字和公司查询最新的流程定义
//func GetProcdefLatestByNameAndCompany(name, company string) (*Procdef, error) {
//	var p []*Procdef
//	err := db.Where("name=? and company=?", name, company).Order("version desc").Find(&p).Error
//	if err != nil || len(p) == 0 {
//		return nil, err
//	}
//	return p[0], err
//}
func GetProcdefLatestByYewuFormIdAndTargetId(yewuFormId string, targetId int64) (*Procdef, error) {
	var p []*Procdef
	err := db.Where("yewu_form_id=? and target_id=?", yewuFormId, targetId).Order("version desc").Find(&p).Error
	if err != nil || len(p) == 0 {
		return nil, err
	}
	return p[0], err
}

// GetProcdefByID 根据流程定义
func GetProcdefByID(id int64) (*Procdef, error) {
	var p = &Procdef{}
	err := db.Where("id=?", id).Find(p).Error
	return p, err
}

// DelProcdefByID del by id
// 根据id删除
func DelProcdefByID(id int) error {
	err := db.Where("id = ?", id).Delete(&Procdef{}).Error
	return err
}

// DelProcdefByIDTx DelProcdefByIDTx
//func DelProcdefByIDTx(id int64, tx *gorm.DB) error {
//	return tx.Where("id = ?", id).Delete(&Procdef{}).Error
//}

// FindProcdefsWithCountAndPaged return result with total count and error
// 返回查询结果和总条数
func FindProcdefsWithCountAndPaged(pageIndex, pageSize int, maps map[string]interface{}) (datas []*Procdef, count int, err error) {
	err = db.Select("id,name,version,yewu_form_id,yewu_name,target_id").Where(maps).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&datas).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	err = db.Model(&Procdef{}).Where(maps).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, count, nil
}

// MoveProcdefToHistoryByIDTx 将流程定义移至历史纪录表
//func MoveProcdefToHistoryByIDTx(ID int64, tx *gorm.DB) error {
//	err := tx.Exec("insert into procdef_history select * from procdef where id=?", ID).Error
//	if err != nil {
//		return err
//	}
//	return DelProcdefByIDTx(ID, tx)
//}
