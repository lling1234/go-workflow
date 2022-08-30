package logic

import (
	"context"
	"log"
	"time"

	"act/common/act/procdef"
	"act/common/cache"
	"act/rpc/internal/svc"
	"act/rpc/ms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveProcDefLogic {
	return &SaveProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程定义
func (l *SaveProcDefLogic) SaveProcDef(in *ms.ProcDefReq) (*ms.ProcDefReply, error) {
	// todo: add your logic here and delete this line
	log.Println("SaveProcDef 流程定义------rpc 222222")
	log.Println("req", in)
	// 插入 start
	// tx, err := l.svcCtx.ActStore.Tx(l.ctx)
	// if err != nil {
	// 	return &ms.ProcDefReply{}, err
	// }

	// _, err = tx.ProcDef.Create().
	// 	SetName(in.Name).SetCode(in.Code).SetYewuName(in.YewuName).SetYewuFormID(in.YewuName).
	// 	SetRemainHours(int(in.RemainHours)).SetResource(in.Resource).
	// 	SetCreateUserID(111).SetCreateUserName("ling").SetCreateTime(time.Now()).SetVersion(1).SetTargetID(222).
	// 	Save(l.ctx)

	// if err != nil {
	// 	tx.Rollback()
	// 	return &ms.ProcDefReply{}, err
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	return &ms.ProcDefReply{}, err
	// }
	// 插入 end

	// 查询 start
	a, err := l.svcCtx.ActStore.ProcDef.Query().Where(procdef.IDIn(800650655701041152)).Only(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("aaaaaa11111111", a)
	errCache := l.svcCtx.ActStore.Cache.Add(l.ctx, "cache.Key{}", cache.Entry{Columns: []string{"111", "222"}}, time.Duration(1111*time.Second))
	if errCache != nil {
		return nil, err
	}
	log.Println("a.Resource",a.Resource)
	log.Println("aaaaaa2222222222", a)
	// 查询 end

	return &ms.ProcDefReply{}, nil
}
