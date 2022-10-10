package schema

import (
	"go-wflow/common/utils/snowflake"
	"time"

	"github.com/qkbyte/ent"
	"github.com/qkbyte/ent/schema/field"
)

// Procdef holds the schema definition for the Procdef entity.
type Procdef struct {
	ent.Schema
}

// Config of the Procdef.
func (Procdef) Config() ent.Config {
	return ent.Config{Table: "wflow_proc_def"}
}

// Fields of the Procdef Mixin.
func (Procdef) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Procdef.
func (Procdef) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(snowflake.SnowGen.NextVal).Comment("雪花ID"),
		field.String("name").MaxLen(200).Comment("流程名称").Optional(),
		field.String("code").MaxLen(100).Comment("流程编码").Optional(),
		field.Uint64("version").Comment("版本").Optional(),
		field.String("resource").MaxLen(10000).Comment("流程图数据"),

		field.Uint64("createUser").Comment("创建人ID").Optional(),
		field.String("createUserName").MaxLen(50).Comment("创建人名称").Optional(),
		field.Time("createTime").Default(time.Now()).Comment("流程创建时间").Optional(),
		field.Uint64("targetID").Comment("组织ID").Optional(),
		field.Uint64("formID").Comment("业务表单ID").Optional(),

		field.String("formName").MaxLen(100).Comment("业务表单名称").Optional(),
		field.Uint64("appID").Comment("应用ID").Optional(),
		field.String("appName").MaxLen(100).Comment("应用名称").Optional(),
		field.Uint64("remainHours").Default(0).Comment("审批限定时间").Optional(),

		field.Uint64("isActive").Default(0).Comment("流程是否生效,0:否,1:是").Optional(),
		field.Time("updateTime").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Time("delTime").Comment("流程删除时间").Optional(),
		field.Uint64("delUser").Comment("删除人ID").Optional(),
		field.Uint64("updateUser").Comment("修改人ID").Optional(),
	}
}

// Edges of the Procdef.
func (Procdef) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Procdef.
func (Procdef) Indexes() []ent.Index {
	return []ent.Index{}
}

// Hooks of the Procdef.
func (Procdef) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policy of the Procdef.
func (Procdef) Policy() ent.Policy {
	return nil
}
