package logic

import (
	"context"
	"errors"
	"log"
	"time"

	"act/common/act/task"
	"act/rpc/internal/flow"
	"act/rpc/internal/model"
	"act/rpc/internal/svc"
	const2 "act/rpc/internal/workflow-const"
	"act/rpc/ms"

	// "orginone/common/schema/task"

	// "github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveExecutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveExecutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExecutionLogic {
	return &SaveExecutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程审批
func (l *SaveExecutionLogic) SaveExecution(in *ms.ExecutionReq) (*ms.ExecutionReply, error) {
	// todo: add your logic here and delete this line
	log.Println("SaveExecution 流程审批")
	// TODO 传入参数
	// err := l.MoveStage()
	// if err != nil {
	// 	return nil, err
	// }
	return &ms.ExecutionReply{}, nil
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
		// var task = model.Task{
		// 	Identity: model.Identity{
		// 		ID: util2.NewSnowflakeId(),
		// 	},
		// 	NodeID:     flow.NodeTypes[flow.NOTIFIER],
		// 	Level:      level,
		// 	ProcInstID: procInstID,
		// 	IsFinished: const2.IsFinishYes,
		// }
		//task.IsFinished = 1
		err := l.svcCtx.ActStore.IdentityLink.Create().
			SetCreateTime(time.Now()).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		_, err = l.svcCtx.ActStore.Task.Create().
			SetNodeID(flow.NodeTypes[flow.NOTIFIER]).SetLevel(level).
			SetProcInstID(procInstID).SetIsFinished(int8(const2.IsFinishYes)).
			SetCreateTime(time.Now()).SetNodeID("开始").SetLevel(0).
			SetClaimTime(time.Now()).SetMemberCount(1).SetAgreeNum(1).SetActType("or").
			Save(l.ctx)
		if err != nil {
			return err
		}
		// _, err := task.NewTaskTx(tx)
		// if err != nil {
		// 	return err
		// }
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
	// TODO var task = getNewTask(nodeInfos, level, procInstID) //新任务
	// TODO task.NodeIDEQ(nodeInfos)
	taskQuery := l.svcCtx.ActStore.Task.Query().Where(task.LevelEQ(level), task.ProcInstIDEQ(procInstID))
	if taskQuery == nil {
		return errors.New("err  task")
	}
	var procInst = &model.ProcInst{ // 流程实例要更新的字段
		NodeID: nodeInfos[level].NodeID,
	}
	if (level + 1) != len(nodeInfos) { // 下一步不是【结束】
		// 生成新的任务
		//  TODO taskID, err := task.NewTaskTx(tx)
		taskCreate, err := l.svcCtx.ActStore.Task.Create().
			SetCreateTime(time.Now()).SetClaimTime(time.Now()).
			Save(l.ctx)
		if err != nil {
			return err
		}
		taskID := taskCreate.ID
		// 添加candidate group
		//err = AddCandidateTx(dataId, step, taskID, procInstID, tx)
		// TODO err = AddCandidate(nodeInfos[level], taskID, targetId, procInstID, tx)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		// 更新流程实例
		procInst.TaskID = int64(taskID)
		// task.IsFinished = const2.IsFinishNo
		procInst.IsFinished = const2.IsFinishNo
		if level > 1 {
			procInst.State = flow.DEALING
		}
		// TODO err = UpdateProcInst(procInst, tx)
		updateRow, err := l.svcCtx.ActStore.ProcInst.Update().
			SetNodeID(nodeInfos[level].NodeID).SetTaskID(int64(taskID)).SetState(flow.DEALING).SetIsFinished(int8(const2.IsFinishNo)).
			Save(l.ctx)

		if err != nil {
			return err
		}
		if updateRow < 0 {
			return err
		}
	} else { // 最后一步直接结束
		// 生成新的任务
		// task.IsFinished = const2.IsFinishYes
		// task.ClaimTime = time.Now()
		// TODO taskID, err := task.NewTaskTx(tx)
		taskCreate, err := l.svcCtx.ActStore.Task.Create().
			SetIsFinished(int8(const2.IsFinishYes)).SetClaimTime(time.Now()).SetCreateTime(time.Now()).
			SetCreateTime(time.Now()).SetClaimTime(time.Now()).
			Save(l.ctx)
		if err != nil {
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
		updateRow, err := l.svcCtx.ActStore.ProcInst.Update().
			SetTaskID(int64(taskID)).SetEndTime(time.Now()).SetState(flow.HAVEPASS).
			SetIsFinished(int8(const2.IsFinishYes)).
			Save(l.ctx)

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
// 驳回
func (l *SaveIdentityLinkLogic) MoveToPrevStage(nodeInfos []*flow.NodeInfo, userID int64, dataId int64, targetId int64,
	procInstID int64, step int) error {
	// 生成新的任务
	// TODO var task = getNewTask(nodeInfos, step, procInstID) //新任务
	// taksID, err := task.NewTaskTx(tx)
	taskQueryArr := l.svcCtx.ActStore.Task.Query().Where(task.StepEQ(step), task.ProcInstIDEQ(procInstID)).
		IDsX(l.ctx)
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
	updateRow, err := l.svcCtx.ActStore.ProcInst.Update().
		SetNodeID(string(nodeInfos[step].NodeID)).SetTaskID(int64(taskID)).SetProcDefID(procInstID).
		SetEndTime(time.Now()).SetState(flow.HAVEPASS).
		Save(l.ctx)

	if err != nil {
		return err
	}
	if updateRow == 0 {
		return err
	}
	// if step == 0 { // 流程回到起始位置，注意起始位置为0,
	// 	err = AddCandidateUserTx(userID, dataId, step, taksID, procInstID, tx)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }
	// // 添加candidate group
	// err = AddIdentityLinkCandidateTx(targetId, step, taksID, procInstID, tx)
	// if err != nil {
	// 	return err
	// }
	return nil
}
