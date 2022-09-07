package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ProcInst holds the schema definition for the ProcInst entity.
type ProcInst struct {
	ent.Schema
}

// Table Name
func (ProcInst) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_proc_inst"},
	}
}

// Fields of the ProcInst.
func (ProcInst) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("proc_def_id").Comment("流程定义id"),
		field.String("title").MaxLen(100).Comment("发起流程标题").Optional(),
		field.String("code").MaxLen(100).Comment("流程编号").Optional(),
		field.Int64("target_id").Comment("组织id").Optional(),
		field.String("node_id").MaxLen(50).Comment("节点id").Optional(),
		field.Int64("task_id").Comment("当前审批任务id").Optional(),
		field.Time("start_time").Default(time.Now()).Comment("发起时间").Optional(),
		field.Time("end_time").Comment("结束时间").Optional(),
		field.Int64("start_user_id").Comment("发起人id").Optional(),
		field.String("start_user_name").MaxLen(50).Comment("发起人姓名").Optional(),
		field.Int32("is_finished").Default(0).Comment("流程是否结束,0:未结束,1:已结束").Optional(),
		field.Int32("state").
			Comment("流程状态,类型为:1待处理、2处理中、3驳回、4已撤回、5未通过、6已通过、7废弃").
			Optional(),
		field.Int64("data_id").Comment("流程绑定数据id").Optional(),
		field.Int32("is_del").Default(0).Comment("是否删除").Optional(),
		field.Time("create_time").Default(time.Now()).Comment("创建时间").Optional(),
		field.Int32("remain_hours").Comment("审批限定时间").Optional(),
		field.Time("update_time").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Int32("flow_type").Default(1).Comment("流程类型 1、普通流程  2、并行流程").Optional(),
		field.String("remark").MaxLen(1000).Comment("备注").Optional(),
		field.Time("del_time").Default(time.Now()).Comment("流程删除时间").Optional(),
		field.Int64("del_user_id").Comment("删除人id").Optional(),
	}
}

// Edges of the ProcInst.
func (ProcInst) Edges() []ent.Edge {
	return nil
}
