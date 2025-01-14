// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-wflow/kernel/ent/nodenotifiers"
	"go-wflow/kernel/ent/predicate"

	"github.com/qkbyte/ent/dialect/sql"
	"github.com/qkbyte/ent/dialect/sql/sqlgraph"
	"github.com/qkbyte/ent/schema/field"
)

// NodenotifiersDelete is the builder for deleting a Nodenotifiers entity.
type NodenotifiersDelete struct {
	config
	hooks    []Hook
	mutation *NodenotifiersMutation
}

// Where appends a list predicates to the NodenotifiersDelete builder.
func (nd *NodenotifiersDelete) Where(ps ...predicate.Nodenotifiers) *NodenotifiersDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NodenotifiersDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nd.hooks) == 0 {
		affected, err = nd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodenotifiersMutation)
			if !ok {
				return nil, fmt.Errorf("意外突变类型 %T", m)
			}
			nd.mutation = mutation
			affected, err = nd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nd.hooks) - 1; i >= 0; i-- {
			if nd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: 未初始化挂钩 (forgotten import ent/runtime?)")
			}
			mut = nd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NodenotifiersDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NodenotifiersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: nodenotifiers.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: nodenotifiers.FieldID,
			},
		},
		UpdateUser: nd.mutation.updateUser,
	}
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
}

// NodenotifiersDeleteOne is the builder for deleting a single Nodenotifiers entity.
type NodenotifiersDeleteOne struct {
	nd *NodenotifiersDelete
}

// Exec executes the deletion query.
func (ndo *NodenotifiersDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nodenotifiers.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NodenotifiersDeleteOne) ExecX(ctx context.Context) {
	ndo.nd.ExecX(ctx)
}
