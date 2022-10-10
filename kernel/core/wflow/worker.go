package wflow

import (
	"encoding/json"
	"errors"
	"github.com/mumushuiding/util"
	"go-wflow/common/flow"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow/general"
	"log"
	"time"
)

type StartProcInstReq struct {
	FormId      uint64
	AppId       uint64
	DefId       uint64
	RefId       uint64
	RemainHours uint64
	Title       string
	DataId      uint64
	Resource    string
	Notifiers   string
	FlowType    uint64
}

func (w *WflowProvider) StartView(in *kernel.StartViewReq) (*kernel.StartViewResp, error) {
	procdef := w.FindDefByFormId(in.FormId, in.AppId)
	var (
		node      *flow.Node
		list      []*flow.NodeInfo
		notifiers []int64
	)
	err := util.Str2Struct(procdef.Resource, &node)
	if err != nil {
		return nil, err
	}
	list, notifiers = flow.ShiftNodeTree(node, list, notifiers)
	notifiers = general.RemoveRepByLoop(notifiers)
	listByte, err := json.Marshal(list)
	notifiersByte, err := json.Marshal(notifiers)
	var flowType uint64 = 1
	if flow.CheckConCurrentFlow(node) {
		flowType = 2
	}
	return &kernel.StartViewResp{
		DataId:    in.DataId,
		DefId:     procdef.ID,
		AppId:     in.AppId,
		FormId:    in.FormId,
		Resource:  string(listByte),
		Notifiers: string(notifiersByte),
		FlowType:  flowType,
	}, nil
}

func (w *WflowProvider) StartProcess(in *kernel.StartProcInstReq) (*kernel.Nil, error) {
	var err error
	inst := w.findExistProcInst(in.DefId, in.DataId)
	if inst == nil {
		inst, err = w.CreateProcInst(StartProcInstReq{
			DefId:       in.DefId,
			DataId:      in.DataId,
			Title:       in.Title,
			RemainHours: in.RemainHours,
			Resource:    in.Resource,
		})

	} else {
		oldInstId := inst.ID
		inst, err = w.CreateProcInst(StartProcInstReq{
			DefId:       in.DefId,
			RefId:       oldInstId,
			DataId:      in.DataId,
			Title:       in.Title,
			RemainHours: in.RemainHours,
			Resource:    in.Resource,
			FlowType:    in.FlowType,
		})
	}
	//生成第一个流程节点任务
	_, err = w.CreateTask(SaveTaskReq{
		NodeID:     flow.NodeTypes[flow.ROOT],
		InstId:     inst.ID,
		IsFinished: 1,
	})
	log.Println("genNodeTree before")

	nodeInfos, _ := w.toNodeInfos(inst.ID, in.Resource)

	w.toNotifiers(inst.ID, in.Notifiers)

	log.Println("genNodeTree after")
	w.GenNormalProcess(nodeInfos[1], inst.ID, in.DataId)

	return &kernel.Nil{}, err
}

func (w *WflowProvider) GenNormalProcess(node *flow.NodeInfo, instId uint64, dataId uint64) error {
	secondTaskId, err := w.createTaskAndIdentityLink(node, instId)
	log.Println("secondTaskId", secondTaskId)
	if err != nil {
		return err
	}
	err = w.UpdateProcInst(UpdateProcInstReq{
		NodeId: node.NodeID,
		DataId: dataId,
		TaskId: secondTaskId,
		State:  flow.PENDING,
	})
	return err
}

func (w *WflowProvider) toNodeInfos(instId uint64, resource string) ([]*flow.NodeInfo, error) {
	var list []*flow.NodeInfo
	err := util.Str2Struct(resource, &list)
	if err != nil {
		return nil, errors.New("NodeInfo节点转换错误")
	}
	for _, v := range list {
		err = w.CreateNodeDetails(v, instId)
	}
	return list, nil
}
func (w *WflowProvider) toNotifiers(instId uint64, notifiers string) ([]uint64, error) {
	var userIds []uint64
	err := util.Str2Struct(notifiers, &userIds)
	if err != nil {
		return nil, errors.New("Notifiers节点转换错误")
	}
	for _, v := range userIds {
		err = w.CreateNodeNotifiers(instId, v)

	}
	return userIds, nil
}

func (w *WflowProvider) createTaskAndIdentityLink(n *flow.NodeInfo, instId uint64) (uint64, error) {
	log.Println("into createTaskAndIdentityLink", n.Type)
	taskId, _ := w.CreateTask(SaveTaskReq{
		NodeID:     n.NodeID,
		InstId:     instId,
		IsFinished: 0,
	})
	//保存审批人
	for _, v := range n.AssignedUser {
		w.CreateIdentityLink(SaveIdentityLinkReq{
			InstId:   instId,
			TaskId:   taskId,
			UserId:   uint64(v.ID),
			UserName: v.Name,
		})
	}
	return taskId, nil
}

func (w *WflowProvider) CompleteProcess(in *kernel.CompleteProcInstReq) (*kernel.Nil, error) {

	inst, err := w.FindProcInstByDataId(in.DataId)
	log.Println("aaaa", inst.TaskID)
	valid, err := w.IsApprovalValid(inst.TaskID)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("当前用户没有审批权限")
	}
	log.Println("bbbbb", inst.TaskID)
	// 1.找出该数据的流程最新审批节点
	t, err := w.FindTaskById(inst.TaskID)
	if err != nil {
		return nil, err
	}
	log.Printf("taskId %d", inst.TaskID)
	var isTaskFinished uint64 = 1
	identityLink, err := w.FindIdentityLinkByTaskId(inst.TaskID)
	//2、根据最新节点ID和用户ID找到对应的审批人
	if err != nil {
		return nil, err
	}
	log.Println(3333333)
	err = w.UpdateIdentityLink(UpdateIdentityLinkReq{
		UserId:  identityLink.UserID,
		TaskId:  inst.TaskID,
		Comment: in.Comment,
		Result:  in.Result,
		IsDeal:  1,
	})

	var isInstFinish uint64
	var approvalState uint64
	newTaskId := inst.TaskID
	nodeInfo, err := w.FindNodeDetail(inst.ID, t.NodeId)
	//nodeInfo, err := convertNodeDetailToNodeInfo(detail)
	//nodeInfo, err := l.findNodeDetail(inst.ID, t.NodeId)
	if err != nil {
		return nil, err
	}
	currentNodeId := nodeInfo.NodeID
	log.Printf("nodeInfo:%v \n", nodeInfo)
	log.Printf("currentNodeId:%s \n", currentNodeId)
	switch in.Result {
	case flow.HAVEPASS:
		//如果审批通过，我们首先判断其是会签还是或签。会签的话，需要判断是否所有人审批完。
		//只有当会签所有人审批完和或签有一个人审批，就完成了该流程节点的审批。
		log.Println(5555555)
		//if err != nil {
		//	return nil, err
		//}
		nextNodeInfo := &flow.NodeInfo{
			NodeID: flow.NodeTypes[flow.END],
			Mode:   flow.MODE_OR,
		}
		if nodeInfo.NextId != flow.NodeTypes[flow.END] {
			nextNodeInfo, err = w.FindNodeDetail(inst.ID, nodeInfo.NextId)
			//nextNodeInfo, err = convertNodeDetailToNodeInfo(detail)
			//nextNodeInfo, err = l.findNodeDetail(inst.ID, nodeInfo.NextId)
			if err != nil {
				return nil, err
			}
		}
		log.Println(66666, nodeInfo.NextId)
		if nodeInfo.Mode == flow.MODE_OR {
			//isInstFinish = flow.IsEndNode(nodeInfo)
			//log.Printf("flow.MODE_OR:%s,isInstFinish:%d \n", flow.MODE_OR, isInstFinish)
			if nodeInfo.NextId == flow.NodeTypes[flow.END] {
				isInstFinish = 1
				approvalState = flow.HAVEPASS
				newTaskId, err = w.CreateTask(SaveTaskReq{
					NodeID:     flow.NodeTypes[flow.END],
					InstId:     inst.ID,
					IsFinished: isInstFinish,
				})

			} else {
				log.Println("flow.DEALING", flow.DEALING)
				approvalState = flow.DEALING
				newTaskId, err = w.moveNextStage(inst.ID, nextNodeInfo)
			}
		} else if nodeInfo.Mode == flow.MODE_AND {

			//1.判断是否该节点的所有人都已经审批
			//2.是否为末级节点
			isTaskCompleted, err := w.IsAllApprovalComplete(inst.TaskID, identityLink.UserID)
			log.Println("flow.MODE_AND", flow.MODE_AND, isTaskCompleted)
			if err != nil {
				return nil, err
			}
			if isTaskCompleted {
				log.Println("taskCompleted:", nodeInfo.NextId)
				if nodeInfo.NextId == flow.NodeTypes[flow.END] {
					log.Println("process end")
					isInstFinish = 1
					approvalState = flow.HAVEPASS
					currentNodeId = flow.NodeTypes[flow.END]
					newTaskId, err = w.CreateTask(SaveTaskReq{
						NodeID:     currentNodeId,
						InstId:     inst.ID,
						IsFinished: 1,
					})

				} else {
					log.Println("andmoveNextStage")
					approvalState = flow.DEALING
					newTaskId, err = w.moveNextStage(inst.ID, nextNodeInfo)
					if err != nil || newTaskId == 0 {
						return nil, err
					}
				}

			} else {
				currentNodeId = nodeInfo.NodeID
				log.Println("andDEALING")
				approvalState = flow.DEALING
				isTaskFinished = 0
			}

		}
		if isInstFinish == 1 {
			currentNodeId = flow.NodeTypes[flow.END]
			approvalState = flow.HAVEPASS

			err = w.SetNodeNotifierPermit(inst.ID)
		}
		//6、审批不通过 直接结束流程 更新流程实例 is_finished=1 state = 7
	case flow.NOTPASS:
		log.Println("req.Result", in.Result)
		isInstFinish = 1
		approvalState = flow.NOTPASS
		//7.驳回  驳回到上一节点审批 驳回到指定节点审批 直接结束流程
	case flow.REJECT:
		approvalState = flow.REJECT
		log.Println("nodeInfo.Refuse:", nodeInfo.Refuse)
		if nodeInfo.Refuse.Type == flow.TO_BEFORE {
			log.Println("TO_BEFORE")
			//驳回至发起人
			if nodeInfo.PrevId == flow.NodeTypes[flow.ROOT] {
				log.Println("rejectToRoot")
				err = w.rejectToRoot(inst.ID, in.DataId, in.Comment, inst.TaskID)
				//删除本次抄送人的数据，避免多次抄送
				err = w.DeleteNodeNotifiersByID(inst.ID)
				err = w.DeleteNodeDetailsByID(inst.ID)
				return &kernel.Nil{}, err
			}
			if err != nil {
				return nil, err
			}
			//nodeInfo, err = l.findNodeDetail(inst.ID, nodeInfo.PrevId)
			nodeInfo, err = w.FindNodeDetail(inst.ID, nodeInfo.PrevId)
			//nodeInfo, err = convertNodeDetailToNodeInfo(detail)
			newTaskId, err = w.moveNextStage(inst.ID, nodeInfo)
			//if err != nil || newTaskId == 0 {
			//	tx.Rollback()
			//	return nil, err
			//}
		} else if nodeInfo.Refuse.Type == flow.TO_NODE {
			if nodeInfo.Refuse.Target == flow.NodeTypes[flow.ROOT] {
				err = w.rejectToRoot(inst.ID, in.DataId, in.Comment, inst.TaskID)
				//删除本次抄送人的数据，避免多次抄送
				err = w.DeleteNodeNotifiersByID(inst.ID)
				err = w.DeleteNodeDetailsByID(inst.ID)

				return &kernel.Nil{}, err
			}
			log.Println("TO_NODE")
			targetNodeInfo, _ := w.FindNodeDetail(inst.ID, nodeInfo.Refuse.Target)
			newTaskId, err = w.moveNextStage(inst.ID, targetNodeInfo)
			//if err != nil || newTaskId == 0 {
			//	tx.Rollback()
			//	return nil, err
			//}
		} else if nodeInfo.Refuse.Type == flow.TO_END {
			log.Println("TO_END")
			approvalState = flow.DISCARD
			isInstFinish = 1
		}
	}
	err = w.UpdateTask(UpdateTaskReq{
		TaskId:     inst.TaskID,
		IsFinished: isTaskFinished,
	})
	var flowCode string
	if approvalState == flow.HAVEPASS {
		flowCode = w.generateCode()
	}

	//8.反向更新流程实例表
	err = w.UpdateProcInst(UpdateProcInstReq{
		TaskId:     newTaskId,
		DataId:     in.DataId,
		NodeId:     currentNodeId,
		State:      approvalState,
		IsFinished: isInstFinish,
		Code:       flowCode,
	})

	return &kernel.Nil{}, nil
}

func (w *WflowProvider) rejectToRoot(instId uint64, dataId uint64, comment string, oldTaskId uint64) error {
	taskId, err := w.CreateTask(SaveTaskReq{
		NodeID:     flow.NodeTypes[flow.ROOT],
		InstId:     instId,
		IsFinished: 1,
	})

	err = w.UpdateIdentityLink(UpdateIdentityLinkReq{
		UserId:  general.MyUserId,
		TaskId:  oldTaskId,
		Comment: comment,
		Result:  flow.REJECT,
		IsDeal:  1,
	})

	err = w.UpdateProcInst(UpdateProcInstReq{
		TaskId: taskId,
		DataId: dataId,
		NodeId: flow.NodeTypes[flow.ROOT],
		State:  flow.REJECTTOROOT,
	})
	return err
}

func (w *WflowProvider) generateCode() string {
	return "ORGINONE" + time.Now().Format("20060102")
}

//审批的时候，同时把审批后的抄送节点生成
func (w *WflowProvider) moveNextStage(instId uint64, nodeInfo *flow.NodeInfo) (uint64, error) {
	if nodeInfo == nil {
		taskId, err := w.CreateTask(SaveTaskReq{
			NodeID:     flow.NodeTypes[flow.END],
			InstId:     instId,
			IsFinished: 1,
		})
		if err != nil {
			return 0, err
		}
		return taskId, nil
	}
	////判断节点是否为转发节点，如果是转发节点的话，就需要生成节点任务、抄送人数据
	//4.保存流程任务
	taskId, err := w.CreateTask(SaveTaskReq{
		NodeID:     nodeInfo.NodeID,
		InstId:     instId,
		IsFinished: 0,
	})
	log.Println("moveNextStage", taskId)
	if err != nil {
		return 0, err
	}
	log.Println("moveNextStage ing")
	//保存审批人
	for _, v := range nodeInfo.AssignedUser {
		_, err = w.CreateIdentityLink(SaveIdentityLinkReq{
			InstId:   instId,
			TaskId:   taskId,
			UserId:   uint64(v.ID),
			UserName: v.Name,
		})
		if err != nil {
			return 0, err
		}
	}
	return taskId, nil
}
