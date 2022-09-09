package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ConcurrentNode holds the schema definition for the ConcurrentNode entity.
type ConcurrentNode struct {
	ent.Schema
}

// Table Name
func (ConcurrentNode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_concurrent_node"},
	}
}

// Fields of the IdentityLink.
func (ConcurrentNode) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("proc_inst_id").Comment("流程实例id"),
		field.String("node_id").MaxLen(50).Comment("节点id").Optional(),
		field.String("node_info").MaxLen(500).Comment("节点信息").Optional(),
		field.String("prev_id").MaxLen(500).Comment("上一节点ID").Optional(),
		field.String("next_id").MaxLen(500).Comment("下一节点ID").Optional(),
		field.Int32("state").Comment("节点状态 1、未审批  2、正在审批  3、已审批完成").Optional(),
		field.Int32("is_del").Default(0).Comment("是否删除,0:未删除,1:已删除").Optional(),
		field.Time("create_time").Default(time.Now()).Comment("流程创建时间").Optional(),
		field.Time("update_time").Default(time.Now()).Comment("流程修改时间").Optional(),
	}
}

// Edges of the IdentityLink.
func (ConcurrentNode) Edges() []ent.Edge {
	return nil
}