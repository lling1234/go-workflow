package service

import (
	const2 "act/rpc/internal/workflow-const"
	util2 "act/rpc/internal/workflow-util"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/mumushuiding/util"
	"act/rpc/internal/flow"
	"act/rpc/internal/model"
)

// ProcessReceiver 接收页面传递参数
type ProcessReceiver struct {
	UserID     int64  `json:"userId"`
	ProcInstID string `json:"procInstID"`
	Username   string `json:"username"`
	//Company    string             `json:"company"`
	ProcName string `json:"procName"`
	Title    string `json:"title"`
	//Department string             `json:"department"`
	//Var *map[string]string `json:"var"`

	TargetId   int64  `json:"targetId"`
	YewuFormId string `json:"yewuFormId"`
	DataId     int64  `json:"dataId"`
}

// ProcessPageReceiver 分页参数
type ProcessPageReceiver struct {
	util.Page
	// 我分管的部门
	//Departments []string `json:"departments"`
	// 我所属于的用户组或者角色
	//Groups   []string `josn:"groups"`
	UserID   int64  `json:"userID"`
	Username string `json:"username"`
	//Company    string   `json:"company"`
	ProcName   string `json:"procName"`
	ProcInstID string `json:"procInstID"`
	TargetId   int64  `json:"targetId"`
}

type MyCompleteProcInstReceiver struct {
	util.Page
	UserID   int64 `json:"userID"`
	TargetId int64 `json:"targetId"`
}

//var copyLock sync.Mutex

// GetDefaultProcessPageReceiver GetDefaultProcessPageReceiver
func GetDefaultProcessPageReceiver() *ProcessPageReceiver {
	var p = ProcessPageReceiver{}
	p.PageIndex = const2.DefaultPageIndex
	p.PageSize = const2.DefaultPageSize
	return &p
}
func findAll(pr *ProcessPageReceiver) ([]*model.ProcInst, int, error) {
	var page = util.Page{}
	page.PageRequest(pr.PageIndex, pr.PageSize)
	return model.FindProcInsts(pr.UserID, pr.ProcName, pr.TargetId, pr.PageIndex, pr.PageSize)
}

// FindProcInstByID FindProcInstByID
func FindProcInstByID(id int64) (string, error) {
	data, err := model.FindProcInstByID(id)
	if err != nil {
		return "", err
	}
	return util.ToJSONStr(data)
}

func FindProcInstStructByID(id int64) (*model.ProcInst, error) {
	data, err := model.FindProcInstByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FindMyCompleteProcInst(userID int64, targetId int64, pageIndex, pageSize int) (string, error) {
	inst, err := model.FindMyCompleteProcInst(userID, targetId, pageIndex, pageSize)
	if err != nil {
		return "", err
	}
	return util.ToJSONStr(inst)
}

// FindAllPageAsJSON FindAllPageAsJSON
func FindAllPageAsJSON(pr *ProcessPageReceiver) (string, error) {
	datas, count, err := findAll(pr)
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, pr.PageIndex, pr.PageSize)
}

func FindAllMyPendingInst(pr *MyCompleteProcInstReceiver) (string, error) {
	inst, err := model.FindAllMyPendingInst(pr.UserID, pr.TargetId, pr.PageIndex, pr.PageSize)
	if err != nil {
		return "", err
	}
	return util.ToJSONStr(inst)
}

// FindMyProcInstByToken FindMyProcInstByToken
// 根据token获取流程信息
//func FindMyProcInstByToken(token string, receiver *ProcessPageReceiver) (string, error) {
//	// 根据 token 获取用户信息
//	userinfo, err := GetUserinfoFromRedis(token)
//	if err != nil {
//		return "", err
//	}
//	if len(userinfo.Company) == 0 {
//		return "", errors.New("保存在redis中的【用户信息 userinfo】字段 company 不能为空")
//	}
//	if len(userinfo.ID) == 0 {
//		return "", errors.New("保存在redis中的【用户信息 userinfo】字段 ID 不能为空")
//	}
//	receiver.Company = userinfo.Company
//	receiver.Departments = userinfo.Departments
//	receiver.Groups = userinfo.Roles
//	receiver.UserID = userinfo.ID
//	// str, _ = util.ToJSONStr(receiver)
//	// fmt.Printf("receiver:%s\n", str)
//	return FindAllPageAsJSON(receiver)
//}

// StartProcessInstanceByToken 启动流程
//func StartProcessInstanceByToken(token string, p *ProcessReceiver) (int64, error) {
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
//	if len(userinfo.Department) == 0 {
//		return 0, errors.New("保存在redis中的【用户信息 userinfo】字段 department 不能为空")
//	}
//	p.Company = userinfo.Company
//	p.Department = userinfo.Department
//	p.UserID = userinfo.ID
//	p.Username = userinfo.Username
//	return p.StartProcessInstanceByID(p.Var)
//}

// StartProcessInstanceByID 启动流程
func (p *ProcessReceiver) StartProcessInstanceByID() (int64, error) {
	// runtime.GOMAXPROCS(2)
	// 获取流程定义
	node, prodefID, err := GetResourceByYewuFormIdAndTargetId(p.YewuFormId, p.TargetId)
	if err != nil {
		return 0, err
	}
	//--------以下需要添加事务-----------------
	level := 0 // 0 为开始节点
	tx := model.GetTx()
	// 新建流程实例
	var procInst = model.ProcInst{
		Identity: model.Identity{
			ID:         util2.NewSnowflakeId(),
			CreateTime: time.Now(),
		},
		ProcDefID:     prodefID,
		Title:         p.Title,
		StartTime:     time.Now(),
		StartUserID:   p.UserID,
		StartUserName: p.Username,
		TargetId:      p.TargetId,
		DataId:        p.DataId,
		State:         flow.PENDING,
		IsDel:         const2.IsDelNo,
	}
	//开启事务
	// times = time.Now()
	//procInst.ID = util2.NewSnowflakeId()
	procInstID, err := CreateProcInstTx(&procInst, tx) // 事务
	// fmt.Printf("启动流程实例耗时：%v", time.Since(times))
	exec := &model.Execution{
		Identity: model.Identity{
			ID: util2.NewSnowflakeId(),
		},
		ProcDefID:  prodefID,
		ProcInstID: procInstID,
	}
	task := &model.Task{
		Identity: model.Identity{
			ID: util2.NewSnowflakeId(),
		},
		NodeID:     "开始",
		ProcInstID: procInstID,
		//Assignee:      p.UserID,
		IsFinished:    1,
		ClaimTime:     time.Now(),
		Level:         level,
		MemberCount:   1,
		UnCompleteNum: 0,
		ActType:       "or",
		AgreeNum:      1,
	}
	// 生成执行流，一串运行节点
	_, err = GenerateExec(exec, node, p.UserID, tx) //事务
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	err = util.Str2Struct(exec.NodeInfos, &nodeInfos)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// -----------------生成新任务-------------
	if nodeInfos[0].ActType == "and" {
		task.UnCompleteNum = nodeInfos[0].MemberCount
		task.MemberCount = nodeInfos[0].MemberCount
	}
	_, err = NewTaskTx(task, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// fmt.Printf("生成新任务耗时：%v", time.Since(times))
	//--------------------流转------------------
	// times = time.Now()
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	// 流程移动到下一环节
	err = MoveStage(nodeInfos, p.UserID, p.Username, p.DataId, p.TargetId, "启动流程", task.ID, procInstID, level, flow.HAVEPASS, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// fmt.Printf("流转到下一流程耗时：%v", time.Since(times))
	// fmt.Println("--------------提交事务----------")
	tx.Commit() //结束事务
	return procInstID, err
}

// CreateProcInstByID 新建流程实例
// func CreateProcInstByID(processDefinitionID int, startUserID string) (int, error) {
// 	var procInst = model.ProcInst{
// 		ProcDefID:   processDefinitionID,
// 		StartTime:   util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS),
// 		StartUserID: startUserID,
// 	}
// 	return procInst.Save()
// }

// CreateProcInstTx CreateProcInstTx
// 开户事务
func CreateProcInstTx(procInst *model.ProcInst, tx *gorm.DB) (int64, error) {

	return procInst.SaveTx(tx)
}

// SetProcInstFinish SetProcInstFinish
// 设置流程结束
func SetProcInstFinish(procInstID int64, endTime time.Time, tx *gorm.DB) error {
	var p = &model.ProcInst{}
	p.ID = procInstID
	p.EndTime = endTime
	p.IsFinished = 1
	return p.UpdateTx(tx)
}

// StartByMyself 我发起的流程
func StartByMyself(receiver *ProcessPageReceiver) (string, error) {
	var page = util.Page{}
	page.PageRequest(receiver.PageIndex, receiver.PageSize)
	datas, count, err := model.StartByMyself(receiver.UserID, receiver.TargetId, receiver.PageIndex, receiver.PageSize)
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, receiver.PageIndex, receiver.PageSize)
}

// FindProcNotify 查询抄送我的
func FindProcNotify(receiver *ProcessPageReceiver) (string, error) {
	var page = util.Page{}
	page.PageRequest(receiver.PageIndex, receiver.PageSize)
	datas, count, err := model.FindProcNotify(receiver.UserID, receiver.TargetId, receiver.PageIndex, receiver.PageSize)
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, receiver.PageIndex, receiver.PageSize)
}

// UpdateProcInst UpdateProcInst
// 更新流程实例
func UpdateProcInst(procInst *model.ProcInst, tx *gorm.DB) error {
	return procInst.UpdateTx(tx)
}

// MoveFinishedProcInstToHistory MoveFinishedProcInstToHistory
//func MoveFinishedProcInstToHistory() error {
//	// 要注意并发，可能会运行多个app实例
//	// 加锁
//	copyLock.Lock()
//	defer copyLock.Unlock()
//	// 从pro_inst表查询已经结束的流程
//	proinsts, err := model.FindFinishedProc()
//	if err != nil {
//		return err
//	}
//	if len(proinsts) == 0 {
//		return nil
//	}
//	for _, v := range proinsts {
//		// 复制 proc_inst
//		//duration, err := util.TimeStrSub(v.EndTime, v.StartTime, util.YYYY_MM_DD_HH_MM_SS)
//		//if err != nil {
//		//	return err
//		//}
//		//v.Duration = duration
//		err = copyProcToHistory(v)
//		if err != nil {
//			return err
//		}
//		tx := model.GetTx()
//		// 流程实例的task移至历史纪录
//		err = copyTaskToHistoryByProInstID(v.ID, tx)
//		if err != nil {
//			tx.Rollback()
//			DelProcInstHistoryByID(v.ID)
//			return err
//		}
//		// execution移至历史纪录
//		err = copyExecutionToHistoryByProcInstID(v.ID, tx)
//		if err != nil {
//			tx.Rollback()
//			DelProcInstHistoryByID(v.ID)
//			return err
//		}
//		// identitylink移至历史纪录
//		err = copyIdentitylinkToHistoryByProcInstID(v.ID, tx)
//		if err != nil {
//			tx.Rollback()
//			DelProcInstHistoryByID(v.ID)
//			return err
//		}
//		// 删除流程实例
//		err = DelProcInstByIDTx(v.ID, tx)
//		if err != nil {
//			tx.Rollback()
//			DelProcInstHistoryByID(v.ID)
//			return err
//		}
//		tx.Commit()
//	}
//
//	return nil
//}

// DelProcInstByIDTx DelProcInstByIDTx
func DelProcInstByIDTx(procInstID int64, tx *gorm.DB) error {
	return model.DelProcInstByIDTx(procInstID, tx)
}

//func copyIdentitylinkToHistoryByProcInstID(procInstID int64, tx *gorm.DB) error {
//	return model.CopyIdentitylinkToHistoryByProcInstID(procInstID, tx)
//}
//func copyExecutionToHistoryByProcInstID(procInstID int64, tx *gorm.DB) error {
//	return model.CopyExecutionToHistoryByProcInstIDTx(procInstID, tx)
//}
//func copyProcToHistory(procInst *model.ProcInst) error {
//	return model.SaveProcInstHistory(procInst)
//
//}
//func copyTaskToHistoryByProInstID(procInstID int64, tx *gorm.DB) error {
//	return model.CopyTaskToHistoryByProInstID(procInstID, tx)
//}
