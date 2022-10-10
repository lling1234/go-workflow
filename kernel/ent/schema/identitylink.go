package schema

import (
	"go-wflow/common/utils/snowflake"
	"time"

	"github.com/qkbyte/ent"
	"github.com/qkbyte/ent/schema/field"
)

// Identitylink holds the schema definition for the Identitylink entity.
type Identitylink struct {
	ent.Schema
}

// Config of the Identitylink.
func (Identitylink) Config() ent.Config {
	return ent.Config{Table: "wflow_identity_link"}
}

// Fields of the Identitylink Mixin.
func (Identitylink) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Identitylink.
func (Identitylink) Fields() []ent.Field {
	return []ent.Field{
		// TODO id转为string
		field.Uint64("id").DefaultFunc(snowflake.SnowGen.NextVal).Comment("雪花id"),
		field.Uint64("userID").Comment("审批人id").Optional(),
		field.String("userName").MaxLen(50).Comment("审批人姓名").Optional(),
		field.Uint64("procInstID").Comment("流程实例id"),
		field.Uint64("targetID").Comment("组织id").Optional(),

		field.Uint64("station").Comment("岗位id").Optional(),
		field.String("comment").MaxLen(500).Comment("评论").Optional(),
		field.Uint64("taskID").Comment("节点任务"),
		field.Uint64("result").Comment("审批结果 3驳回、5未通过、6已通过").Optional(),
		field.Time("createTime").Default(time.Now()).Comment("创建时间").Optional(),

		field.Uint64("isDeal").Default(0).Comment("是否已审批 ,0:未审批,1:已审批").Optional(),
		field.Time("updateTime").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Uint64("createUser").Comment("创建人id").Optional(),
		field.Uint64("updateUser").Comment("修改人id").Optional(),
		field.Uint64("attachmentID").Comment("附件id").Optional(),
		
		field.Uint64("version").Comment("版本").Optional(),
	}
}

// Edges of the Identitylink.
func (Identitylink) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Identitylink.
func (Identitylink) Indexes() []ent.Index {
	return []ent.Index{}
}

// Hooks of the Identitylink.
func (Identitylink) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policy of the Identitylink.
func (Identitylink) Policy() ent.Policy {
	return nil
}
