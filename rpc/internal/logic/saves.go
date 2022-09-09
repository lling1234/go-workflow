package logic

import (
	act2 "act/common/act"
	"act/common/flow"
	"act/rpc/general"
	"context"
	"time"
)

type (
	SaveTaskReq struct {
		NodeID         string
		InstId         int64
		IsFinished     int32
		Level          int32
		Step           int32
		MemberApprover string
		Mode           string
	}
	SaveIdentityLinkReq struct {
		InstId   int64
		TaskId   int64
		Step     int32
		UserId   int64
		UserName string
		Mode     string
	}

	SaveExecutionReq struct {
		InstId    int64
		NodeInfos string
	}
	SaveProcInstReq struct {
		DefId       int64
		DataId      int64
		Title       string
		RemainHours int32
	}
)

func SaveTask(in SaveTaskReq, tx *act2.Tx, ctx context.Context) (int64, error) {
	taskCreate := tx.Task.Create().SetNodeID(in.NodeID).SetProcInstID(in.InstId).
		SetIsFinished(in.IsFinished).SetLevel(in.Level).SetStep(in.Step).SetMemberApprover(in.MemberApprover)
	if in.Mode != "" {
		taskCreate.SetMode(in.Mode)
	}
	t, err := taskCreate.Save(ctx)
	return t.ID, err
}

func SaveIdentityLink(in SaveIdentityLinkReq, tx *act2.Tx, ctx context.Context) error {
	_, err := tx.IdentityLink.Create().SetProcInstID(in.InstId).SetTaskID(in.TaskId).SetTargetID(general.TargetId).
		SetIsDeal(0).SetStep(in.Step).SetUserID(in.UserId).SetUserName(in.UserName).Save(ctx)

	return err
}

func SaveExecution(in SaveExecutionReq, tx *act2.Tx, ctx context.Context) (*act2.Execution, error) {
	exec, err := tx.Execution.Create().SetProcInstID(in.InstId).SetNodeInfos(in.NodeInfos).SetStartTime(time.Now()).
		Save(ctx)
	return exec, err
}

func SaveProcInst(in SaveProcInstReq, tx *act2.Tx, ctx context.Context) (*act2.ProcInst, error) {
	inst, err := tx.ProcInst.Create().SetProcDefID(in.DefId).SetDataID(in.DataId).SetTargetID(general.TargetId).SetStartTime(time.Now()).
		SetTitle(in.Title).SetIsFinished(0).SetRemainHours(in.RemainHours).SetStartUserID(general.UserId3).SetStartUserName(general.UserName3).
		SetState(flow.PENDING).Save(ctx)
	return inst, err
}
