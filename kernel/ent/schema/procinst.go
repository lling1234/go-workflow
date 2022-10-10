package schema

import (
	"go-wflow/common/utils/snowflake"
	"time"

	"github.com/qkbyte/ent"
	"github.com/qkbyte/ent/schema/field"
)

// Procinst holds the schema definition for the Procinst entity.
type Procinst struct {
	ent.Schema
}

// Config of the Procinst.
func (Procinst) Config() ent.Config {
	return ent.Config{Table: "wflow_proc_inst"}
}

// Fields of the Procinst Mixin.
func (Procinst) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Procinst.
func (Procinst) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(snowflake.SnowGen.NextVal).Comment("雪花id"),
		field.Uint64("procDefID").Comment("流程定义id"),
		field.Uint64("refID").Comment("关联id").Optional(),
		field.String("title").MaxLen(100).Comment("发起流程标题").Optional(),
		field.String("code").MaxLen(100).Comment("流程编号").Optional(),

		field.Uint64("targetID").Comment("组织id").Optional(),
		field.String("resource").MaxLen(10000).Comment("流程图数据").Optional(),
		field.String("nodeID").MaxLen(50).Comment("当前审批节点id").Optional(),
		field.Uint64("taskID").Comment("当前审批任务id").Optional(),
		field.String("conNodeIDs").MaxLen(1000).Comment("当前审批节点id（并行流程）").Optional(),

		field.String("conTaskIDs").MaxLen(500).Comment("当前审批任务id（并行流程）").Optional(),
		field.Uint64("isFinished").Default(0).Comment("流程是否结束,0:未结束,1:已结束").Optional(),
		field.Uint64("state").
			Comment("流程状态,类型为:1待处理、2处理中、3驳回、4已撤回、5未通过、6已通过、7废弃").
			Optional(),
		field.Uint64("dataID").Comment("流程绑定数据id").Optional(),
		field.Uint64("updateUser").Comment("修改人id").Optional(),

		field.Time("createTime").Default(time.Now()).Comment("创建时间").Optional(),
		field.Time("finishTime").Default(time.Now()).Comment("流程结束时间").Optional(),
		field.Uint64("createUser").Comment("创建人id").Optional(),
		field.String("createUsername").MaxLen(50).Comment("发起人姓名").Optional(),
		field.Uint64("remainHours").Comment("审批限定时间").Optional(),
		
		field.Time("updateTime").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Uint64("flowType").Default(1).Comment("流程类型 1、普通流程  2、并行流程").Optional(),
		field.String("remark").MaxLen(1000).Comment("备注").Optional(),
		field.Time("delTime").Comment("流程删除时间").Optional(),
		field.Uint64("delUser").Comment("删除人id").Optional(),

		field.Uint64("version").Comment("版本").Optional(),

	}
}

// Edges of the Procinst.
func (Procinst) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Procinst.
func (Procinst) Indexes() []ent.Index {
	return []ent.Index{}
}

// Hooks of the Procinst.
func (Procinst) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policy of the Procinst.
func (Procinst) Policy() ent.Policy {
	return nil
}
