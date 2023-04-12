// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-wflow/kernel/ent/predicate"
	"go-wflow/kernel/ent/procinst"

	"github.com/qkbyte/ent/dialect/sql"
	"github.com/qkbyte/ent/dialect/sql/sqlgraph"
	"github.com/qkbyte/ent/schema/field"
)

// ProcinstDelete is the builder for deleting a Procinst entity.
type ProcinstDelete struct {
	config
	hooks    []Hook
	mutation *ProcinstMutation
}

// Where appends a list predicates to the ProcinstDelete builder.
func (pd *ProcinstDelete) Where(ps ...predicate.Procinst) *ProcinstDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProcinstDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProcinstMutation)
			if !ok {
				return nil, fmt.Errorf("意外突变类型 %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: 未初始化挂钩 (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProcinstDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProcinstDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: procinst.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: procinst.FieldID,
			},
		},
		UpdateUser: pd.mutation.updateUser,
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// ProcinstDeleteOne is the builder for deleting a single Procinst entity.
type ProcinstDeleteOne struct {
	pd *ProcinstDelete
}

// Exec executes the deletion query.
func (pdo *ProcinstDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{procinst.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProcinstDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}