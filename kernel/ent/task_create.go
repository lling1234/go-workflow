// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-wflow/kernel/ent/task"
	"time"

	"github.com/qkbyte/ent/dialect/sql/sqlgraph"
	"github.com/qkbyte/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetNodeId sets the "nodeId" field.
func (tc *TaskCreate) SetNodeId(s string) *TaskCreate {
	tc.mutation.SetNodeId(s)
	return tc
}

// SetNillableNodeId sets the "nodeId" field if the given value is not nil.
func (tc *TaskCreate) SetNillableNodeId(s *string) *TaskCreate {
	if s != nil {
		tc.SetNodeId(*s)
	}
	return tc
}

// SetProcInstID sets the "procInstID" field.
func (tc *TaskCreate) SetProcInstID(u uint64) *TaskCreate {
	tc.mutation.SetProcInstID(u)
	return tc
}

// SetCreateTime sets the "createTime" field.
func (tc *TaskCreate) SetCreateTime(t time.Time) *TaskCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreateTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetClaimTime sets the "claimTime" field.
func (tc *TaskCreate) SetClaimTime(t time.Time) *TaskCreate {
	tc.mutation.SetClaimTime(t)
	return tc
}

// SetNillableClaimTime sets the "claimTime" field if the given value is not nil.
func (tc *TaskCreate) SetNillableClaimTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetClaimTime(*t)
	}
	return tc
}

// SetIsFinished sets the "isFinished" field.
func (tc *TaskCreate) SetIsFinished(u uint64) *TaskCreate {
	tc.mutation.SetIsFinished(u)
	return tc
}

// SetNillableIsFinished sets the "isFinished" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsFinished(u *uint64) *TaskCreate {
	if u != nil {
		tc.SetIsFinished(*u)
	}
	return tc
}

// SetUpdateTime sets the "updateTime" field.
func (tc *TaskCreate) SetUpdateTime(t time.Time) *TaskCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (tc *TaskCreate) SetNillableUpdateTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetCreateUser sets the "createUser" field.
func (tc *TaskCreate) SetCreateUser(u uint64) *TaskCreate {
	tc.mutation.SetCreateUser(u)
	return tc
}

// SetNillableCreateUser sets the "createUser" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreateUser(u *uint64) *TaskCreate {
	if u != nil {
		tc.SetCreateUser(*u)
	}
	return tc
}

// SetUpdateUser sets the "updateUser" field.
func (tc *TaskCreate) SetUpdateUser(u uint64) *TaskCreate {
	tc.mutation.SetUpdateUser(u)
	return tc
}

// SetNillableUpdateUser sets the "updateUser" field if the given value is not nil.
func (tc *TaskCreate) SetNillableUpdateUser(u *uint64) *TaskCreate {
	if u != nil {
		tc.SetUpdateUser(*u)
	}
	return tc
}

// SetVersion sets the "version" field.
func (tc *TaskCreate) SetVersion(u uint64) *TaskCreate {
	tc.mutation.SetVersion(u)
	return tc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tc *TaskCreate) SetNillableVersion(u *uint64) *TaskCreate {
	if u != nil {
		tc.SetVersion(*u)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(u uint64) *TaskCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableID(u *uint64) *TaskCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("意外突变类型 %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: 未初始化挂钩 (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := task.DefaultCreateTime
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.IsFinished(); !ok {
		v := task.DefaultIsFinished
		tc.mutation.SetIsFinished(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := task.DefaultUpdateTime
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := task.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.ProcInstID(); !ok {
		return &ValidationError{Name: "procInstID", err: errors.New(`ent: missing required field "Task.procInstID"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: task.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.NodeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldNodeId,
		})
		_node.NodeId = value
	}
	if value, ok := tc.mutation.ProcInstID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: task.FieldProcInstID,
		})
		_node.ProcInstID = value
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.ClaimTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldClaimTime,
		})
		_node.ClaimTime = value
	}
	if value, ok := tc.mutation.IsFinished(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: task.FieldIsFinished,
		})
		_node.IsFinished = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: task.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := tc.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: task.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := tc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: task.FieldVersion,
		})
		_node.Version = value
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
				if !ok {
					return nil, fmt.Errorf("意外突变类型 %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
