package logic

import (
	act2 "act/common/act"
	"act/common/act/identitylink"
	"act/common/act/procinst"
	"act/common/act/task"
	"act/common/flow"
	"act/rpc/general"
	"context"
	"strconv"
	"time"
)

type (
	UpdateIdentityLinkReq struct {
		UserId  int64
		TaskId  int64
		Comment string
		Result  int32
		IsDeal  int32
	}

	UpdateTaskReq struct {
		TaskId     int64
		IsFinished int32
	}

	UpdateProcInstReq struct {
		TaskId     int64
		DataId     int64
		NodeId     string
		State      int32
		IsFinished int32
		Code       string
	}
)

func UpdateIdentityLink(in UpdateIdentityLinkReq, tx *act2.Tx, ctx context.Context) error {
	err := tx.IdentityLink.Update().Where(identitylink.TaskIDEQ(in.TaskId), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDelEQ(0), identitylink.IsDealEQ(0)).
		SetUpdateTime(time.Now()).SetComment(in.Comment).SetResult(in.Result).SetIsDeal(1).Exec(ctx)

	return err
}

func UpdateTask(in UpdateTaskReq, tx *act2.Tx, ctx context.Context) error {
	oldTask, err := tx.Task.Query().Where(task.IDEQ(in.TaskId)).First(ctx)
	if err != nil {
		return err
	}
	agreers := ""
	if oldTask.AgreeApprover == "" {
		agreers = strconv.FormatInt(general.MyUserId, 10)
	} else {
		agreers = oldTask.AgreeApprover + "," + strconv.FormatInt(general.MyUserId, 10)
	}
	err = tx.Task.Update().Where(task.IDEQ(in.TaskId)).SetUpdateTime(time.Now()).SetIsFinished(in.IsFinished).SetAgreeApprover(agreers).SetClaimTime(time.Now()).Exec(ctx)

	return err
}

func UpdateProcInst(in UpdateProcInstReq, tx *act2.Tx, ctx context.Context) error {
	procInstUpdate := tx.ProcInst.Update().Where(procinst.DataIDEQ(in.DataId), procinst.StateNotIn(flow.WITHDRAW, flow.DISCARD), procinst.IsDelEQ(0))

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
		procInstUpdate.SetIsFinished(1).SetEndTime(time.Now())
	}
	err := procInstUpdate.SetUpdateTime(time.Now()).SetUpdateUserID(general.MyUserId).Exec(ctx)

	return err
}
