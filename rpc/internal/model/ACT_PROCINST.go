package model

import (
	"strconv"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

// ProcInst 流程实例
type ProcInst struct {
	Identity
	// 流程定义ID
	ProcDefID int64 `json:"procDefId"`
	// 流程定义名
	//ProcDefName string `json:"procDefName"`
	// title 标题
	Title string `json:"title,omitempty"`
	// 用户部门
	//Department string `json:"department"`
	//Company    string `json:"company"`
	// 当前节点
	NodeID string `json:"nodeID,omitempty"`
	// 审批人
	//Candidate string `json:"candidate"`
	// 当前任务
	TaskID    int64     `json:"taskID"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	//Duration      int64  `json:"duration"`
	StartUserID   int64  `json:"startUserId,omitempty"`
	StartUserName string `json:"startUserName,omitempty"`
	IsFinished    int    `gorm:"default:2" json:"isFinished"`

	TargetId int64 `json:"targetId"`
	DataId   int64 `json:"dataId"`
	State    int   `gorm:"default:1" json:"state"`
	IsDel    int   `gorm:"default:2" json:"isDel"`
}

func (ProcInst) TableName() string {
	return "act_proc_inst"
}

// StartByMyself 我发起的流程
func StartByMyself(userID, targetId int64, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	maps := map[string]int64{
		"start_user_id": userID,
		"target_id":     targetId,
	}
	return findProcInsts(maps, pageIndex, pageSize)
}

// FindProcInstByID FindProcInstByID
func FindProcInstByID(id int64) (*ProcInst, error) {
	var data = ProcInst{}
	err := db.Where("id=?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FindProcNotify 查询抄送我的流程
func FindProcNotify(userID, targetId int64, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	var datas []*ProcInst
	var count int
	var sql string
	targetId_ := strconv.FormatInt(targetId, 10)
	userId_ := strconv.FormatInt(userID, 10)

	sql = "select proc_inst_id from identitylink i where i.type='notifier' and i.target_id='" + targetId_ + "' and i.user_id='" + userId_ + "'"

	err := db.Where("id in (" + sql + ")").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("start_time desc").Find(&datas).Error
	if err != nil {
		return datas, count, err
	}
	err = db.Model(&ProcInst{}).Where("id in (" + sql + ")").Count(&count).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}
func findProcInsts(maps map[string]int64, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	var datas []*ProcInst
	var count int
	selectDatas := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			err := db.Where(maps).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("start_time desc").Find(&datas).Error
			in <- err
			wg.Done()
		}()
	}
	selectCount := func(in chan<- error, wg *sync.WaitGroup) {
		err := db.Model(&ProcInst{}).Where(maps).Count(&count).Error
		in <- err
		wg.Done()
	}
	var err1 error
	var wg sync.WaitGroup
	numberOfRoutine := 2
	wg.Add(numberOfRoutine)
	errStream := make(chan error, numberOfRoutine)
	// defer fmt.Println("close channel")
	selectDatas(errStream, &wg)
	selectCount(errStream, &wg)
	wg.Wait()
	defer close(errStream) // 关闭通道
	for i := 0; i < numberOfRoutine; i++ {
		// log.Printf("send: %v", <-errStream)
		if err := <-errStream; err != nil {
			err1 = err
		}
	}
	// fmt.Println("结束")
	return datas, count, err1
}

// FindMyCompleteProcInst 分页查询我审批的流程
func FindMyCompleteProcInst(userID int64, targetId int64, pageIndex, pageSize int) ([]ProcInst, error) {
	var datas []ProcInst
	err := db.Where(" id IN \n ( \n SELECT t.proc_inst_id from act_identity_link i \n LEFT JOIN act_task t on i.task_id = t.id \n where i.user_id=? and i.target_id=? \n ) \n and is_del=2",
		userID, targetId).
		Find(&datas).Error
	return datas, err
}

func FindAllMyPendingInst(userID int64, targetId int64, pageIndex, pageSize int) ([]ProcInst, error) {
	var datas []ProcInst
	err := db.Where(" task_id IN \n (select c.task_id from act_candidate c  \n where not EXISTS (SELECT 1 FROM act_identity_link where user_id=c.user_id and task_id=c.task_id) \n and c.user_id=? and c.target_id=? \n ) and is_del=2",
		userID, targetId).
		Find(&datas).Error
	return datas, err
}

// FindProcInsts FindProcInsts
// 分页查询
func FindProcInsts(userID int64, procName string, targetId int64, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	var datas []*ProcInst
	var count int
	var sql = " target_id='" + strconv.FormatInt(targetId, 10) + "' and is_finished=2 "
	if len(procName) > 0 {
		sql += "and proc_def_name='" + procName + "'"
	}
	// fmt.Println(sql)
	selectDatas := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			//err := db.Scopes(GroupsNotNull(groups, sql), DepartmentsNotNull(departments, sql)).
			//	Or("candidate=? and "+sql, userID).
			//	Offset((pageIndex - 1) * pageSize).Limit(pageSize).
			//	Order("start_time desc").
			//	Find(&datas).Error
			//in <- err
			//wg.Done()
		}()
	}
	selectCount := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			//err := db.Scopes(GroupsNotNull(groups, sql), DepartmentsNotNull(departments, sql)).Model(&ProcInst{}).Or("candidate=? and "+sql, userID).Count(&count).Error
			//in <- err
			//wg.Done()
		}()
	}
	var err1 error
	var wg sync.WaitGroup
	numberOfRoutine := 2
	wg.Add(numberOfRoutine)
	errStream := make(chan error, numberOfRoutine)
	// defer fmt.Println("close channel")
	selectDatas(errStream, &wg)
	selectCount(errStream, &wg)
	wg.Wait()
	defer close(errStream) // 关闭通道

	for i := 0; i < numberOfRoutine; i++ {
		// log.Printf("send: %v", <-errStream)
		if err := <-errStream; err != nil {
			err1 = err
		}
	}
	// fmt.Println("结束")
	return datas, count, err1
}

// Save save
func (p *ProcInst) Save() (int64, error) {
	p.CreateTime = time.Now()
	err := db.Create(p).Error
	if err != nil {
		return 0, err
	}
	return p.ID, nil
}

//SaveTx SaveTx
func (p *ProcInst) SaveTx(tx *gorm.DB) (int64, error) {
	p.CreateTime = time.Now()
	p.EndTime = time.Now()
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return -1, err
	}
	return p.ID, nil
}

// DelProcInstByID DelProcInstByID
func DelProcInstByID(id int64) error {
	return db.Where("id=?", id).Delete(&ProcInst{}).Error
}

// DelProcInstByIDTx DelProcInstByIDTx
// 事务
func DelProcInstByIDTx(id int64, tx *gorm.DB) error {
	return tx.Where("id=?", id).Delete(&ProcInst{}).Error
}

// UpdateTx UpdateTx
func (p *ProcInst) UpdateTx(tx *gorm.DB) error {
	return tx.Model(&ProcInst{}).Updates(p).Error
}

// FindFinishedProc FindFinishedProc
func FindFinishedProc() ([]*ProcInst, error) {
	var datas []*ProcInst
	err := db.Where("is_finished=1").Find(&datas).Error
	return datas, err
}

func FindLeastTaskId(dataId int64) (int64, error) {
	var p = ProcInst{}
	err := db.Where("data_id =? ", dataId).First(&p).Error
	return p.TaskID, err
}
