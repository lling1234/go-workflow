package wflow

import (
	"encoding/json"
	"go-wflow/common/flow"
	"go-wflow/kernel"
	"go-wflow/kernel/ent"
	"time"

	"go-wflow/kernel/core/wflow/general"
	"log"
)

type (
	SaveTaskReq struct {
		NodeID     string
		InstId     uint64
		IsFinished uint64
	}

	SaveIdentityLinkReq struct {
		InstId   uint64
		TaskId   uint64
		UserId   uint64
		UserName string
		Result   uint64
	}
)

func (w *WflowProvider) CreateProcDef(in *kernel.SaveProcDefReq) (*kernel.Procdef, error) {
	version, err := w.GetNewVersion(in.FormId, in.AppId)
	if err != nil {
		return nil, err
	}
	log.Println("CreateProcessDef 111111", in)

	procdefData := w.Client.Procdef.Create(w.PersonId).
		SetName(in.Name).SetCode(in.Code).SetFormID(in.FormId).SetFormName(in.FormName).
		SetAppName(in.AppName).SetAppID(in.AppId).
		SetRemainHours(in.RemainHours).SetResource(in.Resource).
		SetCreateUser(general.MyUserId).SetCreateUserName(general.UserName0).
		SetCreateTime(time.Now()).SetVersion(version).
		SetTargetID(general.TargetId).
		SaveX(w.SrcCtx)

	log.Println("procdefData.CreateTime", procdefData.CreateTime.Format("2006-01-02 15:04:05:06"))

	log.Println("procDef", procdefData.ToMessage().String())

	// TODO  返回值,时间带有时区
	return procdefData.ToMessage(), nil
}

func (w *WflowProvider) CreateProcInst(in StartProcInstReq) (*ent.Procinst, error) {
	procInst := w.Client.Procinst.Create(w.PersonId).
		SetProcDefID(in.DefId).SetTitle(in.Title).
		SetDataID(in.DataId). //.SetFlowType(uint64(in.FlowType))
		SetState(flow.PENDING).
		SetRemainHours(in.RemainHours).SetResource(in.Resource).
		SetTargetID(general.TargetId).
		SaveX(w.SrcCtx)
	log.Println("procinst", procInst)
	// TODO  返回值
	return procInst, nil
}

func (w *WflowProvider) CreateTask(in SaveTaskReq) (uint64, error) {
	task := w.Client.Task.Create(w.PersonId).SetNodeId(in.NodeID).SetProcInstID(in.InstId).
		SetIsFinished(uint64(in.IsFinished)).SetCreateUser(general.MyUserId).
		SaveX(w.SrcCtx)
	log.Println("task", task)
	return task.ID, nil
}

func (w *WflowProvider) CreateIdentityLink(in SaveIdentityLinkReq) (uint64, error) {
	identityLinkCreate := w.Client.Identitylink.Create(w.PersonId).
		SetProcInstID(in.InstId).SetTaskID(in.TaskId).SetTargetID(general.TargetId).
		SetIsDeal(0).SetUserID(in.UserId).SetUserName(in.UserName).SetCreateUser(general.MyUserId)
	if in.Result != 0 {
		identityLinkCreate.SetResult(uint64(in.Result))
	}
	identityLink := identityLinkCreate.SaveX(w.SrcCtx)
	return identityLink.ID, nil
}

func (w *WflowProvider) CreateNodeDetails(n *flow.NodeInfo, instId uint64) error {
	nodeInfo := ""
	if n.AssignedUser != nil && len(n.AssignedUser) > 0 {
		assignedUser, err := json.Marshal(n.AssignedUser)
		if err != nil {
			return err
		}
		nodeInfo = string(assignedUser)
	}
	refuseStr := ""
	if n.Refuse != nil && n.Refuse.Type != "" {
		refuseData, err := json.Marshal(n.Refuse)
		if err != nil {
			return err
		}
		refuseStr = string(refuseData)
	}
	log.Println("SaveNodeDetails ing", n.Mode)
	w.Client.NodeDetails.Create(w.PersonId).SetNodeID(n.NodeID).SetProcInstID(instId).SetPrevID(n.PrevId).SetNextID(n.NextId).
		SetNodeInfo(nodeInfo).SetMode(n.Mode).SetRefuse(refuseStr).SaveX(w.SrcCtx)
	return nil
}

func (w *WflowProvider) CreateNodeNotifiers(instId uint64, userId uint64) error {
	w.Client.Nodenotifiers.Create(w.PersonId).SetProcInstID(instId).SetUserID(userId).SetTargetID(general.TargetId).SetCreateUser(general.MyUserId).
		SaveX(w.SrcCtx)
	return nil
}
