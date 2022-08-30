package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("node_id").MaxLen(50).Comment("节点id").Optional(),
		field.Int("level").Comment("流程层级").Optional(),
		field.Int("step").Comment("流程步数").Optional(),
		field.Int64("proc_inst_id").Comment("流程实例id"),
		field.Time("create_time").Default(time.Now()).Comment("任务创建时间").Optional(),
		field.Time("claim_time").Default(time.Now()).Comment("节点最新审批时间").Optional(),
		field.Int("member_count").Comment("需审批人数").Optional(),
		field.Int("un_complete_num").Comment("未审批人数").Optional(),
		field.Int("agree_num").Comment("已通过人数").Optional(),
		field.Int8("is_finished").Default(2).Comment("任务是否完成 2:未结束 1:已完成").Optional(),
		field.Enum("act_type").Values("and", "or").Default("or").Comment("会签or或签").Optional(),
		field.Int64("data_id").Comment("流程绑定数据ID").Optional(),
		field.Int("is_del").Default(0).Comment("是否删除").Optional(),
		field.Time("update_time").Default(time.Now()).Comment("流程修改时间").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
