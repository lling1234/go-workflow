package logic

import (
	"context"
	"log"

	commonact "act/common/act"
	"act/common/act/procinst"
	"act/common/page"
	commonpage "act/common/page"
	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllProcInstLogic {
	return &FindAllProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

<<<<<<< HEAD
// 流程实例查询,根据id查询，分页查询和根据title过滤
func (l *FindAllProcInstLogic) FindAllProcInst(in *act.IdRequest) (*act.CommonRpcRes, error) {
	log.Println("FindAllProcInst111111111", in)
	store := l.svcCtx.CommonStore.ProcInst.Query()
	//1.根据id查询
	if in.Id >= 1 {
		procinstInfo, err := store.Where(procinst.IDEQ(in.Id)).First(l.ctx)
		if err != nil {
			return nil, err
		}
		res, err := page.JsonResult(procinstInfo, err)
		if err != nil {
			return nil, err
		}
		return &act.CommonRpcRes{Json: res.Json}, nil
	}

	// 2.分页查询和根据title过滤
	total, err := store.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("total222222222", total)
	procinstArr, err := store.Where(procinst.TitleContains(in.Filter)).
		Order(commonact.Asc(procinst.FieldID)).
		Limit(int(in.Limit)).Offset(int(in.Offset)).
		All(l.ctx)
	log.Println("procinstArr3333333", procinstArr)
	res, err := commonpage.PageResult(&page.PageRequest{Limit: in.Limit, Offset: in.Offset}, int64(total), procinstArr, err)
	if err != nil {
		return nil, err
	}
	log.Println("----------------")
	log.Println("res444444444", res.Json)
	log.Println("----------------")
	log.Println("res555555555", res)
	return &act.CommonRpcRes{Json: res.Json}, nil
=======
func (l *FindAllProcInstLogic) FindAllProcInst(in *act.IdRequest) (*act.CommonRpcRes, error) {
	// todo: add your logic here and delete this line

	return &act.CommonRpcRes{}, nil
>>>>>>> 4a54178f732840048a6d221f9338a6e8f5d12ba3
}
