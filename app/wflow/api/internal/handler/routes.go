// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	wflow "go-wflow/app/wflow/api/internal/handler/wflow"
	"go-wflow/app/wflow/api/internal/svc"

	"github.com/qkbyte/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procdef/create",
				Handler: wflow.AddProcDefHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procdef/delete",
				Handler: wflow.DelProcDefHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procdef/update",
				Handler: wflow.UpdateProcDefHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procdef/findall",
				Handler: wflow.FindAllProcDefHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procdef/findone",
				Handler: wflow.FindOneProcDefHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procdef/setactive",
				Handler: wflow.SetActiveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/startview",
				Handler: wflow.StartViewHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/start",
				Handler: wflow.StartProcInstHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/complete",
				Handler: wflow.CompleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/view",
				Handler: wflow.ViewHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/findbydataid",
				Handler: wflow.FindByDataIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/withdraw",
				Handler: wflow.WithdrawHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/procinst/delete",
				Handler: wflow.DelProcInstHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/process/findmystart",
				Handler: wflow.FindMyStartHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/process/findmyunfinish",
				Handler: wflow.FindMyUnFinishHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/process/findmyfinish",
				Handler: wflow.FindMyFinishHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/process/findmypending",
				Handler: wflow.FindMyPendingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/process/findmyapproved",
				Handler: wflow.FindMyApprovedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/orginone/wflow/process/findovertime",
				Handler: wflow.FindOverTimeHandler(serverCtx),
			},
		},
	)
}
