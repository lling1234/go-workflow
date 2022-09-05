// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/task"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetNodeID sets the "node_id" field.
func (tc *TaskCreate) SetNodeID(s string) *TaskCreate {
	tc.mutation.SetNodeID(s)
	return tc
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableNodeID(s *string) *TaskCreate {
	if s != nil {
		tc.SetNodeID(*s)
	}
	return tc
}

// SetLevel sets the "level" field.
func (tc *TaskCreate) SetLevel(i int32) *TaskCreate {
	tc.mutation.SetLevel(i)
	return tc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (tc *TaskCreate) SetNillableLevel(i *int32) *TaskCreate {
	if i != nil {
		tc.SetLevel(*i)
	}
	return tc
}

// SetStep sets the "step" field.
func (tc *TaskCreate) SetStep(i int32) *TaskCreate {
	tc.mutation.SetStep(i)
	return tc
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (tc *TaskCreate) SetNillableStep(i *int32) *TaskCreate {
	if i != nil {
		tc.SetStep(*i)
	}
	return tc
}

// SetProcInstID sets the "proc_inst_id" field.
func (tc *TaskCreate) SetProcInstID(i int64) *TaskCreate {
	tc.mutation.SetProcInstID(i)
	return tc
}

// SetCreateTime sets the "create_time" field.
func (tc *TaskCreate) SetCreateTime(t time.Time) *TaskCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreateTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetClaimTime sets the "claim_time" field.
func (tc *TaskCreate) SetClaimTime(t time.Time) *TaskCreate {
	tc.mutation.SetClaimTime(t)
	return tc
}

// SetNillableClaimTime sets the "claim_time" field if the given value is not nil.
func (tc *TaskCreate) SetNillableClaimTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetClaimTime(*t)
	}
	return tc
}

// SetMemberApprover sets the "member_approver" field.
func (tc *TaskCreate) SetMemberApprover(s string) *TaskCreate {
	tc.mutation.SetMemberApprover(s)
	return tc
}

// SetNillableMemberApprover sets the "member_approver" field if the given value is not nil.
func (tc *TaskCreate) SetNillableMemberApprover(s *string) *TaskCreate {
	if s != nil {
		tc.SetMemberApprover(*s)
	}
	return tc
}

// SetAgreeApprover sets the "agree_approver" field.
func (tc *TaskCreate) SetAgreeApprover(s string) *TaskCreate {
	tc.mutation.SetAgreeApprover(s)
	return tc
}

// SetNillableAgreeApprover sets the "agree_approver" field if the given value is not nil.
func (tc *TaskCreate) SetNillableAgreeApprover(s *string) *TaskCreate {
	if s != nil {
		tc.SetAgreeApprover(*s)
	}
	return tc
}

// SetIsFinished sets the "is_finished" field.
func (tc *TaskCreate) SetIsFinished(i int8) *TaskCreate {
	tc.mutation.SetIsFinished(i)
	return tc
}

// SetNillableIsFinished sets the "is_finished" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsFinished(i *int8) *TaskCreate {
	if i != nil {
		tc.SetIsFinished(*i)
	}
	return tc
}

// SetMode sets the "mode" field.
func (tc *TaskCreate) SetMode(t task.Mode) *TaskCreate {
	tc.mutation.SetMode(t)
	return tc
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tc *TaskCreate) SetNillableMode(t *task.Mode) *TaskCreate {
	if t != nil {
		tc.SetMode(*t)
	}
	return tc
}

// SetDataID sets the "data_id" field.
func (tc *TaskCreate) SetDataID(i int64) *TaskCreate {
	tc.mutation.SetDataID(i)
	return tc
}

// SetNillableDataID sets the "data_id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDataID(i *int64) *TaskCreate {
	if i != nil {
		tc.SetDataID(*i)
	}
	return tc
}

// SetIsDel sets the "is_del" field.
func (tc *TaskCreate) SetIsDel(i int8) *TaskCreate {
	tc.mutation.SetIsDel(i)
	return tc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsDel(i *int8) *TaskCreate {
	if i != nil {
		tc.SetIsDel(*i)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TaskCreate) SetUpdateTime(t time.Time) *TaskCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TaskCreate) SetNillableUpdateTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
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
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
				return nil, fmt.Errorf("act: uninitialized hook (forgotten import act/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Task)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TaskMutation", v)
		}
		node = nv
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
	if _, ok := tc.mutation.ClaimTime(); !ok {
		v := task.DefaultClaimTime
		tc.mutation.SetClaimTime(v)
	}
	if _, ok := tc.mutation.IsFinished(); !ok {
		v := task.DefaultIsFinished
		tc.mutation.SetIsFinished(v)
	}
	if _, ok := tc.mutation.Mode(); !ok {
		v := task.DefaultMode
		tc.mutation.SetMode(v)
	}
	if _, ok := tc.mutation.IsDel(); !ok {
		v := task.DefaultIsDel
		tc.mutation.SetIsDel(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := task.DefaultUpdateTime
		tc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if v, ok := tc.mutation.NodeID(); ok {
		if err := task.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`act: validator failed for field "Task.node_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ProcInstID(); !ok {
		return &ValidationError{Name: "proc_inst_id", err: errors.New(`act: missing required field "Task.proc_inst_id"`)}
	}
	if v, ok := tc.mutation.MemberApprover(); ok {
		if err := task.MemberApproverValidator(v); err != nil {
			return &ValidationError{Name: "member_approver", err: fmt.Errorf(`act: validator failed for field "Task.member_approver": %w`, err)}
		}
	}
	if v, ok := tc.mutation.AgreeApprover(); ok {
		if err := task.AgreeApproverValidator(v); err != nil {
			return &ValidationError{Name: "agree_approver", err: fmt.Errorf(`act: validator failed for field "Task.agree_approver": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Mode(); ok {
		if err := task.ModeValidator(v); err != nil {
			return &ValidationError{Name: "mode", err: fmt.Errorf(`act: validator failed for field "Task.mode": %w`, err)}
		}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = id
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.NodeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldNodeID,
		})
		_node.NodeID = value
	}
	if value, ok := tc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldLevel,
		})
		_node.Level = value
	}
	if value, ok := tc.mutation.Step(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldStep,
		})
		_node.Step = value
	}
	if value, ok := tc.mutation.ProcInstID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
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
	if value, ok := tc.mutation.MemberApprover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldMemberApprover,
		})
		_node.MemberApprover = value
	}
	if value, ok := tc.mutation.AgreeApprover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldAgreeApprover,
		})
		_node.AgreeApprover = value
	}
	if value, ok := tc.mutation.IsFinished(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsFinished,
		})
		_node.IsFinished = value
	}
	if value, ok := tc.mutation.Mode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldMode,
		})
		_node.Mode = value
	}
	if value, ok := tc.mutation.DataID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldDataID,
		})
		_node.DataID = value
	}
	if value, ok := tc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
		_node.UpdateTime = value
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
					return nil, fmt.Errorf("unexpected mutation type %T", m)
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
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = id
				}
				mutation.done = true
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
