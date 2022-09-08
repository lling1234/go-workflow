// Code generated by ent, DO NOT EDIT.

package identitylink

import (
	"act/common/act/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// Step applies equality check predicate on the "step" field. It's identical to StepEQ.
func Step(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStep), v))
	})
}

// ProcInstID applies equality check predicate on the "proc_inst_id" field. It's identical to ProcInstIDEQ.
func ProcInstID(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcInstID), v))
	})
}

// TargetID applies equality check predicate on the "target_id" field. It's identical to TargetIDEQ.
func TargetID(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetID), v))
	})
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDeal applies equality check predicate on the "is_deal" field. It's identical to IsDealEQ.
func IsDeal(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeal), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserName), v))
	})
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserName), v...))
	})
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserName), v...))
	})
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserName), v))
	})
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserName), v))
	})
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserName), v))
	})
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserName), v))
	})
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserName), v))
	})
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserName), v))
	})
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserName), v))
	})
}

// UserNameIsNil applies the IsNil predicate on the "user_name" field.
func UserNameIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserName)))
	})
}

// UserNameNotNil applies the NotNil predicate on the "user_name" field.
func UserNameNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserName)))
	})
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserName), v))
	})
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserName), v))
	})
}

// StepEQ applies the EQ predicate on the "step" field.
func StepEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStep), v))
	})
}

// StepNEQ applies the NEQ predicate on the "step" field.
func StepNEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStep), v))
	})
}

// StepIn applies the In predicate on the "step" field.
func StepIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStep), v...))
	})
}

// StepNotIn applies the NotIn predicate on the "step" field.
func StepNotIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStep), v...))
	})
}

// StepGT applies the GT predicate on the "step" field.
func StepGT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStep), v))
	})
}

// StepGTE applies the GTE predicate on the "step" field.
func StepGTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStep), v))
	})
}

// StepLT applies the LT predicate on the "step" field.
func StepLT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStep), v))
	})
}

// StepLTE applies the LTE predicate on the "step" field.
func StepLTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStep), v))
	})
}

// StepIsNil applies the IsNil predicate on the "step" field.
func StepIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStep)))
	})
}

// StepNotNil applies the NotNil predicate on the "step" field.
func StepNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStep)))
	})
}

// ProcInstIDEQ applies the EQ predicate on the "proc_inst_id" field.
func ProcInstIDEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDNEQ applies the NEQ predicate on the "proc_inst_id" field.
func ProcInstIDNEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDIn applies the In predicate on the "proc_inst_id" field.
func ProcInstIDIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProcInstID), v...))
	})
}

// ProcInstIDNotIn applies the NotIn predicate on the "proc_inst_id" field.
func ProcInstIDNotIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProcInstID), v...))
	})
}

// ProcInstIDGT applies the GT predicate on the "proc_inst_id" field.
func ProcInstIDGT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDGTE applies the GTE predicate on the "proc_inst_id" field.
func ProcInstIDGTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDLT applies the LT predicate on the "proc_inst_id" field.
func ProcInstIDLT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDLTE applies the LTE predicate on the "proc_inst_id" field.
func ProcInstIDLTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcInstID), v))
	})
}

// TargetIDEQ applies the EQ predicate on the "target_id" field.
func TargetIDEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetID), v))
	})
}

// TargetIDNEQ applies the NEQ predicate on the "target_id" field.
func TargetIDNEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTargetID), v))
	})
}

// TargetIDIn applies the In predicate on the "target_id" field.
func TargetIDIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTargetID), v...))
	})
}

// TargetIDNotIn applies the NotIn predicate on the "target_id" field.
func TargetIDNotIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTargetID), v...))
	})
}

// TargetIDGT applies the GT predicate on the "target_id" field.
func TargetIDGT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTargetID), v))
	})
}

// TargetIDGTE applies the GTE predicate on the "target_id" field.
func TargetIDGTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTargetID), v))
	})
}

// TargetIDLT applies the LT predicate on the "target_id" field.
func TargetIDLT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTargetID), v))
	})
}

// TargetIDLTE applies the LTE predicate on the "target_id" field.
func TargetIDLTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTargetID), v))
	})
}

// TargetIDIsNil applies the IsNil predicate on the "target_id" field.
func TargetIDIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTargetID)))
	})
}

// TargetIDNotNil applies the NotNil predicate on the "target_id" field.
func TargetIDNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTargetID)))
	})
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComment), v))
	})
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldComment), v...))
	})
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldComment), v...))
	})
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComment), v))
	})
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComment), v))
	})
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComment), v))
	})
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComment), v))
	})
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComment), v))
	})
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComment), v))
	})
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComment), v))
	})
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldComment)))
	})
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldComment)))
	})
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComment), v))
	})
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComment), v))
	})
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskID), v))
	})
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTaskID), v...))
	})
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...int64) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTaskID), v...))
	})
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskID), v))
	})
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskID), v))
	})
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskID), v))
	})
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v int64) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskID), v))
	})
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResult), v))
	})
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResult), v...))
	})
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResult), v...))
	})
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResult), v))
	})
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResult), v))
	})
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResult), v))
	})
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResult), v))
	})
}

// ResultIsNil applies the IsNil predicate on the "result" field.
func ResultIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResult)))
	})
}

// ResultNotNil applies the NotNil predicate on the "result" field.
func ResultNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResult)))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// IsDelIn applies the In predicate on the "is_del" field.
func IsDelIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIsDel), v...))
	})
}

// IsDelNotIn applies the NotIn predicate on the "is_del" field.
func IsDelNotIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIsDel), v...))
	})
}

// IsDelGT applies the GT predicate on the "is_del" field.
func IsDelGT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDel), v))
	})
}

// IsDelGTE applies the GTE predicate on the "is_del" field.
func IsDelGTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDel), v))
	})
}

// IsDelLT applies the LT predicate on the "is_del" field.
func IsDelLT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDel), v))
	})
}

// IsDelLTE applies the LTE predicate on the "is_del" field.
func IsDelLTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDel), v))
	})
}

// IsDelIsNil applies the IsNil predicate on the "is_del" field.
func IsDelIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsDel)))
	})
}

// IsDelNotNil applies the NotNil predicate on the "is_del" field.
func IsDelNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsDel)))
	})
}

// IsDealEQ applies the EQ predicate on the "is_deal" field.
func IsDealEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeal), v))
	})
}

// IsDealNEQ applies the NEQ predicate on the "is_deal" field.
func IsDealNEQ(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeal), v))
	})
}

// IsDealIn applies the In predicate on the "is_deal" field.
func IsDealIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIsDeal), v...))
	})
}

// IsDealNotIn applies the NotIn predicate on the "is_deal" field.
func IsDealNotIn(vs ...int32) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIsDeal), v...))
	})
}

// IsDealGT applies the GT predicate on the "is_deal" field.
func IsDealGT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeal), v))
	})
}

// IsDealGTE applies the GTE predicate on the "is_deal" field.
func IsDealGTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeal), v))
	})
}

// IsDealLT applies the LT predicate on the "is_deal" field.
func IsDealLT(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeal), v))
	})
}

// IsDealLTE applies the LTE predicate on the "is_deal" field.
func IsDealLTE(v int32) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeal), v))
	})
}

// IsDealIsNil applies the IsNil predicate on the "is_deal" field.
func IsDealIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsDeal)))
	})
}

// IsDealNotNil applies the NotNil predicate on the "is_deal" field.
func IsDealNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsDeal)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.IdentityLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IdentityLink) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IdentityLink) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IdentityLink) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		p(s.Not())
	})
}
