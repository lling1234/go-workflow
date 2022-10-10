package wflow

import (
	"errors"
	"go-wflow/common/flow"
	"go-wflow/common/utils/linq"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow/general"
	"go-wflow/kernel/ent"
	"go-wflow/kernel/ent/identitylink"
	"go-wflow/kernel/ent/nodedetails"
	"go-wflow/kernel/ent/procdef"
	"go-wflow/kernel/ent/procinst"
	"go-wflow/kernel/ent/task"
	"log"

	"github.com/mumushuiding/util"
)

func (w *WflowProvider) GetNewVersion(formId uint64, appId uint64) (uint64, error) {
	defs := w.Client.Procdef.Query().
		Where(procdef.FormIDEQ(formId), procdef.AppIDEQ(appId)).Select(procdef.FieldVersion).
		AllX(w.SrcCtx)
	maxVersion := linq.From(defs).SelectT(func(a *ent.Procdef) uint64 {
		return uint64(a.Version)
	}).Max()
	if maxVersion == nil {
		return 1, nil
	}
	return maxVersion.(uint64) + 1, nil
}

func (w *WflowProvider) findExistProcInst(defId uint64, dataId uint64) *ent.Procinst {
	all, err := w.Client.Procinst.Query().Where(procinst.ProcDefIDEQ(defId), procinst.DataIDEQ(dataId), procinst.IsFinished(0)).All(w.SrcCtx)
	if err != nil || all == nil || len(all) == 0 {
		return nil
	}

	return all[0]
}

func (w *WflowProvider) FindDefByFormId(formId uint64, appId uint64) *ent.Procdef {
	defs := w.Client.Procdef.Query().
		Where(procdef.FormIDEQ(formId), procdef.AppIDEQ(appId), procdef.IsActiveEQ(1)).
		AllX(w.SrcCtx)
	if defs != nil && len(defs) > 0 {
		return defs[0]
	}
	return nil
}

func (w *WflowProvider) FindDefByVersion(in *kernel.QueryProcDefReq) *ent.Procdef {
	defs := w.Client.Procdef.Query().
		Where(procdef.FormIDEQ(in.FormId), procdef.AppIDEQ(in.AppId), procdef.Version(in.Version)).
		AllX(w.SrcCtx)
	log.Println("FindDefByVersion", in.FormId)
	if defs != nil && len(defs) > 0 {
		return defs[0]
	}
	return nil
}

func (w *WflowProvider) FindProcInstByDataId(dataId uint64) (*ent.Procinst, error) {
	insts, err := w.Client.Procinst.Query().Where(procinst.DataIDEQ(dataId)).All(w.SrcCtx)
	log.Printf("inst.len:%d\n")
	if err != nil {
		return nil, err
	}
	if insts == nil || len(insts) == 0 {
		return nil, errors.New("该数据没有对应的流程实例")
	}
	inst := insts[0]
	return inst, nil
}

func (w *WflowProvider) IsAllApprovalComplete(taskId uint64, userId uint64) (bool, error) {
	all, err := w.Client.Identitylink.Query().
		Where(identitylink.TaskIDEQ(taskId), identitylink.IsDeal(0), identitylink.UserIDNEQ(userId)).All(w.SrcCtx)
	if err != nil {
		return false, err
	}
	if all == nil || len(all) == 0 {
		return true, nil
	}
	return false, nil
}

//判断当前用户是否有审批权限
func (w *WflowProvider) IsApprovalValid(taskId uint64) (bool, error) {
	all, err := w.Client.Identitylink.Query().
		Where(identitylink.TaskIDEQ(taskId), identitylink.IsDeal(0), identitylink.UserIDEQ(general.MyUserId)).All(w.SrcCtx)
	if err != nil {
		return false, err
	}
	if all == nil || len(all) == 0 {
		return false, nil
	}
	return true, nil
}

func (w *WflowProvider) FindTaskById(taskId uint64) (*ent.Task, error) {
	t, err := w.Client.Task.Query().Where(task.IDEQ(taskId)).First(w.SrcCtx)
	return t, err
}

func (w *WflowProvider) FindIdentityLinkByTaskId(taskId uint64) (*ent.Identitylink, error) {
	ils, err := w.Client.Identitylink.Query().
		Where(identitylink.TaskIDEQ(taskId), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDealEQ(0)).
		Select(identitylink.FieldID, identitylink.FieldUserID, identitylink.FieldUserName, identitylink.FieldProcInstID, identitylink.FieldTaskID, identitylink.FieldTargetID).
		All(w.SrcCtx)
	if err != nil {
		return nil, err
	}
	if ils == nil || len(ils) == 0 {
		return nil, errors.New("用户没有审批权限")
	}
	return ils[0], nil
}

func (w *WflowProvider) FindNodeDetail(instId uint64, nodeId string) (*flow.NodeInfo, error) {
	datas, err := w.Client.NodeDetails.Query().Where(nodedetails.ProcInstIDEQ(instId), nodedetails.NodeIDEQ(nodeId)).All(w.SrcCtx)
	if err != nil {
		return nil, err
	}
	if datas == nil || len(datas) == 0 {
		return nil, errors.New("没有找到对应的节点对象")
	}
	return convertNodeDetailToNodeInfo(datas[0])
}

func convertNodeDetailToNodeInfo(data *ent.NodeDetails) (*flow.NodeInfo, error) {
	var err error
	var assignedUser []*flow.AssignedUser
	if data.NodeInfo != "" {
		err = util.Str2Struct(data.NodeInfo, &assignedUser)
		log.Println(" Str2Struct ing")
		if err != nil {
			return nil, err
		}
	}
	var refuse *flow.Refuse
	log.Println("data.Refuse", data.Refuse)
	if data.Refuse != "" {
		err = util.Str2Struct(data.Refuse, &refuse)
		if err != nil {
			return nil, err
		}
	}
	log.Println(" findNodeDetail end")
	return &flow.NodeInfo{
		NodeID:       data.NodeID,
		NextId:       data.NextID,
		PrevId:       data.PrevID,
		Mode:         data.Mode,
		AssignedUser: assignedUser,
		Refuse:       refuse,
	}, nil
}
