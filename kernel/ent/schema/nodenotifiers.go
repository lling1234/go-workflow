package schema

import (
	"go-wflow/common/utils/snowflake"
	"time"

	"github.com/qkbyte/ent"
	"github.com/qkbyte/ent/schema/field"
)

// Nodenotifiers holds the schema definition for the Nodenotifiers entity.
type Nodenotifiers struct {
	ent.Schema
}

// Config of the Nodenotifiers.
func (Nodenotifiers) Config() ent.Config {
	return ent.Config{Table: "wflow_node_notifiers"}
}

// Fields of the Nodenotifiers Mixin.
func (Nodenotifiers) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Nodenotifiers.
func (Nodenotifiers) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(snowflake.SnowGen.NextVal).Comment("雪花id"),
		field.Uint64("userID").Comment("抄送人id").Optional(),
		field.Uint64("procInstID").Comment("流程实例id"),
		field.Uint64("targetID").Comment("组织id").Optional(),
		field.String("comment").MaxLen(500).Comment("评论").Optional(),

		field.Time("createTime").Default(time.Now()).Comment("创建时间").Optional(),
		field.Uint64("isPermit").Default(0).Comment("是否允许查看 ,0:不允许，1：允许").Optional(),
		field.Uint64("isDeal").Default(0).Comment("是否已审批 ,0:未审批,1:已审批").Optional(),
		field.Time("updateTime").Default(time.Now()).Comment("流程修改时间").Optional(),

		field.Uint64("createUser").Comment("创建人id").Optional(),
		field.Uint64("updateUser").Comment("修改人id").Optional(),
		field.Uint64("version").Comment("版本").Optional(),

	}
}

// Edges of the Nodenotifiers.
func (Nodenotifiers) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Nodenotifiers.
func (Nodenotifiers) Indexes() []ent.Index {
	return []ent.Index{}
}

// Hooks of the Nodenotifiers.
func (Nodenotifiers) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policy of the Nodenotifiers.
func (Nodenotifiers) Policy() ent.Policy {
	return nil
}
