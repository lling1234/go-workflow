package act

import (
	"context"
	"log"

	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/actclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
1、需要判断撤回人员是否为发起人。（只有发起人才能撤回）
2、撤回将流程实例表state=4,isFinish=1,endTime=now,updateTime=now
*/
func (l *WithdrawLogic) Withdraw(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
	log.Println("req 1111111", req.DataId)
	// 1.通过token中解析用户id
	// userID := 101
	// 2.1根据dataID在proc_inst数据库表中查询procdefID
	_, err = l.svcCtx.Rpc.Withdraw(l.ctx, &actclient.DataIdReq{DataId: req.DataId})
	log.Println("err 3333333", err)
	if err != nil {
		return types.GetErrorCommonResponse("撤回失败！")
	}
	// 2.2根据procdefID在proc_def数据库表中查询到create_user_id

	// 3.UserID和create_user_id比较不相等返回，无权限撤回

	// 4.UserID和create_user_id比较相等,将流程实例表state=4,isFinish=1,endTime=now,updateTime=now

	return types.GetSuccessCommonResponse("撤回成功！")
}
