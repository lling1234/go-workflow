package service

import (
	"errors"
	"fmt"
	const2 "act/rpc/internal/workflow-const"
	util2 "act/rpc/internal/workflow-util"
	"sync"
	"time"

	"act/rpc/internal/flow"

	"github.com/jinzhu/gorm"

	"act/rpc/internal/model"
)

// TaskReceiver 任务
type TaskReceiver struct {
	//TaskID   int64  `json:"taskID"`
	UserID   int64  `json:"userID,omitempty"`
	UserName string `json:"username,omitempty"`
	//Pass     string `json:"pass,omitempty"`
	ProcInstID int64  `json:"procInstID,omitempty"`
	Comment    string `json:"comment,omitempty"`
	DataId     int64  `json:"dataId"`
	Result     int    `json:"result,omitempty"`
	TargetId   int64  `json:"targetId"`
}
type WithdrawReceiver struct {
	UserID     int64  `json:"userID,omitempty"`
	UserName   string `json:"username,omitempty"`
	ProcInstID int64  `json:"procInstID,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

var completeLock sync.Mutex

// NewTask 新任务
func NewTask(t *model.Task) (int64, error) {
	if len(t.NodeID) == 0 {
		return 0, errors.New("request param nodeID can not be null / 任务当前所在节点nodeId不能为空！")
	}
	t.CreateTime = time.Now()
	return t.NewTask()
}

// NewTaskTx NewTaskTx
// 开启事务
func NewTaskTx(t *model.Task, tx *gorm.DB) (int64, error) {
	if len(t.NodeID) == 0 {
		return 0, errors.New("request param nodeID can not be null / 任务当前所在节点nodeId不能为空！")
	}
	t.CreateTime = time.Now()
	return t.NewTaskTx(tx)
}

// DeleteTask 删除任务
//func DeleteTask(id int64) error {
//	return model.DeleteTask(id)
//}

// GetTaskByID GetTaskById
func GetTaskByID(id int64) (task *model.Task, err error) {
	return model.GetTaskByID(id)
}

// GetTaskLastByProInstID GetTaskLastByProInstID
func GetTaskLastByProInstID(procInstID int64) (*model.Task, error) {
	return model.GetTaskLastByProInstID(procInstID)
}

// CompleteByToken 通过token 审批任务
//func CompleteByToken(token string, receiver *TaskReceiver) error {
//	userinfo, err := GetUserinfoFromRedis(token)
//	if err != nil {
//		return err
//	}
//	pass, err := strconv.ParseBool(receiver.Pass)
//	if err != nil {
//		return err
//	}
//	err = Complete(receiver.TaskID, userinfo.ID, userinfo.Username, userinfo.Company, receiver.Comment, receiver.Candidate, pass)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// Complete Complete
// 审批
func Complete(taskID int64, userID int64, username string, dataId int64, targetId int64, comment string, result int) error {
	tx := model.GetTx()
	err := CompleteTaskTx(taskID, userID, username, dataId, targetId, comment, result, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func RejectToPre(taskID int64, userID int64, username string, dataId int64, targetId int64, comment string, result int) error {
	tx := model.GetTx()
	err := CompleteTaskTx(taskID, userID, username, dataId, targetId, comment, result, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func RejectToStart(taskID int64, userID int64, username string, dataId int64, targetId int64, comment string, result int) error {
	tx := model.GetTx()
	err := CompleteTaskTx(taskID, userID, username, dataId, targetId, comment, result, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateTaskWhenComplete UpdateTaskWhenComplete
func UpdateTaskWhenComplete(taskID int64, result int, tx *gorm.DB) (*model.Task, error) {
	// 获取task
	completeLock.Lock()         // 关锁
	defer completeLock.Unlock() //解锁
	// 查询任务
	// TODO entsql SELECT * FROM `act_task`  WHERE (id=9212524636351135768)
	task, err := GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("任务【" + fmt.Sprintf("%d", task.ID) + "】不存在")
	}
	// 判断是否已经结束
	if task.IsFinished == 1 {
		if task.NodeID == "结束" {
			return nil, errors.New("流程已经结束")
		}
		return nil, errors.New("任务【" + fmt.Sprintf("%d", taskID) + "】已经被审批过了！！")
	}
	// 设置处理人和处理时间
	//task.Assignee = userID
	task.ClaimTime = time.Now()
	// ----------------会签 （默认全部通过才结束），只要存在一个不通过，就结束，然后流转到上一步
	//同意
	if result == flow.HAVEPASS {
		task.AgreeNum++
	}
	// 未审批人数减一
	task.UnCompleteNum--
	// 判断是否结束
	if task.UnCompleteNum == 0 {
		task.IsFinished = 1
	}
	// TODO entsql
	/*
		UPDATE `act_task`
		SET `act_type` = 'and',
		`agree_num` = 1,
		`claim_time` = '2022-08-30 15:25:25',
		`create_time` = '2022-08-30 15:23:45',
		`id` = 9212524636351135768,
		`is_finished` = 1,
		`level` = 1,
		`member_count` = 1,
		`node_id` = '1659320194369',
		`proc_inst_id` = 5469563801642631168
		WHERE `act_task`.`id` = 9212524636351135768
	*/
	err = task.UpdateTx(tx)
	// str, _ := util.ToJSONStr(task)
	// log.Println(str)
	if err != nil {
		return nil, err
	}
	return task, nil
}

//CompleteTaskTx 执行任务
func CompleteTaskTx(taskID int64, userID int64, username string, dataId int64, targetId int64, comment string, result int, tx *gorm.DB) error {
	//更新任务
	// TODO entsql
	task, err := UpdateTaskWhenComplete(taskID, result, tx)
	if err != nil {
		return err
	}
	// 如果是会签
	// TODO entsql  SELECT count(*) FROM `act_identity_link`  WHERE (user_id=122298283  and task_id=9212524636351135768)
	if task.ActType == "and" {
		// fmt.Println("------------------是会签，判断用户是否已经审批过了，避免重复审批-------")
		// 判断用户是否已经审批过了（存在会签的情况）
		yes, err := IfParticipantByTaskID(userID, taskID)
		if err != nil {
			tx.Rollback()
			return err
		}
		if yes {
			tx.Rollback()
			return errors.New("您已经审批过了，请等待他人审批！）")
		}
	}

	// 查看任务的未审批人数是否为0，不为0就不流转
	if result != 0 {
		// 添加参与人
		/*TODO entsql
		INSERT INTO `act_identity_link`
		(`id`,`create_time`,`type`,`user_id`,`user_name`,`task_id`,`target_id`,`step`,`proc_inst_id`,`comment`,`result`)
		VALUES (5893377269455290368,'2022-08-30 15:25:25','participant',122298283,'小春',9212524636351135768,1727882,1,5469563801642631168,'小春评论1111',7)*/
		err := AddParticipantTx(userID, username, targetId, comment, task.ID, task.ProcInstID, task.Level, result, tx)
		if err != nil {
			return err
		}
		//return nil
	}
	// 流转到下一流程
	// nodeInfos, err := GetExecNodeInfosByProcInstID(task.ProcInstID)
	// if err != nil {
	// 	return err
	// }
	err = MoveStageByProcInstID(userID, username, dataId, targetId, comment, taskID, task.ProcInstID, task.Level, result, tx)
	if err != nil {
		return err
	}

	return nil
}

// WithDrawTaskByToken 撤回任务
//func WithDrawTaskByToken(token string, receiver *TaskReceiver) error {
//	userinfo, err := GetUserinfoFromRedis(token)
//	if err != nil {
//		return err
//	}
//	if userinfo.ID == 0 {
//		return errors.New("保存在redis中的【用户信息 userinfo】字段 ID 不能为空！！")
//	}
//	if len(userinfo.Username) == 0 {
//		return errors.New("保存在redis中的【用户信息 userinfo】字段 username 不能为空！！")
//	}
//	if len(userinfo.Company) == 0 {
//		return errors.New("保存在redis中的【用户信息 userinfo】字段 company 不能为空")
//	}
//	return WithDrawTask(receiver.TaskID, receiver.ProcInstID, userinfo.ID, userinfo.Username, receiver.DataId, receiver.Comment)
//}

// WithDrawTask 撤回任务
func WithDrawTask(procInstID int64, userID int64, comment string) error {
	procInst, err := FindProcInstStructByID(procInstID)
	if err != nil {
		return err
	}
	if procInst == nil {
		return errors.New("查询procInst结果为空")
	}
	if procInst.State != flow.PENDING {
		return errors.New("该流程不符合撤回条件")
	}
	if userID != procInst.StartUserID {
		return errors.New("必须由发起人撤回")
	}
	task, err := GetTaskByID(procInst.TaskID)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("查询task结果为空")
	}
	task.IsFinished = 1
	task.ClaimTime = time.Now()
	tx := model.GetTx()
	task.UpdateTx(tx)
	newTask := &model.Task{
		Identity: model.Identity{
			ID: util2.NewSnowflakeId(),
		},
		NodeID:        "开始",
		ProcInstID:    procInst.ID,
		IsFinished:    1,
		ClaimTime:     time.Now(),
		Level:         0,
		MemberCount:   1,
		UnCompleteNum: 0,
		ActType:       "or",
		AgreeNum:      1,
	}
	newTask.UpdateTx(tx)
	procInst.TaskID = newTask.ID
	procInst.State = flow.WITHDRAW
	procInst.UpdateTx(tx)
	return nil
}

// WithDrawTask 撤回任务
//func WithDrawTask(taskID int64, procInstID int64, userID int64, username string, dataId int64, targetId int64, comment string) error {
//	var err1, err2 error
//	var currentTask, lastTask *model.Task
//	var timesx time.Time
//	var wg sync.WaitGroup
//	timesx = time.Now()
//	wg.Add(2)
//	go func() {
//		currentTask, err1 = GetTaskByID(taskID)
//		wg.Done()
//	}()
//	go func() {
//		lastTask, err2 = GetTaskLastByProInstID(procInstID)
//		wg.Done()
//	}()
//	wg.Wait()
//	if err1 != nil {
//		if err1 == gorm.ErrRecordNotFound {
//			return errors.New("任务不存在")
//		}
//		return err1
//	}
//	if err2 != nil {
//		if err2 == gorm.ErrRecordNotFound {
//			return errors.New("找不到流程实例id为【" + fmt.Sprintf("%d", procInstID) + "】的任务，无权撤回")
//		}
//		return err2
//	}
//	// str1,_:=util.ToJSONStr(currentTask)
//	// str2,_:=util.ToJSONStr(lastTask)
//	// fmt.Println(str1)
//	// fmt.Println(str2)
//	if currentTask.Level == 0 {
//		return errors.New("开始位置无法撤回")
//	}
//	//if lastTask.Assignee != userID {
//	//	return errors.New("只能撤回本人审批过的任务！！")
//	//}
//	if currentTask.IsFinished == 1 {
//		return errors.New("已经审批结束，无法撤回！")
//	}
//	if currentTask.UnCompleteNum != currentTask.MemberCount {
//		return errors.New("已经有人审批过了，无法撤回！")
//	}
//	sub := currentTask.Level - lastTask.Level
//	if math.Abs(float64(sub)) != 1 {
//		return errors.New("只能撤回相邻的任务！！")
//	}
//	//var pass = false
//	//if sub < 0 {
//	//	pass = true
//	//}
//	fmt.Printf("判断是否可以撤回,耗时：%v\n", time.Since(timesx))
//	timesx = time.Now()
//	tx := model.GetTx()
//	// 更新当前的任务
//	currentTask.IsFinished = 1
//	err := currentTask.UpdateTx(tx)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//	// 撤回
//	err = MoveStageByProcInstID(userID, username, dataId, targetId, comment, currentTask.ID, procInstID, currentTask.Level, 5, tx)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	fmt.Printf("撤回流程耗时：%v\n", time.Since(timesx))
//	return nil
//}

// MoveStageByProcInstID MoveStageByProcInstID
func MoveStageByProcInstID(userID int64, username string, dataId int64, targetId int64, comment string, taskID int64, procInstID int64, level int, result int, tx *gorm.DB) (err error) {
	// TODO entsql SELECT node_infos FROM `act_execution`  WHERE (proc_inst_id=5469563801642631168)
	nodeInfos, err := GetExecNodeInfosByProcInstID(procInstID)
	if err != nil {
		return err
	}
	return MoveStage(nodeInfos, userID, username, dataId, targetId, comment, taskID, procInstID, level, result, tx)
}

// MoveStage 流程流转
func MoveStage(nodeInfos []*flow.NodeInfo, userID int64, username string, dataId int64, targetId int64, comment string, taskID int64, procInstID int64, level int, result int, tx *gorm.DB) (err error) {
	// 添加上一步的参与人
	//err = AddParticipantTx(userID, username, dataId, comment, taskID, procInstID, level, tx)
	//AddCandidate(nodeInfos, taskID, procInstID, tx)
	if err != nil {
		return err
	}
	if result != 0 {
		level++
		if level-1 > len(nodeInfos) {
			return errors.New("已经结束无法流转到下一个节点")
		}
	} else {
		level--
		if level < 0 {
			return errors.New("处于开始位置，无法回退到上一个节点")
		}
	}
	// 指定下一步执行人
	//if len(candidate) > 0 {
	//	nodeInfos[step].Aprover = candidate
	//}
	// 判断下一流程： 如果是审批人是：抄送人
	// fmt.Printf("下一审批人类型：%s\n", nodeInfos[step].AproverType)
	// fmt.Println(nodeInfos[step].AproverType == flow.NodeTypes[flow.NOTIFIER])
	if nodeInfos[level].AproverType == flow.NodeTypes[flow.NOTIFIER] {
		// 生成新的任务
		var task = model.Task{
			Identity: model.Identity{
				ID: util2.NewSnowflakeId(),
			},
			NodeID:     flow.NodeTypes[flow.NOTIFIER],
			Level:      level,
			ProcInstID: procInstID,
			IsFinished: const2.IsFinishYes,
		}
		//task.IsFinished = 1

		_, err := task.NewTaskTx(tx)
		if err != nil {
			return err
		}
		//添加候选人
		/*TODO entsql
		INSERT INTO `act_candidate`
		(`id`,`create_time`,`user_id`,`user_name`,`proc_inst_id`,`target_id`,`task_id`,`is_finish`)
		VALUES (5893408212094189568,'2022-08-30 15:25:25',107777777,'大壮',5469563801642631168,1727882,9212524636351135769,2)*/
		err = AddCandidate(nodeInfos[level], taskID, targetId, procInstID, tx)
		if err != nil {
			return err
		}
		// 添加抄送人
		//err = AddNotifierTx(dataId, level, procInstID, tx)
		//if err != nil {
		//	return err
		//}
		return MoveStage(nodeInfos, userID, username, dataId, targetId, comment, taskID, procInstID, level, result, tx)
	}
	if result == flow.HAVEPASS {
		return MoveToNextStage(nodeInfos, targetId, procInstID, level, tx)
	}
	return MoveToPrevStage(nodeInfos, userID, dataId, targetId, procInstID, level, tx)
}

// MoveToNextStage MoveToNextStage
//通过
func MoveToNextStage(nodeInfos []*flow.NodeInfo, targetId int64, procInstID int64, level int, tx *gorm.DB) error {
	var task = getNewTask(nodeInfos, level, procInstID) //新任务
	var procInst = &model.ProcInst{                     // 流程实例要更新的字段
		NodeID: nodeInfos[level].NodeID,
	}
	if (level + 1) != len(nodeInfos) { // 下一步不是【结束】
		/*TODO entsql
		INSERT INTO `act_task`
		(`create_time`,`node_id`,`level`,`proc_inst_id`,`claim_time`,`member_count`,
		`un_complete_num`,`agree_num`,`act_type`)
		VALUES ('2022-08-30 15:25:25','1659320226296',2,5469563801642631168,'2022-08-30 15:25:25',1,1,0,'or')
		*/
		//TODO 在哪里调用，为什么要调用 entsql SELECT `is_finished` FROM `act_task`  WHERE (id = 9212524636351135769)
		// 生成新的任务
		taskID, err := task.NewTaskTx(tx)
		if err != nil {
			return err
		}
		// 添加candidate group
		/*TODO entsql
		INSERT INTO `act_candidate`
		(`id`,`create_time`,`user_id`,`user_name`,`proc_inst_id`,`target_id`,`task_id`,`is_finish`)
		VALUES (5893408212094189568,'2022-08-30 15:25:25',107777777,'大壮',5469563801642631168,1727882,9212524636351135769,2)*/
		//err = AddCandidateTx(dataId, step, taskID, procInstID, tx)
		err = AddCandidate(nodeInfos[level], taskID, targetId, procInstID, tx)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		// 更新流程实例
		procInst.TaskID = taskID
		task.IsFinished = const2.IsFinishNo
		procInst.IsFinished = const2.IsFinishNo
		if level > 1 {
			procInst.State = flow.DEALING
		}
		/*TODO entsql
		UPDATE `act_proc_inst`
		SET `is_finished` = 2, `node_id` = '1659320226296', `state` = 2, `task_id` = 9212524636351135769*/
		err = UpdateProcInst(procInst, tx)
		if err != nil {
			return err
		}
	} else { // 最后一步直接结束
		// 生成新的任务
		task.IsFinished = const2.IsFinishYes
		task.ClaimTime = time.Now()
		/*TODO entsql
		INSERT INTO `act_task` (`create_time`,`node_id`,`level`,`proc_inst_id`,`claim_time`,`agree_num`,`is_finished`) VALUES ('2022-08-30 15:25:51','结束',3,5469563801642631168,'2022-08-30 15:25:51',0,1)
		*/
		//TODO entsql   SELECT `member_count`, `un_complete_num`, `act_type` FROM `act_task`  WHERE (id = 9212524636351135770)
		taskID, err := task.NewTaskTx(tx)
		if err != nil {
			return err
		}
		// 删除候选用户组
		/*TODO entsql
		DELETE FROM `act_identity_link`  WHERE (proc_inst_id=5469563801642631168 and type='candidate')*/
		err = DelCandidateByProcInstID(procInstID, tx)
		if err != nil {
			return err
		}
		// 更新流程实例
		procInst.TaskID = taskID
		procInst.EndTime = time.Now()
		procInst.IsFinished = const2.IsFinishYes
		procInst.State = flow.HAVEPASS
		//procInst.Candidate = "审批结束"
		/* TODO entsql
		UPDATE `act_proc_inst`
		SET `end_time` = '2022-08-30 15:25:51', `is_finished` = 1, `node_id` = '结束',
		`state` = 7, `task_id` = 9212524636351135770
		*/
		err = UpdateProcInst(procInst, tx)
		if err != nil {
			return err
		}
		//删除待办数据
		go func() {
			candidate := model.Candidate{}
			// TODO entsql DELETE FROM `act_candidate`  WHERE (proc_inst_id=5469563801642631168 )
			err = candidate.DelByProcInstID(procInstID, tx)
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

// MoveToPrevStage MoveToPrevStage
// 驳回
func MoveToPrevStage(nodeInfos []*flow.NodeInfo, userID int64, dataId int64, targetId int64, procInstID int64, step int, tx *gorm.DB) error {
	// 生成新的任务
	var task = getNewTask(nodeInfos, step, procInstID) //新任务
	taksID, err := task.NewTaskTx(tx)
	if err != nil {
		return err
	}
	var procInst = &model.ProcInst{ // 流程实例要更新的字段
		NodeID: nodeInfos[step].NodeID,
		//Candidate: nodeInfos[step].Aprover,
		TaskID: taksID,
	}
	procInst.ID = procInstID
	err = UpdateProcInst(procInst, tx)
	if err != nil {
		return err
	}
	if step == 0 { // 流程回到起始位置，注意起始位置为0,
		err = AddCandidateUserTx(userID, dataId, step, taksID, procInstID, tx)
		if err != nil {
			return err
		}
		return nil
	}
	// 添加candidate group
	err = AddIdentityLinkCandidateTx(targetId, step, taksID, procInstID, tx)
	if err != nil {
		return err
	}
	return nil
}
func getNewTask(nodeInfos []*flow.NodeInfo, level int, procInstID int64) *model.Task {
	var task = &model.Task{ // 新任务
		NodeID: nodeInfos[level].NodeID,
		Level:  level,
		//CreateTime:    currentTime,
		ProcInstID:    procInstID,
		MemberCount:   nodeInfos[level].MemberCount,
		UnCompleteNum: nodeInfos[level].MemberCount,
		ActType:       nodeInfos[level].ActType,
	}
	return task
}
