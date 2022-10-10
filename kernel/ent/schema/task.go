package schema

import (
	"go-wflow/common/utils/snowflake"
	"time"

	"github.com/qkbyte/ent"
	"github.com/qkbyte/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Config of the Task.
func (Task) Config() ent.Config {
	return ent.Config{Table: "wflow_task"}
}

// Fields of the Task Mixin.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(snowflake.SnowGen.NextVal).Comment("雪花id"),
		field.String("nodeId").Comment("节点id").Optional(),
		field.Uint64("procInstID").Comment("流程实例id"),
		field.Time("createTime").Default(time.Now()).Comment("任务创建时间").Optional(),
		field.Time("claimTime").Comment("节点最新审批时间").Optional(),

		field.Uint64("isFinished").Default(0).Comment("任务是否完成 0:未结束 1:已完成").Optional(),
		field.Time("updateTime").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Uint64("createUser").Comment("创建人id").Optional(),
		field.Uint64("updateUser").Comment("修改人id").Optional(),
		field.Uint64("version").Comment("版本").Optional(),

	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Task.
func (Task) Indexes() []ent.Index {
	return []ent.Index{}
}

// Hooks of the Task.
func (Task) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policy of the Task.
func (Task) Policy() ent.Policy {
	return nil
}
