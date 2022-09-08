package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ProcDef holds the schema definition for the ProcDef entity.
type ProcDef struct {
	ent.Schema
}

// Table Name
func (ProcDef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_proc_def"},
	}
}

// Fields of the ProcDef.
func (ProcDef) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").MaxLen(200).Comment("流程名称").Optional(),
		field.String("code").MaxLen(100).Comment("流程编码").Optional(),
		field.Int32("version").Comment("版本").Optional(),
		field.String("resource").MaxLen(10000).Comment("流程图数据"),
		field.Int64("create_user_id").Comment("创建人id").Optional(),
		field.String("create_user_name").MaxLen(50).Comment("创建人名称").Optional(),
		field.Time("create_time").Default(time.Now()).Comment("流程创建时间").Optional(),
		field.Int64("target_id").Comment("组织ID").Optional(),
		field.String("form_id").MaxLen(50).Comment("业务表单ID").Optional(),
		field.String("form_name").MaxLen(50).Comment("业务表单名称").Optional(),
		field.Int32("remain_hours").Default(0).Comment("审批限定时间").Optional(),
		field.Int32("is_del").Default(0).Comment("是否删除,0:未删除,1:已删除").Optional(),
		field.Int32("is_active").Default(0).Comment("流程是否生效,0:否,1:是").Optional(),
		field.Time("update_time").Default(time.Now()).Comment("流程修改时间").Optional(),
		field.Time("del_time").Comment("流程删除时间").Optional(),
		field.Int64("del_user_id").Comment("删除人id").Optional(),
		field.Int64("update_user_id").Comment("修改人id").Optional(),
	}
}

// Edges of the ProcDef.
func (ProcDef) Edges() []ent.Edge {
	return nil
}
