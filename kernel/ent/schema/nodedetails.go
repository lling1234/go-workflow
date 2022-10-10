package schema

import (
	"go-wflow/common/utils/snowflake"
	"time"

	"github.com/qkbyte/ent"
	"github.com/qkbyte/ent/schema/field"
)

// NodeDetails holds the schema definition for the NodeDetails entity.
type NodeDetails struct {
	ent.Schema
}

// Config of the NodeDetails.
func (NodeDetails) Config() ent.Config {
	return ent.Config{Table: "wflow_node_details"}
}

// Fields of the NodeDetails Mixin.
func (NodeDetails) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the NodeDetails.
func (NodeDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(snowflake.SnowGen.NextVal).Comment("雪花id"),
		field.Uint64("procInstID").Comment("流程实例id"),
		field.String("nodeID").MaxLen(50).Comment("节点id").Optional(),
		field.String("nodeInfo").MaxLen(500).Comment("节点信息").Optional(),
		field.String("refuse").MaxLen(50).Comment("拒绝策略 直接结束、驳回至上一级、驳回至指定人").Optional(),

		field.String("prevID").MaxLen(500).Comment("上一节点ID").Optional(),
		field.String("nextID").MaxLen(500).Comment("下一节点ID").Optional(),
		field.String("mode").MaxLen(50).Default("or").Comment("会签or或签").Optional(),
		field.Time("createTime").Default(time.Now()).Comment("流程创建时间").Optional(),

		field.Time("updateTime").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Uint64("createUser").Comment("创建人id").Optional(),
		field.Uint64("updateUser").Comment("修改人id").Optional(),
		field.Uint64("version").Comment("版本").Optional(),
	}
}

// Edges of the NodeDetails.
func (NodeDetails) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the NodeDetails.
func (NodeDetails) Indexes() []ent.Index {
	return []ent.Index{}
}

// Hooks of the NodeDetails.
func (NodeDetails) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policy of the NodeDetails.
func (NodeDetails) Policy() ent.Policy {
	return nil
}
