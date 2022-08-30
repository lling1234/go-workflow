package service

import (
	"errors"
	const2 "act/rpc/internal/workflow-const"
	util2 "act/rpc/internal/workflow-util"
	"log"
	"sync"
	"time"

	"act/rpc/internal/flow"

	"github.com/mumushuiding/util"

	"act/rpc/internal/model"
)

var saveLock sync.Mutex

// Procdef 流程定义表
type Procdef struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	// 流程定义json字符串
	Resource  *flow.Node `json:"resource"`
	PageSize  int        `json:"pageSize"`
	PageIndex int        `json:"pageIndex"`

	Code        string `json:"code"`
	TargetId    int64  `json:"targetId"`
	YewuFormId  string `json:"yewuFormId"`
	YewuName    string `json:"yewuName"`
	RemainHours int    `json:"remainHours"`
	UserId      int64  `json:"userId"`
	UserName    string `json:"userName"`
}

// GetProcdefByID 根据流程定义id获取流程定义
func GetProcdefByID(id int64) (*model.Procdef, error) {
	return model.GetProcdefByID(id)
}

// GetProcdefLatestByNameAndCompany GetProcdefLatestByNameAndCompany
// 根据流程定义名字和公司查询流程定义
//func GetProcdefLatestByNameAndCompany(name, company string) (*model.Procdef, error) {
//	return model.GetProcdefLatestByYewuFormIdAndTargetId(name, company)
//}
func GetProcdefLatestByYewuFormIdAndTargetId(yewuFormId string, targetId int64) (*model.Procdef, error) {
	return model.GetProcdefLatestByYewuFormIdAndTargetId(yewuFormId, targetId)
}

// GetResourceByNameAndCompany GetResourceByNameAndCompany
// 获取流程定义配置信息
//func GetResourceByNameAndCompany(name, company string) (*flow.Node, int64, string, error) {
//	prodef, err := GetProcdefLatestByNameAndCompany(name, company)
//	if err != nil {
//		return nil, 0, "", err
//	}
//	if prodef == nil {
//		return nil, 0, "", errors.New("流程【" + name + "】不存在")
//	}
//	node := &flow.Node{}
//	err = util.Str2Struct(prodef.Resource, node)
//	return node, prodef.ID, prodef.Name, err
//}
func GetResourceByYewuFormIdAndTargetId(yewuFormId string, targetId int64) (*flow.Node, int64, error) {
	prodef, err := GetProcdefLatestByYewuFormIdAndTargetId(yewuFormId, targetId)
	if err != nil {
		return nil, 0, err
	}
	if prodef == nil {
		return nil, 0, errors.New("流程【" + yewuFormId + "】不存在")
	}
	node := &flow.Node{}
	err = util.Str2Struct(prodef.Resource, node)
	return node, prodef.ID, err
}

// GetResourceByID GetResourceByID
// 根据id查询流程定义
func GetResourceByID(id int64) (*flow.Node, int64, error) {
	prodef, err := GetProcdefByID(id)
	if err != nil {
		return nil, 0, err
	}
	node := &flow.Node{}
	err = util.Str2Struct(prodef.Resource, node)
	return node, prodef.ID, err
}

// SaveProcdefByToken SaveProcdefByToken
//func (p *Procdef) SaveProcdefByToken(token string) (int64, error) {
//	// 根据 token 获取用户信息
//	userinfo, err := GetUserinfoFromRedis(token)
//	if err != nil {
//		return 0, err
//	}
//	if len(userinfo.Company) == 0 {
//		return 0, errors.New("保存在redis中的【用户信息 userinfo】字段 company 不能为空")
//	}
//	if len(userinfo.Username) == 0 {
//		return 0, errors.New("保存在redis中的【用户信息 userinfo】字段 username 不能为空")
//	}
//	if len(userinfo.ID) == 0 {
//		return 0, errors.New("保存在redis中的【用户信息 userinfo】字段 ID 不能为空")
//	}
//	p.Company = userinfo.Company
//	p.Userid = userinfo.ID
//	p.Username = userinfo.Username
//	return p.SaveProcdef()
//}

// SaveProcdef 保存
func (p *Procdef) SaveProcdef() (id int64, err error) {
	// 流程定义有效性检验
	//err = IsProdefValid(p.Resource)
	if err != nil {
		return 0, err
	}
	log.Println("p.Resource00000",p.Resource)
	resource, err := util.ToJSONStr(p.Resource)
	if err != nil {
		return 0, err
	}
	log.Println("resource  111", resource)
	var procdef = model.Procdef{
		Identity: model.Identity{
			ID:         util2.NewSnowflakeId(),
			CreateTime: time.Now(),
		},
		Name:           p.Name,
		Code:           p.Code,
		Resource:       resource,
		TargetId:       p.TargetId,
		YewuFormId:     p.YewuFormId,
		YewuName:       p.YewuName,
		RemainHours:    p.RemainHours,
		CreateUserId:   p.UserId,
		CreateUserName: p.UserName,
	}
	return SaveProcdef(&procdef)
}

// SaveProcdef 保存
func SaveProcdef(p *model.Procdef) (id int64, err error) {
	// 参数是否为空判定
	saveLock.Lock()
	defer saveLock.Unlock()
	old, err := GetProcdefLatestByYewuFormIdAndTargetId(p.YewuFormId, p.TargetId)
	if err != nil {
		return 0, err
	}
	//p.DeployTime = util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS)
	p.CreateTime = time.Now()
	p.IsDel = const2.IsDelNo // 2为未删除
	if old == nil {
		p.Version = 1
		return p.Save()
	}
	tx := model.GetTx()
	// 保存新版本
	p.Version = old.Version + 1
	err = p.SaveTx(tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// 转移旧版本
	//err = model.MoveProcdefToHistoryByIDTx(old.ID, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return p.ID, nil
}

// ExistsProcdefByNameAndCompany if exists
// 查询流程定义是否存在
func ExistsProcdefByNameAndCompany(yewuFormId string, targetId int64) (yes bool, version int, err error) {
	p, err := GetProcdefLatestByYewuFormIdAndTargetId(yewuFormId, targetId)
	if p == nil {
		return false, 1, err
	}
	version = p.Version + 1
	return true, version, err
}

// FindAllPageAsJSON find by page and  transform result to string
// 分页查询并将结果转换成 json 字符串
func (p *Procdef) FindAllPageAsJSON() (string, error) {
	datas, count, err := p.FindAll()
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, p.PageIndex, p.PageSize)
}

// FindAll FindAll
func (p *Procdef) FindAll() ([]*model.Procdef, int, error) {
	var page = util.Page{}
	page.PageRequest(p.PageIndex, p.PageSize)
	maps := p.getMaps()
	return model.FindProcdefsWithCountAndPaged(page.PageIndex, page.PageSize, maps)
}
func (p *Procdef) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if len(p.YewuFormId) > 0 {
		maps["yewuFormId"] = p.YewuFormId
	}
	if p.TargetId != 0 {
		maps["targetId"] = p.TargetId
	}
	return maps
}

// DelProcdefByID del by id
func DelProcdefByID(id int) error {
	return model.DelProcdefByID(id)
}

// IsProdefValid 流程定义格式是否有效
func IsProdefValid(node *flow.Node) error {

	return flow.IfProcessConifgIsValid(node)
}
