package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// IdentityLink holds the schema definition for the IdentityLink entity.
type IdentityLink struct {
	ent.Schema
}

// Fields of the IdentityLink.
func (IdentityLink) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("审批人id").Optional(),
		field.String("user_name").MaxLen(50).Comment("审批人姓名").Optional(),
		field.Int("step").Comment("审批步数").Optional(),
		field.Int64("proc_inst_id").Comment("流程实例id").Optional(),
		field.Int64("target_id").Comment("岗位id").Optional(),
		field.String("comment").MaxLen(500).Comment("评论").Optional(),
		field.Int64("task_id").Comment("节点任务").Optional(),
		field.Int("result").Comment("审批结果").Optional(),
		field.Time("create_time").Default(time.Now()).Comment("创建时间").Optional(),
		field.Int8("is_del").Default(0).Comment("是否删除,0:未删除,1:已删除").Optional(),
		field.Int8("is_deal").Default(0).Comment("是否已审批 ,0:未审批,1:已审批").Optional(),
	}
}

// Edges of the IdentityLink.
func (IdentityLink) Edges() []ent.Edge {
	return nil
}
