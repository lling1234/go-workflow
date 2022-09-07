package act

import (
<<<<<<< HEAD
=======
	"act/rpc/types/act"
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

<<<<<<< HEAD
<<<<<<<< HEAD:api/internal/logic/act/findbydataidlogic.go
type FindByDataIdLogic struct {
========
type FindMyFinishLogic struct {
>>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb:api/internal/logic/act/findmyfinishlogic.go
=======
type FindByDataIdLogic struct {
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

<<<<<<< HEAD
<<<<<<<< HEAD:api/internal/logic/act/findbydataidlogic.go
func NewFindByDataIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByDataIdLogic {
	return &FindByDataIdLogic{
========
func NewFindMyFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyFinishLogic {
	return &FindMyFinishLogic{
>>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb:api/internal/logic/act/findmyfinishlogic.go
=======
func NewFindByDataIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByDataIdLogic {
	return &FindByDataIdLogic{
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

<<<<<<< HEAD
<<<<<<<< HEAD:api/internal/logic/act/findbydataidlogic.go
func (l *FindByDataIdLogic) FindByDataId(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
========
func (l *FindMyFinishLogic) FindMyFinish(req *types.PageReq) (resp *types.CommonResponse, err error) {
>>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb:api/internal/logic/act/findmyfinishlogic.go
	// todo: add your logic here and delete this line

	return
=======
func (l *FindByDataIdLogic) FindByDataId(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
	reply, err := l.svcCtx.Rpc.FindProcInstByDataId(l.ctx, &act.DataIdReq{
		DataId: req.DataId,
	})

	return types.GetCommonResponse(err, reply)
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
}
