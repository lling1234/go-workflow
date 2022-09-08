package logic

import (
	"act/rpc/constant"
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
	// TODO userid=101
	var userid int64 = 101
	//1、"待处理", 2、"处理中", 3、 "驳回", 4、"已撤回" ,5、 "未通过",6、 "已通过", 7、"废弃"
	//流程状态为 1、"待处理", 2、"处理中", 3、 "驳回" 才可以被撤回
	if userid == procinstInfo.StartUserID && procinstInfo.State < constant.WITHDRAW {
		_, err := tx.ProcInst.Update().
			Where(procinst.ProcDefID(procinstInfo.ProcDefID)).
			SetState(constant.WITHDRAW).SetIsFinished(1).SetEndTime(time.Now()).SetUpdateTime(time.Now()).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		return &act.Nil{}, errors.New("人员未找到！")
	}
	err = tx.Commit()
<<<<<<< HEAD
	if err != nil {
		return nil, err
	}
=======
>>>>>>> 4a54178f732840048a6d221f9338a6e8f5d12ba3
	// 3.UserID和create_user_id比较不相等返回，无权限撤回

	// 4.UserID和create_user_id比较相等,将流程实例表state=4,isFinish=1,endTime=now,updateTime=now

	return &act.Nil{}, err
}
