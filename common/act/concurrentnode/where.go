// Code generated by ent, DO NOT EDIT.

package concurrentnode

import (
	"act/common/act/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProcInstID applies equality check predicate on the "proc_inst_id" field. It's identical to ProcInstIDEQ.
func ProcInstID(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcInstID), v))
	})
}

// ProcDefID applies equality check predicate on the "proc_def_id" field. It's identical to ProcDefIDEQ.
func ProcDefID(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcDefID), v))
	})
}

// NodeID applies equality check predicate on the "node_id" field. It's identical to NodeIDEQ.
func NodeID(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNodeID), v))
	})
}

// NodeInfo applies equality check predicate on the "node_info" field. It's identical to NodeInfoEQ.
func NodeInfo(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNodeInfo), v))
	})
}

// PrevID applies equality check predicate on the "prev_id" field. It's identical to PrevIDEQ.
func PrevID(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrevID), v))
	})
}

// NextID applies equality check predicate on the "next_id" field. It's identical to NextIDEQ.
func NextID(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextID), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ProcInstIDEQ applies the EQ predicate on the "proc_inst_id" field.
func ProcInstIDEQ(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDNEQ applies the NEQ predicate on the "proc_inst_id" field.
func ProcInstIDNEQ(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDIn applies the In predicate on the "proc_inst_id" field.
func ProcInstIDIn(vs ...int64) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProcInstID), v...))
	})
}

// ProcInstIDNotIn applies the NotIn predicate on the "proc_inst_id" field.
func ProcInstIDNotIn(vs ...int64) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProcInstID), v...))
	})
}

// ProcInstIDGT applies the GT predicate on the "proc_inst_id" field.
func ProcInstIDGT(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDGTE applies the GTE predicate on the "proc_inst_id" field.
func ProcInstIDGTE(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDLT applies the LT predicate on the "proc_inst_id" field.
func ProcInstIDLT(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcInstID), v))
	})
}

// ProcInstIDLTE applies the LTE predicate on the "proc_inst_id" field.
func ProcInstIDLTE(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcInstID), v))
	})
}

// ProcDefIDEQ applies the EQ predicate on the "proc_def_id" field.
func ProcDefIDEQ(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcDefID), v))
	})
}

// ProcDefIDNEQ applies the NEQ predicate on the "proc_def_id" field.
func ProcDefIDNEQ(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcDefID), v))
	})
}

// ProcDefIDIn applies the In predicate on the "proc_def_id" field.
func ProcDefIDIn(vs ...int64) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProcDefID), v...))
	})
}

// ProcDefIDNotIn applies the NotIn predicate on the "proc_def_id" field.
func ProcDefIDNotIn(vs ...int64) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProcDefID), v...))
	})
}

// ProcDefIDGT applies the GT predicate on the "proc_def_id" field.
func ProcDefIDGT(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcDefID), v))
	})
}

// ProcDefIDGTE applies the GTE predicate on the "proc_def_id" field.
func ProcDefIDGTE(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcDefID), v))
	})
}

// ProcDefIDLT applies the LT predicate on the "proc_def_id" field.
func ProcDefIDLT(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcDefID), v))
	})
}

// ProcDefIDLTE applies the LTE predicate on the "proc_def_id" field.
func ProcDefIDLTE(v int64) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcDefID), v))
	})
}

// NodeIDEQ applies the EQ predicate on the "node_id" field.
func NodeIDEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNodeID), v))
	})
}

// NodeIDNEQ applies the NEQ predicate on the "node_id" field.
func NodeIDNEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNodeID), v))
	})
}

// NodeIDIn applies the In predicate on the "node_id" field.
func NodeIDIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNodeID), v...))
	})
}

// NodeIDNotIn applies the NotIn predicate on the "node_id" field.
func NodeIDNotIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNodeID), v...))
	})
}

// NodeIDGT applies the GT predicate on the "node_id" field.
func NodeIDGT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNodeID), v))
	})
}

// NodeIDGTE applies the GTE predicate on the "node_id" field.
func NodeIDGTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNodeID), v))
	})
}

// NodeIDLT applies the LT predicate on the "node_id" field.
func NodeIDLT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNodeID), v))
	})
}

// NodeIDLTE applies the LTE predicate on the "node_id" field.
func NodeIDLTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNodeID), v))
	})
}

// NodeIDContains applies the Contains predicate on the "node_id" field.
func NodeIDContains(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNodeID), v))
	})
}

// NodeIDHasPrefix applies the HasPrefix predicate on the "node_id" field.
func NodeIDHasPrefix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNodeID), v))
	})
}

// NodeIDHasSuffix applies the HasSuffix predicate on the "node_id" field.
func NodeIDHasSuffix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNodeID), v))
	})
}

// NodeIDEqualFold applies the EqualFold predicate on the "node_id" field.
func NodeIDEqualFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNodeID), v))
	})
}

// NodeIDContainsFold applies the ContainsFold predicate on the "node_id" field.
func NodeIDContainsFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNodeID), v))
	})
}

// NodeInfoEQ applies the EQ predicate on the "node_info" field.
func NodeInfoEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoNEQ applies the NEQ predicate on the "node_info" field.
func NodeInfoNEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoIn applies the In predicate on the "node_info" field.
func NodeInfoIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNodeInfo), v...))
	})
}

// NodeInfoNotIn applies the NotIn predicate on the "node_info" field.
func NodeInfoNotIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNodeInfo), v...))
	})
}

// NodeInfoGT applies the GT predicate on the "node_info" field.
func NodeInfoGT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoGTE applies the GTE predicate on the "node_info" field.
func NodeInfoGTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoLT applies the LT predicate on the "node_info" field.
func NodeInfoLT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoLTE applies the LTE predicate on the "node_info" field.
func NodeInfoLTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoContains applies the Contains predicate on the "node_info" field.
func NodeInfoContains(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoHasPrefix applies the HasPrefix predicate on the "node_info" field.
func NodeInfoHasPrefix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoHasSuffix applies the HasSuffix predicate on the "node_info" field.
func NodeInfoHasSuffix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoIsNil applies the IsNil predicate on the "node_info" field.
func NodeInfoIsNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNodeInfo)))
	})
}

// NodeInfoNotNil applies the NotNil predicate on the "node_info" field.
func NodeInfoNotNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNodeInfo)))
	})
}

// NodeInfoEqualFold applies the EqualFold predicate on the "node_info" field.
func NodeInfoEqualFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNodeInfo), v))
	})
}

// NodeInfoContainsFold applies the ContainsFold predicate on the "node_info" field.
func NodeInfoContainsFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNodeInfo), v))
	})
}

// PrevIDEQ applies the EQ predicate on the "prev_id" field.
func PrevIDEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrevID), v))
	})
}

// PrevIDNEQ applies the NEQ predicate on the "prev_id" field.
func PrevIDNEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrevID), v))
	})
}

// PrevIDIn applies the In predicate on the "prev_id" field.
func PrevIDIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrevID), v...))
	})
}

// PrevIDNotIn applies the NotIn predicate on the "prev_id" field.
func PrevIDNotIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrevID), v...))
	})
}

// PrevIDGT applies the GT predicate on the "prev_id" field.
func PrevIDGT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrevID), v))
	})
}

// PrevIDGTE applies the GTE predicate on the "prev_id" field.
func PrevIDGTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrevID), v))
	})
}

// PrevIDLT applies the LT predicate on the "prev_id" field.
func PrevIDLT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrevID), v))
	})
}

// PrevIDLTE applies the LTE predicate on the "prev_id" field.
func PrevIDLTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrevID), v))
	})
}

// PrevIDContains applies the Contains predicate on the "prev_id" field.
func PrevIDContains(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrevID), v))
	})
}

// PrevIDHasPrefix applies the HasPrefix predicate on the "prev_id" field.
func PrevIDHasPrefix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrevID), v))
	})
}

// PrevIDHasSuffix applies the HasSuffix predicate on the "prev_id" field.
func PrevIDHasSuffix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrevID), v))
	})
}

// PrevIDEqualFold applies the EqualFold predicate on the "prev_id" field.
func PrevIDEqualFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrevID), v))
	})
}

// PrevIDContainsFold applies the ContainsFold predicate on the "prev_id" field.
func PrevIDContainsFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrevID), v))
	})
}

// NextIDEQ applies the EQ predicate on the "next_id" field.
func NextIDEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextID), v))
	})
}

// NextIDNEQ applies the NEQ predicate on the "next_id" field.
func NextIDNEQ(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNextID), v))
	})
}

// NextIDIn applies the In predicate on the "next_id" field.
func NextIDIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNextID), v...))
	})
}

// NextIDNotIn applies the NotIn predicate on the "next_id" field.
func NextIDNotIn(vs ...string) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNextID), v...))
	})
}

// NextIDGT applies the GT predicate on the "next_id" field.
func NextIDGT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNextID), v))
	})
}

// NextIDGTE applies the GTE predicate on the "next_id" field.
func NextIDGTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNextID), v))
	})
}

// NextIDLT applies the LT predicate on the "next_id" field.
func NextIDLT(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNextID), v))
	})
}

// NextIDLTE applies the LTE predicate on the "next_id" field.
func NextIDLTE(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNextID), v))
	})
}

// NextIDContains applies the Contains predicate on the "next_id" field.
func NextIDContains(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNextID), v))
	})
}

// NextIDHasPrefix applies the HasPrefix predicate on the "next_id" field.
func NextIDHasPrefix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNextID), v))
	})
}

// NextIDHasSuffix applies the HasSuffix predicate on the "next_id" field.
func NextIDHasSuffix(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNextID), v))
	})
}

// NextIDEqualFold applies the EqualFold predicate on the "next_id" field.
func NextIDEqualFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNextID), v))
	})
}

// NextIDContainsFold applies the ContainsFold predicate on the "next_id" field.
func NextIDContainsFold(v string) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNextID), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int32) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int32) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// IsDelIn applies the In predicate on the "is_del" field.
func IsDelIn(vs ...int32) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIsDel), v...))
	})
}

// IsDelNotIn applies the NotIn predicate on the "is_del" field.
func IsDelNotIn(vs ...int32) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIsDel), v...))
	})
}

// IsDelGT applies the GT predicate on the "is_del" field.
func IsDelGT(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDel), v))
	})
}

// IsDelGTE applies the GTE predicate on the "is_del" field.
func IsDelGTE(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDel), v))
	})
}

// IsDelLT applies the LT predicate on the "is_del" field.
func IsDelLT(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDel), v))
	})
}

// IsDelLTE applies the LTE predicate on the "is_del" field.
func IsDelLTE(v int32) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDel), v))
	})
}

// IsDelIsNil applies the IsNil predicate on the "is_del" field.
func IsDelIsNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsDel)))
	})
}

// IsDelNotNil applies the NotNil predicate on the "is_del" field.
func IsDelNotNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsDel)))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.ConcurrentNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ConcurrentNode) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ConcurrentNode) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
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
func Not(p predicate.ConcurrentNode) predicate.ConcurrentNode {
	return predicate.ConcurrentNode(func(s *sql.Selector) {
		p(s.Not())
	})
}
