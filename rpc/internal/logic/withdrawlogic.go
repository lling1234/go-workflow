package logic

import (
	"act/common/act/procdef"
	"act/common/act/procinst"
	"act/rpc/general"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WithdrawLogic) Withdraw(in *act.DataIdReq) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 2.2根据procdefID在proc_def数据库表中查询到create_user_id
	procinst, err := tx.ProcInst.Query().Where(procinst.DataIDEQ(in.DataId)).First(l.ctx)
	if err != nil {
		return nil, err
	}

	procdef, err := tx.ProcDef.Query().Where(procdef.IDEQ(procinst.ProcDefID)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	userid := general.UserId0
	if userid == procdef.CreateUserID {
		_, err := tx.ProcInst.Update().
			SetState(4).SetIsFinished(1).SetEndTime(time.Now()).SetUpdateTime(time.Now()).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		return &act.Nil{}, nil
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 3.UserID和create_user_id比较不相等返回，无权限撤回

	// 4.UserID和create_user_id比较相等,将流程实例表state=4,isFinish=1,endTime=now,updateTime=now

	return &act.Nil{}, nil
}
