package logic

import (
	"context"
	"errors"
	"log"
	"time"

	"act/common/act/procinst"
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
	log.Println("in 2222222", in.DataId)
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 2.2根据procdefID在proc_def数据库表中查询到create_user_id
	procinstInfo, err := tx.ProcInst.Query().Where(procinst.DataIDEQ(in.DataId)).First(l.ctx)
	if err != nil {
		return nil, err
	}

	// TODO StartUserID=11025 StartUserName=xiaoming
	var userid int64 = 11025
	if userid == procinstInfo.StartUserID {
		_, err := tx.ProcInst.Update().
			Where(procinst.ProcDefID(procinstInfo.ProcDefID)).
			SetState(4).SetIsFinished(1).SetEndTime(time.Now()).SetUpdateTime(time.Now()).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		return &act.Nil{}, errors.New("人员未找到！")
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	// 3.UserID和create_user_id比较不相等返回，无权限撤回

	// 4.UserID和create_user_id比较相等,将流程实例表state=4,isFinish=1,endTime=now,updateTime=now

	return &act.Nil{}, nil
}
