package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Execution holds the schema definition for the Execution entity.
type Execution struct {
	ent.Schema
}

// Fields of the Execution.
func (Execution) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("proc_inst_id").Comment("流程实例id"),
		field.Int64("proc_def_id").Comment("流程定义id"),
		field.String("node_infos").MaxLen(5000).Comment("节点审批执行顺序").Optional(),
		field.Time("start_time").Default(time.Now()).Comment("开始时间").Optional(),
		field.Int8("is_del").Default(0).Comment("是否删除,0:未删除,1:已删除").Optional(),
		field.Time("create_time").Default(time.Now()).Comment("创建时间").Optional(),
		field.Time("update_time").Default(time.Now()).Comment("流程修改时间").Optional(),
	}
}

// Edges of the Execution.
func (Execution) Edges() []ent.Edge {
	return nil
}
