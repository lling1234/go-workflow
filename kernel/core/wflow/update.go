package wflow

import (
	"go-wflow/common/flow"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow/general"
	"go-wflow/kernel/ent/identitylink"
	"go-wflow/kernel/ent/nodenotifiers"
	"go-wflow/kernel/ent/procdef"
	"go-wflow/kernel/ent/procinst"
	"go-wflow/kernel/ent/task"
	"time"
)

type (
	UpdateIdentityLinkReq struct {
		UserId  uint64
		TaskId  uint64
		Comment string
		Result  uint64
		IsDeal  uint64
	}

	UpdateTaskReq struct {
		TaskId     uint64
		IsFinished uint64
	}

	UpdateProcInstReq struct {
		TaskId     uint64
		DataId     uint64
		NodeId     string
		State      uint64
		IsFinished uint64
		Code       string
		Resource   string
	}
)

func (w *WflowProvider) UpdateProcDef(in *kernel.SaveProcDefReq) error {
	err := w.Client.Procdef.Update(w.PersonId).
		Where(procdef.FormIDEQ(in.FormId), procdef.AppIDEQ(in.AppId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(in.Version)).
		SetName(in.Name).SetCode(in.Code).SetFormID(in.FormId).SetFormName(in.FormName).
		SetAppName(in.AppName).SetAppID(in.AppId).
		SetRemainHours(in.RemainHours).SetResource(in.Resource).
		SetVersion(in.Version).
		// SetCreateTime(time.Now()).SetVersion(uint64(in.Version)).
		Exec(w.SrcCtx)
	if err != nil {
		return err
	}
	return nil
}

func (w *WflowProvider) UpdateIdentityLink(in UpdateIdentityLinkReq) error {
	w.Client.Identitylink.Update(w.PersonId).Where(identitylink.TaskIDEQ(in.TaskId), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDealEQ(0)).
		SetUpdateTime(time.Now()).SetComment(in.Comment).SetIsDeal(1).SetResult(uint64(in.Result)).
		ExecX(w.SrcCtx)
	return nil
}

func (w *WflowProvider) UpdateNodeNotifierByInstId(instId uint64) error {
	w.Client.Nodenotifiers.Update(w.PersonId).Where(nodenotifiers.ProcInstIDEQ(instId)).SetIsPermit(1).ExecX(w.SrcCtx)
	return nil
}

func (w *WflowProvider) UpdateTask(in UpdateTaskReq) error {
	w.Client.Task.Update(w.PersonId).Where(task.IDEQ(in.TaskId)).SetUpdateTime(time.Now()).SetIsFinished(uint64(in.IsFinished)).SetClaimTime(time.Now()).
		ExecX(w.SrcCtx)
	return nil
}

func (w *WflowProvider) UpdateProcInst(in UpdateProcInstReq) error {
	procInstUpdate := w.Client.Procinst.Update(w.PersonId).Where(procinst.DataIDEQ(in.DataId), procinst.StateNotIn(flow.WITHDRAW, flow.DISCARD))

	if in.TaskId != 0 {
		procInstUpdate.SetNodeID(in.NodeId).SetTaskID(in.TaskId)
	}
	if in.State != 0 {
		procInstUpdate.SetState(in.State)
	}
	if in.State == flow.HAVEPASS {
		procInstUpdate.SetCode(in.Code)
	}
	if in.IsFinished == 1 {
		procInstUpdate.SetIsFinished(1).SetFinishTime(time.Now())
	}
	procInstUpdate.SetResource(in.Resource).SetUpdateTime(time.Now()).SetUpdateUser(general.MyUserId).
		ExecX(w.SrcCtx)

	return nil
}
// TODO 错误处理和error
func (w *WflowProvider) SetIsActive(in *kernel.QueryProcDefReq) (*kernel.Nil, error) {
	w.Client.Procdef.Update(w.PersonId).Where(procdef.FormIDEQ(in.FormId), procdef.AppIDEQ(in.AppId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(in.Version)).
		SetIsActive(1).SetUpdateTime(time.Now()).SetVersion(in.Version).Exec(w.SrcCtx)
	w.Client.Procdef.Update(w.PersonId).Where(procdef.FormIDEQ(in.FormId), procdef.AppIDEQ(in.AppId), procdef.TargetIDEQ(general.TargetId), procdef.VersionNEQ(in.Version)).
		SetIsActive(0).SetVersion(in.Version).Exec(w.SrcCtx)
	return nil, nil
}

func (w *WflowProvider) SetNodeNotifierPermit(instId uint64) error {
	err := w.Client.Nodenotifiers.Update(w.PersonId).Where(nodenotifiers.ProcInstID(instId)).SetIsPermit(1).Exec(w.SrcCtx)
	return err
}
