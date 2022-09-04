package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"act/common/act"
	"act/common/act/execution"
	"act/common/act/identitylink"
	"act/common/act/procinst"
	"act/common/act/task"
	"act/rpc/internal/flow"
	"act/rpc/internal/svc"
	const2 "act/rpc/internal/workflow-const"
	"act/rpc/ms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveIdentityLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveIdentityLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveIdentityLinkLogic {
	return &SaveIdentityLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程审批
func (l *SaveIdentityLinkLogic) SaveIdentityLink(in *ms.IdentityLinkReq) (*ms.IdentityLinkReply, error) {
	// todo: add your logic here and delete this line
	log.Println("SaveIdentityLink 流程开始------rpc 222222")
	log.Println("req", in)

	store := l.svcCtx.ActStore

	/*
		1.taskId
		entsql SELECT * FROM `act_proc_inst`  WHERE (data_id =1 ) ORDER BY `act_proc_inst`.`id` ASC LIMIT 1
	*/
	// TODO dataID
	procInstQuery, err := store.ProcInst.Query().
		Where(procinst.DataIDEQ(0)).
		Limit(1).Order(act.Asc(procinst.FieldID)).
		Only(l.ctx)
	// procInstQuery, err := store.ProcInst.Query().Where(procinst.DataIDEQ(in.dataID)).Limit(1).Only(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("procInstQuery", procInstQuery)
	taskId := procInstQuery.TaskID
	log.Println("taskId", taskId)

	/*
		2.查询 用户是否包含在审批人中
		entsql SELECT count(*) FROM `act_candidate`  WHERE ( task_id=9212524636351135768  and user_id=122298283)
	*/

	/*
		3.
		entsql SELECT count(*) FROM `act_identity_link`  WHERE ( task_id=9212524636351135768  and user_id=122298283)
	*/
	identityLinkCount, err := store.IdentityLink.Query().
		Where(identitylink.TaskIDEQ(int64(taskId)), identitylink.UserIDEQ(in.UserId)).Count(l.ctx)
	if err != nil {
		return nil, err
	}
	flag := false

	if identityLinkCount > 0 {
		flag = true
	}
	if flag {
		return &ms.IdentityLinkReply{Result: "不能重复审批"}, nil
	}

	// 4.将审批驳回与通过/不通过分开
	switch in.Result {
	case flow.REJECTTOPRE:
		// TODO
		err = l.Complete(int64(taskId), in.UserId, in.UserName, 1, in.TargetId, in.Comment, int(in.Result))

	case flow.REJECTTOSTART:
		// TODO
		err = l.Complete(int64(taskId), in.UserId, in.UserName, 1, in.TargetId, in.Comment, int(in.Result))

	default:
		err = l.Complete(int64(taskId), in.UserId, in.UserName, 1, in.TargetId, in.Comment, int(in.Result))
	}
	if err != nil {
		return nil, err
	}

	return &ms.IdentityLinkReply{}, nil

}

// 7个参数
func (l *SaveIdentityLinkLogic) Complete(taskID int64, userID int64, username string, dataId int64, targetId int64, comment string, result int) error {
	var err error
	mtx := sync.Mutex{}
	mtx.Lock()
	defer mtx.Unlock()
	store, err := l.svcCtx.ActStore.Tx(l.ctx)
	if err != nil {
		return err
	}

	/*
		1.更新任务
		1.1查询任务 entsql SELECT * FROM `act_task`  WHERE (id=9212524636351135768)
	*/
	taskQuery, err := store.Task.Query().Where(task.IDEQ(int(taskID))).Only(l.ctx)
	if err != nil {
		store.Rollback()
		return err
	}
	if taskQuery == nil {
		return errors.New("任务【" + fmt.Sprintf("%d", taskQuery.ID) + "】不存在")
	}
	// 1.2判断是否已经结束
	if taskQuery.IsFinished == int8(const2.IsFinishYes) {
		if taskQuery.NodeID == "结束" {
			return errors.New("流程已经结束")
		}
		return errors.New("任务【" + fmt.Sprintf("%d", taskID) + "】已经被审批过了！！")
	}
	//同意
	if result == flow.HAVEPASS {
		taskQuery.AgreeNum++
	}
	// 未审批人数减一
	taskQuery.UnCompleteNum--
	// 判断是否结束
	if taskQuery.UnCompleteNum == 0 {
		taskQuery.IsFinished = int8(const2.IsFinishYes)
	}

	/*
		 1.2 更新任务,更新4个字段
		 entsql
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
	// TODO
	err = store.Task.Update().
		SetAgreeNum(taskQuery.AgreeNum).SetIsFinished(taskQuery.IsFinished).SetUnCompleteNum(taskQuery.UnCompleteNum).
		SetClaimTime(time.Now()).SetCreateTime(time.Now()).
		Exec(l.ctx)
	if err != nil {
		store.Rollback()
		return err
	}
	/*
		2.如果是会签
		entsql  SELECT count(*) FROM `act_identity_link`  WHERE (user_id=122298283  and task_id=9212524636351135768)
	*/
	if taskQuery.ActType == "and" {
		identityLinkCount, err := store.IdentityLink.Query().
			Where(identitylink.UserIDEQ(userID), identitylink.TargetIDEQ(taskID)).
			Count(l.ctx)
		if err != nil {
			store.Rollback()
			return err
		}
		if identityLinkCount > 0 {
			return errors.New("您已经审批过了，请等待他人审批！）")
		}

	}
	// 3.查看任务的未审批人数是否为0，不为0就不流转
	/* entsql
	INSERT INTO `act_identity_link`
	(`id`,`create_time`,`type`,`user_id`,`user_name`,`task_id`,`target_id`,`step`,`proc_inst_id`,`comment`,`result`)
	VALUES (5893377269455290368,'2022-08-30 15:25:25','participant',122298283,'小春',9212524636351135768,1727882,1,5469563801642631168,'小春评论1111',7)
	*/
	if result != 0 {
		_, err := store.IdentityLink.Create().
			SetCreateTime(time.Now()).SetUserID(userID).SetUserName(username).
			SetTaskID(taskID).SetTargetID(targetId).SetStep(taskQuery.Step).
			SetProcInstID(taskQuery.ProcInstID).SetComment(comment).SetResult(result).
			Save(l.ctx)
		if err != nil {
			store.Rollback()
			return err
		}
	}
	// 4.流转到下一流程
	/*
	 entsql SELECT node_infos FROM `act_execution`  WHERE (proc_inst_id=5469563801642631168)
	*/
	executionQuery, err := store.Execution.Query().Where(execution.ProcInstID(taskQuery.ProcInstID)).
		Only(l.ctx)
	if err != nil {
		store.Rollback()
		return err
	}
	err = store.Commit()
	if err != nil {
		store.Rollback()
		return err
	}
	nodeInfos := executionQuery.NodeInfos
	log.Println("nodeInfos", nodeInfos)
	// TODO nodeInfos
	l.MoveStage(nil, userID, username, dataId, targetId, comment, targetId, executionQuery.ProcInstID, taskQuery.Level, result)
	return nil
}
func (l *SaveIdentityLinkLogic) MoveStage(nodeInfos []*flow.NodeInfo, userID int64, username string, dataId int64,
	targetId int64, comment string, taskID int64, procInstID int64, level int, result int) (err error) {
	// 添加上一步的参与人
	//err = AddParticipantTx(userID, username, dataId, comment, taskID, procInstID, level, tx)
	//AddCandidate(nodeInfos, taskID, procInstID, tx)
	// if err != nil {
	// 	return err
	// }

	//1.前置判断：已经结束无法流转到下一个节点；处于开始位置，无法回退到上一个节点
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
		_, err = l.svcCtx.ActStore.Task.Create().
			SetNodeID(flow.NodeTypes[flow.NOTIFIER]).SetLevel(level).
			SetProcInstID(procInstID).SetIsFinished(int8(const2.IsFinishYes)).
			SetCreateTime(time.Now()).SetNodeID("开始").SetLevel(0).
			SetClaimTime(time.Now()).SetMemberCount(1).SetAgreeNum(1).SetActType("or").
			Save(l.ctx)
		if err != nil {
			return err
		}

		//TODO 添加候选人
		// err = AddCandidate(nodeInfos[level], taskID, targetId, procInstID, tx)
		// if err != nil {
		// 	return err
		// }
		// 添加抄送人
		//err = AddNotifierTx(dataId, level, procInstID, tx)
		//if err != nil {
		//	return err
		//}
		return l.MoveStage(nodeInfos, userID, username, dataId, targetId, comment, taskID, procInstID, level, result)
	}
	if result == flow.HAVEPASS {
		return l.MoveToNextStage(nodeInfos, targetId, procInstID, level)
	}
	return l.MoveToPrevStage(nodeInfos, userID, dataId, targetId, procInstID, level)
}

// MoveToNextStage MoveToNextStage
//通过
func (l *SaveIdentityLinkLogic) MoveToNextStage(nodeInfos []*flow.NodeInfo,
	targetId int64, procInstID int64, level int) error {
	tx, err := l.svcCtx.ActStore.Tx(l.ctx)
	if err != nil {
		return err
	}
	//新任务
	taskQuery := l.svcCtx.ActStore.Task.Query().Where(task.LevelEQ(level), task.ProcInstIDEQ(procInstID))
	if taskQuery == nil {
		return errors.New("err  task")
	}
	// 下一步不是【结束】
	if (level + 1) != len(nodeInfos) {
		// 生成新的任务
		/* entsql
		INSERT INTO `act_task`
		(`create_time`,`node_id`,`level`,`proc_inst_id`,`claim_time`,`member_count`,
		`un_complete_num`,`agree_num`,`act_type`)
		VALUES ('2022-08-30 15:25:25','1659320226296',2,5469563801642631168,'2022-08-30 15:25:25',1,1,0,'or')
		*/
		taskCreate, err := tx.Task.Create().
			SetNodeID(nodeInfos[level].NodeID).SetLevel(level).SetCreateTime(time.Now()).
			SetProcInstID(procInstID).SetMemberCount(nodeInfos[level].MemberCount).
			SetUnCompleteNum(nodeInfos[level].MemberCount).SetActType(task.ActType(nodeInfos[level].ActType)).
			SetCreateTime(time.Now()).SetClaimTime(time.Now()).
			Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
		taskID := taskCreate.ID
		//TODO 在哪里调用，为什么要调用 entsql SELECT `is_finished` FROM `act_task`  WHERE (id = 9212524636351135769)

		// 添加candidate group
		/*TODO entsql
		INSERT INTO `act_candidate`
		(`id`,`create_time`,`user_id`,`user_name`,`proc_inst_id`,`target_id`,`task_id`,`is_finish`)
		VALUES (5893408212094189568,'2022-08-30 15:25:25',107777777,'大壮',5469563801642631168,1727882,9212524636351135769,2)*/
		//err = AddCandidateTx(dataId, step, taskID, procInstID, tx)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		// 更新流程实例,流程实例要更新的字段
		/*
			 entsql
			UPDATE `act_proc_inst`
			SET `is_finished` = 2, `node_id` = '1659320226296', `state` = 2, `task_id` = 9212524636351135769
		*/
		updateRow, err := tx.ProcInst.Update().
			SetNodeID(nodeInfos[level].NodeID).SetTaskID(int64(taskID)).
			SetState(flow.DEALING).SetIsFinished(int8(const2.IsFinishNo)).
			Save(l.ctx)

		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		if updateRow < 0 {
			return err
		}
	} else { // 最后一步直接结束
		// 生成新的任务
		/* entsql
		INSERT INTO `act_task`
		 (`create_time`,`node_id`,`level`,`proc_inst_id`,`claim_time`,`agree_num`,`is_finished`)
		 VALUES ('2022-08-30 15:25:51','结束',3,5469563801642631168,'2022-08-30 15:25:51',0,1)
		*/
		//TODO entsql   SELECT `member_count`, `un_complete_num`, `act_type` FROM `act_task`  WHERE (id = 9212524636351135770)
		taskCreate, err := tx.Task.Create().
			SetCreateTime(time.Now()).SetNodeID(nodeInfos[level].NodeID).SetLevel(level).
			SetProcInstID(procInstID).SetClaimTime(time.Now()).SetAgreeNum(0).SetIsFinished(int8(const2.IsFinishYes)).
			Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
		log.Println("else taskCreate", taskCreate)
		taskID := taskCreate.ID
		// 删除候选用户组
		// err = DelCandidateByProcInstID(procInstID, tx)
		// if err != nil {
		// 	return err
		// }
		// 更新流程实例
		// procInst.TaskID = taskID
		// procInst.EndTime = time.Now()
		// procInst.IsFinished = const2.IsFinishYes
		// procInst.State = flow.HAVEPASS
		// TODO procInst.Candidate = "审批结束"
		// TODO err = UpdateProcInst(procInst, tx)

		/* TODO entsql
		UPDATE `act_proc_inst`
		SET `end_time` = '2022-08-30 15:25:51', `is_finished` = 1, `node_id` = '结束',
		`state` = 7, `task_id` = 9212524636351135770
		*/
		updateRow, err := tx.ProcInst.Update().
			SetEndTime(time.Now()).SetIsFinished(int8(const2.IsFinishYes)).SetNodeID(nodeInfos[level].NodeID).
			SetState(flow.HAVEPASS).SetTaskID(int64(taskID)).
			Save(l.ctx)

		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		if updateRow == 0 {
			return err
		}
		//删除待办数据
		// go func() {
		// 	candidate := model.Candidate{}
		// 	err = candidate.DelByProcInstID(procInstID, tx)
		// }()
		// if err != nil {
		// 	return err
		// }
	}
	return nil
}

// MoveToPrevStage MoveToPrevStage
// TODO驳回
func (l *SaveIdentityLinkLogic) MoveToPrevStage(nodeInfos []*flow.NodeInfo, userID int64, dataId int64,
	targetId int64, procInstID int64, step int) error {
	// 生成新的任务
	// TODO var task = getNewTask(nodeInfos, step, procInstID) //新任务
	// taksID, err := task.NewTaskTx(tx)
	tx, err := l.svcCtx.ActStore.Tx(l.ctx)
	if err != nil {
		return err
	}
	taskQueryArr, err := l.svcCtx.ActStore.Task.Query().
		Where(task.StepEQ(step), task.ProcInstIDEQ(procInstID)).
		Where(task.NodeIDEQ(nodeInfos[step].NodeID), task.MemberCountEQ(nodeInfos[step].MemberCount)).
		Where(task.UnCompleteNumEQ(nodeInfos[step].MemberCount)).
		Where(task.ActTypeEQ(task.ActType(nodeInfos[step].ActType))).
		IDs(l.ctx)
	if err != nil {
		return err
	}
	if len(taskQueryArr) == 0 {
		return errors.New("err  MoveToPrevStage")
	}
	log.Println("taskQueryArr", taskQueryArr)
	taskID := taskQueryArr[0]
	// var procInst = &model.ProcInst{ // 流程实例要更新的字段
	// 	NodeID: nodeInfos[step].NodeID,
	// 	//Candidate: nodeInfos[step].Aprover,
	// 	TaskID: taksID,
	// }
	// procInst.ID = procInstID
	// TODO  err = UpdateProcInst(procInst, tx)
	// TODO // 	//Candidate: nodeInfos[step].Aprover,
	// nodeInfosStr:=convertor.ToString(nodeInfos)
	updateRow, err := tx.ProcInst.Update().
		SetNodeID(string(nodeInfos[step].NodeID)).SetTaskID(int64(taskID)).SetProcDefID(procInstID).
		SetEndTime(time.Now()).SetState(flow.HAVEPASS).
		Save(l.ctx)

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	if updateRow == 0 {
		return err
	}
	//TODO if step == 0 { // 流程回到起始位置，注意起始位置为0,
	// 	err = AddCandidateUserTx(userID, dataId, step, taksID, procInstID, tx)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }
	// //TODO 添加candidate group
	// err = AddIdentityLinkCandidateTx(targetId, step, taksID, procInstID, tx)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// // 流程开始
// func (l *SaveIdentityLinkLogic) SaveIdentityLink(in *ms.IdentityLinkReq) (*ms.IdentityLinkReply, error) {
// 	// todo: add your logic here and delete this line
// 	log.Println("SaveIdentityLink 流程开始------rpc 222222")
// 	log.Println("req", in)
// 	var err error

// 	/*
// 		1. SELECT * FROM `act_proc_def`
// 		WHERE (yewu_form_id='1239348' and target_id=1727882)
// 		ORDER BY version desc
// 	*/
// 	// 根据yewu_form_id和target_id从act_proc_def表中，获取id
// 	store := l.svcCtx.ActStore

// 	procDefResult := store.ProcDef.Query().
// 		Where(procdef.YewuFormID("1239348"), procdef.TargetIDEQ(1727882)).
// 		Order(act.Asc(procdef.FieldVersion))

// 	procdefIdArr, err := procDefResult.IDs(l.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	procdefId := procdefIdArr[0]
// 	log.Println("procdefId", procdefId)

// 	// node解析
// 	node := &flow.Node{}
// 	// TODO 根据id查询resource
// 	var resourceStr string = `{"name":"发起人","type":"start","nodeId":"sid-startevent","childNode":{"type":"route","nodeId":"1659320135116","prevId":"sid-startevent","childNode":{"name":"审批人","type":"approver","nodeId":"1659320226296","prevId":"1659320135116","properties":{"actionerRules":[{"type":"target_management","approverNames":"大壮","approverIds":"107777777","actType":"or","level":2,"autoUp":true}]}},"conditionNodes":[{"name":"条件1","type":"condition","nodeId":"1659320135126","prevId":"1659320135116","childNode":{"name":"审批人","type":"approver","nodeId":"1659320194369","prevId":"1659320135126","properties":{"actionerRules":[{"type":"target_management","approverNames":"小春","approverIds":"122298283","actType":"and","level":1,"autoUp":true}]}},"properties":{"conditions":[[{"type":"dingtalk_actioner_range_condition","paramKey":"amount","paramLabel":"5000","upperBound":"3000"}]]}},{"name":"条件2","type":"condition","nodeId":"1659320135136","prevId":"1659320135126","childNode":{"name":"审批人","type":"approver","nodeId":"1659320201085","prevId":"1659320135136","properties":{"actionerRules":[{"type":"target_label","approverNames":"小龙","approverIds":"122298287","memberCount":1,"actType":"or"}]}},"properties":{"conditions":[[{"type":"dingtalk_actioner_range_condition","paramKey":"amount2","paramLabel":"5000","lowerBoundEqual":"3000"}]]}}]}}`
// 	procDefResource := act.ProcDef{Resource: resourceStr}
// 	err = util.Str2Struct(procDefResource.Resource, node)
// 	if err != nil {
// 		return nil, err
// 	}
// 	level := 0 // 0 为开始节点
// 	log.Println("level", level)
// 	// ---------------------------------------------
// 	// 新建流程实例
// 	/*	 2.
// 	INSERT INTO `act_proc_inst`
// 	(`id`,`create_time`,`proc_def_id`,`title`,`node_id`,`task_id`,
// 	`start_time`,`end_time`,`start_user_id`,`start_user_name`,`target_id`,`data_id`,`state`)
// 	VALUES (6817417568194756608,'2022-08-19 17:08:33',8725662109012164608,'','',0,
// 	'2022-08-19 17:08:33','2022-08-19 17:08:33',12025,'yxq',1727882,0,1)
// 	*/
// 	procInstInfo, err := store.ProcInst.Create().SetCreateTime(time.Now()).SetProcDefID(11).SetTitle("111").
// 		SetNodeID("111").SetTaskID(11).SetStartTime(time.Now()).SetEndTime(time.Now().AddDate(0, 0, 1)).SetStartUserID(12025).SetStartUserName("ling").
// 		SetTargetID(1727882).SetDataID(0).SetState(1).
// 		Save(l.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	procInstId := procInstInfo.ID
// 	log.Println("procInstId", procInstId)
// 	/* 3.
// 	// 生成执行流，一串运行节点
// 	INSERT INTO `act_execution`
// 	(`id`,`create_time`,`proc_inst_id`,`proc_def_id`,`is_active`,`start_time`,`node_infos`)
// 	VALUES (6817391092489617408,'2022-08-19 17:08:33',6817417568194756608,8725662109012164608,
// 	1,'2022-08-19 17:08:33','[{"nodeId":"开始","type":"starter","aproverType":"","memberCount":0,"level":0,"actType":"","approverNames":"","approverIds":"⻹"},{"nodeId":"1659320194369","type":"target_management","aproverType":"approver","memberCount":1,"level":1,"actType":"and","approverNames":"小春","approverIds":"122298283"},{"nodeId":"1659320226296","type":"target_management","aproverType":"approver","memberCount":1,"level":2,"actType":"or","approverNames":"大壮","approverIds":"107777777"},{"nodeId":"结束","type":"","aproverType":"","memberCount":0,"level":3,"actType":"","approverNames":"","approverIds":""}]')
// 	*/
// 	execution, err := store.Execution.Create().SetCreateTime(time.Now()).
// 		SetProcInstID(int64(random.RandInt(1, 100000000000))).
// 		SetProcDefID(8725662109012164608).SetIsActive(1).
// 		SetStartTime(time.Now()).SetNodeInfos(resourceStr).
// 		Save(l.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("execution", execution)
// 	/* 4.
// 	// -----------------生成新任务-------------
// 	INSERT INTO `act_task`
// 	(`id`,`create_time`,`node_id`,`level`,`proc_inst_id`,`claim_time`,
// 	`member_count`,`agree_num`,`act_type`,`is_finished`)
// 	VALUES (6817391092489617408,'2022-08-19 17:08:33','开始',0,6817417568194756608,'2022-08-19 17:08:33',
// 	1,1,'or',1)
// 	*/
// 	_, err = store.Task.Create().
// 		SetCreateTime(time.Now()).SetNodeID("开始").SetLevel(0).SetProcInstID(6817417568194756608).
// 		SetClaimTime(time.Now()).SetMemberCount(1).SetAgreeNum(1).SetActType("or").SetIsFinished(1).
// 		Save(l.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// TODO
// 	// SELECT `un_complete_num` FROM `act_task`  WHERE (id = 6817391092489617408)
// 	// SELECT `is_finished` FROM `act_task`  WHERE (id = 9212524636351135751)
// 	taskFinsh := store.Task.Query().Where(task.ID(6817391092489617408))
// 	taskFinsh.IDs(l.ctx)
// 	/* 5.

// 	INSERT INTO `act_candidate`
// 	(`id`,`create_time`,`user_id`,`user_name`,`proc_inst_id`,`target_id`,`task_id`,`is_finish`)
// 	VALUES (6817318104469700608,'2022-08-19 17:08:33',122298283,'小春',
// 	6817417568194756608,1727882,9212524636351135751,2)
// 	*/
// 	/* 6.
// 	// 流程移动到下一环节
// 	UPDATE `act_proc_inst`
// 	SET `is_finished` = 2, `node_id` = '1659320194369', `task_id` = 9212524636351135751
// 	*/
// 	// 6.1 根据id和ProcDefID查询获取流程实例
// 	procdefList, err := store.ProcInst.Query().Where(procinst.ID(826790988067274752), procinst.ProcDefID(800650655701041152)).First(l.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("procdefList", procdefList)
// 	// 6.2 更新
// 	procdefUpdate, err := procdefList.Update().SetIsFinished(2).SetNodeID("1659320194369").SetTaskID(9212524636351135751).Save(l.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("procdefUpdate", procdefUpdate)

// 	// TODO下一个节点 同意 MoveToNextStage或驳回MoveToPrevStage
// 	return &ms.IdentityLinkReply{}, nil
// }
