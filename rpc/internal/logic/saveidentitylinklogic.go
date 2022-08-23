package logic

import (
	"context"
	"log"
	"time"

	"act/common/act"
	"act/common/act/procdef"
	"act/common/act/procinst"
	"act/common/act/task"

	"act/rpc/internal/flow"
	"act/rpc/internal/svc"
	"act/rpc/ms"

	"github.com/duke-git/lancet/v2/random"

	"github.com/mumushuiding/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveIdentityLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveIdentityLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveIdentityLinkLogic {
	return &SaveIdentityLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程开始
func (l *SaveIdentityLinkLogic) SaveIdentityLink(in *ms.IdentityLinkReq) (*ms.IdentityLinkReply, error) {
	// todo: add your logic here and delete this line
	log.Println("SaveIdentityLink 流程开始------rpc 222222")
	log.Println("req", in)
	var err error

	/*
		1. SELECT * FROM `act_proc_def`
		WHERE (yewu_form_id='1239348' and target_id=1727882)
		ORDER BY version desc
	*/
	// 根据yewu_form_id和target_id从act_proc_def表中，获取id
	store := l.svcCtx.ActStore

	procDefResult := store.ProcDef.Query().
		Where(procdef.YewuFormID("1239348"), procdef.TargetIDEQ(1727882)).
		Order(act.Asc(procdef.FieldVersion))

	procdefIdArr, err := procDefResult.IDs(l.ctx)
	if err != nil {
		return nil, err
	}
	procdefId := procdefIdArr[0]
	log.Println("procdefId", procdefId)

	// node解析
	node := &flow.Node{}
	// TODO 根据id查询resource
	var resourceStr string = `{"name":"发起人","type":"start","nodeId":"sid-startevent","childNode":{"type":"route","nodeId":"1659320135116","prevId":"sid-startevent","childNode":{"name":"审批人","type":"approver","nodeId":"1659320226296","prevId":"1659320135116","properties":{"actionerRules":[{"type":"target_management","approverNames":"大壮","approverIds":"107777777","actType":"or","level":2,"autoUp":true}]}},"conditionNodes":[{"name":"条件1","type":"condition","nodeId":"1659320135126","prevId":"1659320135116","childNode":{"name":"审批人","type":"approver","nodeId":"1659320194369","prevId":"1659320135126","properties":{"actionerRules":[{"type":"target_management","approverNames":"小春","approverIds":"122298283","actType":"and","level":1,"autoUp":true}]}},"properties":{"conditions":[[{"type":"dingtalk_actioner_range_condition","paramKey":"amount","paramLabel":"5000","upperBound":"3000"}]]}},{"name":"条件2","type":"condition","nodeId":"1659320135136","prevId":"1659320135126","childNode":{"name":"审批人","type":"approver","nodeId":"1659320201085","prevId":"1659320135136","properties":{"actionerRules":[{"type":"target_label","approverNames":"小龙","approverIds":"122298287","memberCount":1,"actType":"or"}]}},"properties":{"conditions":[[{"type":"dingtalk_actioner_range_condition","paramKey":"amount2","paramLabel":"5000","lowerBoundEqual":"3000"}]]}}]}}`
	procDefResource := act.ProcDef{Resource: resourceStr}
	err = util.Str2Struct(procDefResource.Resource, node)
	if err != nil {
		return nil, err
	}
	level := 0 // 0 为开始节点
	log.Println("level", level)
	// ---------------------------------------------
	// 新建流程实例
	/*	 2.
	INSERT INTO `act_proc_inst`
	(`id`,`create_time`,`proc_def_id`,`title`,`node_id`,`task_id`,
	`start_time`,`end_time`,`start_user_id`,`start_user_name`,`target_id`,`data_id`,`state`)
	VALUES (6817417568194756608,'2022-08-19 17:08:33',8725662109012164608,'','',0,
	'2022-08-19 17:08:33','2022-08-19 17:08:33',12025,'yxq',1727882,0,1)
	*/
	procInstInfo, err := store.ProcInst.Create().SetCreateTime(time.Now()).SetProcDefID(11).SetTitle("111").
		SetNodeID("111").SetTaskID(11).SetStartTime(time.Now()).SetEndTime(time.Now().AddDate(0, 0, 1)).SetStartUserID(12025).SetStartUserName("ling").
		SetTargetID(1727882).SetDataID(0).SetState(1).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	procInstId := procInstInfo.ID
	log.Println("procInstId", procInstId)
	/* 3.
	// 生成执行流，一串运行节点
	INSERT INTO `act_execution`
	(`id`,`create_time`,`proc_inst_id`,`proc_def_id`,`is_active`,`start_time`,`node_infos`)
	VALUES (6817391092489617408,'2022-08-19 17:08:33',6817417568194756608,8725662109012164608,
	1,'2022-08-19 17:08:33','[{"nodeId":"开始","type":"starter","aproverType":"","memberCount":0,"level":0,"actType":"","approverNames":"","approverIds":"⻹"},{"nodeId":"1659320194369","type":"target_management","aproverType":"approver","memberCount":1,"level":1,"actType":"and","approverNames":"小春","approverIds":"122298283"},{"nodeId":"1659320226296","type":"target_management","aproverType":"approver","memberCount":1,"level":2,"actType":"or","approverNames":"大壮","approverIds":"107777777"},{"nodeId":"结束","type":"","aproverType":"","memberCount":0,"level":3,"actType":"","approverNames":"","approverIds":""}]')
	*/
	execution, err := store.Execution.Create().SetCreateTime(time.Now()).
		SetProcInstID(int64(random.RandInt(1, 100000000000))).
		SetProcDefID(8725662109012164608).SetIsActive(1).
		SetStartTime(time.Now()).SetNodeInfos(resourceStr).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("execution", execution)
	/* 4.
	// -----------------生成新任务-------------
	INSERT INTO `act_task`
	(`id`,`create_time`,`node_id`,`level`,`proc_inst_id`,`claim_time`,
	`member_count`,`agree_num`,`act_type`,`is_finished`)
	VALUES (6817391092489617408,'2022-08-19 17:08:33','开始',0,6817417568194756608,'2022-08-19 17:08:33',
	1,1,'or',1)
	*/
	_, err = store.Task.Create().
		SetCreateTime(time.Now()).SetNodeID("开始").SetLevel(0).SetProcInstID(6817417568194756608).
		SetClaimTime(time.Now()).SetMemberCount(1).SetAgreeNum(1).SetActType("or").SetIsFinished(1).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	// TODO
	// SELECT `un_complete_num` FROM `act_task`  WHERE (id = 6817391092489617408)
	// SELECT `is_finished` FROM `act_task`  WHERE (id = 9212524636351135751)
	taskFinsh := store.Task.Query().Where(task.ID(6817391092489617408))
	taskFinsh.IDs(l.ctx)
	/* 5.

	INSERT INTO `act_candidate`
	(`id`,`create_time`,`user_id`,`user_name`,`proc_inst_id`,`target_id`,`task_id`,`is_finish`)
	VALUES (6817318104469700608,'2022-08-19 17:08:33',122298283,'小春',
	6817417568194756608,1727882,9212524636351135751,2)
	*/
	/* 6.
	// 流程移动到下一环节
	UPDATE `act_proc_inst`
	SET `is_finished` = 2, `node_id` = '1659320194369', `task_id` = 9212524636351135751
	*/
	// 6.1 根据id和ProcDefID查询获取流程实例
	procdefList, err := store.ProcInst.Query().Where(procinst.ID(826790988067274752), procinst.ProcDefID(800650655701041152)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("procdefList", procdefList)
	// 6.2 更新
	procdefUpdate, err := procdefList.Update().SetIsFinished(2).SetNodeID("1659320194369").SetTaskID(9212524636351135751).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("procdefUpdate", procdefUpdate)

	// TODO下一个节点 同意 MoveToNextStage或驳回MoveToPrevStage 
	return &ms.IdentityLinkReply{}, nil
}
