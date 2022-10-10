package wflow

import (
	"go-wflow/kernel"
	"go-wflow/kernel/ent/identitylink"
	"go-wflow/kernel/ent/nodedetails"
	"go-wflow/kernel/ent/nodenotifiers"
	"go-wflow/kernel/ent/procdef"
	"go-wflow/kernel/ent/procinst"
	"go-wflow/kernel/ent/task"
)

// TODO 软删除

func (w *WflowProvider) DeleteProcDefByID(in *kernel.IdReq) (*kernel.Nil, error) {
	_, err := w.Client.Procdef.Delete(in.Id).
		Where(procdef.ID(in.Id)).
		Exec(w.SrcCtx)

	return &kernel.Nil{}, err
}

func (w *WflowProvider) DeleteProcInstByID(id uint64) error {
	_, err := w.Client.Procinst.Delete(id).
		Where(procinst.IDEQ(id)).
		Exec(w.SrcCtx)
	if err != nil {
		return err
	}
	return nil
}

func (w *WflowProvider) DeleteTaskByID(id uint64) error {
	_, err := w.Client.Task.Delete(id).
		Where(task.IDEQ(id)).
		Exec(w.SrcCtx)
	if err != nil {
		return err
	}
	return nil
}

func (w *WflowProvider) DeleteIdentityLinkByID(id uint64) error {
	_, err := w.Client.Identitylink.Delete(id).
		Where(identitylink.IDEQ(id)).
		Exec(w.SrcCtx)
	if err != nil {
		return err
	}
	return nil
}

func (w *WflowProvider) DeleteNodeDetailsByID(id uint64) error {
	_, err := w.Client.NodeDetails.Delete(id).
		Where(nodedetails.IDEQ(id)).
		Exec(w.SrcCtx)
	if err != nil {
		return err
	}
	return nil
}

func (w *WflowProvider) DeleteNodeNotifiersByID(id uint64) error {
	_, err := w.Client.Nodenotifiers.Delete(id).
		Where(nodenotifiers.IDEQ(id)).
		Exec(w.SrcCtx)
	if err != nil {
		return err
	}
	return nil
}
