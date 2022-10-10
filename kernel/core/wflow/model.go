package wflow

import (
	"context"
	"go-wflow/common/flow"
	"go-wflow/kernel"

	"go-wflow/kernel/ent"
	schema "go-wflow/kernel/ent"
)

type (
	WflowInterface interface {
		CreateProcDef(in *kernel.SaveProcDefReq) (*kernel.Procdef, error)
		CreateProcInst(in StartProcInstReq) (*ent.Procinst, error)
		CreateTask(in SaveTaskReq) (uint64, error)
		CreateIdentityLink(in SaveIdentityLinkReq) (uint64, error)
		CreateNodeDetails(in *flow.NodeInfo, instId uint64) error
		CreateNodeNotifiers(instId uint64, userId uint64) error

		DeleteProcDefByID(in *kernel.IdReq) (*kernel.Nil, error)
		DeleteProcInstByID(id uint64) error
		DeleteTaskByID(id uint64) error
		DeleteIdentityLinkByID(id uint64) error
		DeleteNodeDetailsByID(id uint64) error
		DeleteNodeNotifiersByID(id uint64) error

		UpdateProcDef(in *kernel.SaveProcDefReq) error
		UpdateIdentityLink(in UpdateIdentityLinkReq) error
		UpdateNodeNotifierByInstId(instId uint64) error
		UpdateTask(in UpdateTaskReq) error
		UpdateProcInst(in UpdateProcInstReq) error
		SetIsActive(in *kernel.QueryProcDefReq) (*kernel.Nil, error)
		SetNodeNotifierPermit(instId uint64) error

		StartProcess(in *kernel.StartProcInstReq) (*kernel.Nil, error)
		CompleteProcess(in *kernel.CompleteProcInstReq) (*kernel.Nil, error)
		StartView(in *kernel.StartViewReq) (*kernel.StartViewResp, error)

		GetNewVersion(formId uint64, appId uint64) (uint64, error)
		FindDefByFormId(formId uint64, appId uint64) *ent.Procdef
		FindDefByVersion(in *kernel.QueryProcDefReq) *ent.Procdef
		FindProcInstByDataId(dataId uint64) (*ent.Procinst, error)
		IsAllApprovalComplete(taskId uint64, userId uint64) (bool, error)
		IsApprovalValid(taskId uint64) (bool, error)
		FindTaskById(taskId uint64) (*ent.Task, error)
		FindIdentityLinkByTaskId(taskId uint64) (*ent.Identitylink, error)
		FindNodeDetail(instId uint64, nodeId string) (*flow.NodeInfo, error)
	}

	WflowProvider struct {
		PersonId uint64
		Client   *schema.Client
		SrcCtx   context.Context
	}
)

// TODO 初始化一次
func New(personId uint64, client *ent.Client, ctx context.Context) WflowInterface {
	return &WflowProvider{
		PersonId: personId,
		Client:   client,
		SrcCtx:   ctx,
		// Client:   schema.FromContext(ctx),
	}
}

func (w *WflowProvider) ctx() {

}
