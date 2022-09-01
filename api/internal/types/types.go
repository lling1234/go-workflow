// Code generated by goctl. DO NOT EDIT.
package types

import "act/api/flow"

type SaveProcDef struct {
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	FormId      string     `json:"formId"`
	FormName    string     `json:"formName"`
	RemainHours int32      `json:"remainHours"`
	Resource    *flow.Node `json:"resource"`
}

type ProcDefIdReq struct {
	ProcDefId int64 `json:"formId"`
}

type StartProcess struct {
	Title  string `json:"name"`
	FormId string `json:"formId"`
	DataId int64  `json:"dataId"`
}

type CompeleteTask struct {
	DataId  int64  `json:"dataId"`
	Result  int    `json:"result"`
	Comment string `json:"comment"`
}

type Withdraw struct {
	DataId int64 `json:"dataId"`
}

type FindByTargetId struct {
	TargetId int64 `json:"targetId"`
}

type FindMyStartProcess struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

type FindMyApproveProcess struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

type CommonResponse struct {
	Code    int32       `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}
