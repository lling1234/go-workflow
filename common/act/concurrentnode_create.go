// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/concurrentnode"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConcurrentNodeCreate is the builder for creating a ConcurrentNode entity.
type ConcurrentNodeCreate struct {
	config
	mutation *ConcurrentNodeMutation
	hooks    []Hook
}

// SetProcInstID sets the "proc_inst_id" field.
func (cnc *ConcurrentNodeCreate) SetProcInstID(i int64) *ConcurrentNodeCreate {
	cnc.mutation.SetProcInstID(i)
	return cnc
}

// SetProcDefID sets the "proc_def_id" field.
func (cnc *ConcurrentNodeCreate) SetProcDefID(i int64) *ConcurrentNodeCreate {
	cnc.mutation.SetProcDefID(i)
	return cnc
}

// SetNodeID sets the "node_id" field.
func (cnc *ConcurrentNodeCreate) SetNodeID(s string) *ConcurrentNodeCreate {
	cnc.mutation.SetNodeID(s)
	return cnc
}

// SetNodeInfo sets the "node_info" field.
func (cnc *ConcurrentNodeCreate) SetNodeInfo(s string) *ConcurrentNodeCreate {
	cnc.mutation.SetNodeInfo(s)
	return cnc
}

// SetNillableNodeInfo sets the "node_info" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillableNodeInfo(s *string) *ConcurrentNodeCreate {
	if s != nil {
		cnc.SetNodeInfo(*s)
	}
	return cnc
}

// SetPrevID sets the "prev_id" field.
func (cnc *ConcurrentNodeCreate) SetPrevID(s string) *ConcurrentNodeCreate {
	cnc.mutation.SetPrevID(s)
	return cnc
}

// SetNillablePrevID sets the "prev_id" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillablePrevID(s *string) *ConcurrentNodeCreate {
	if s != nil {
		cnc.SetPrevID(*s)
	}
	return cnc
}

// SetNextID sets the "next_id" field.
func (cnc *ConcurrentNodeCreate) SetNextID(s string) *ConcurrentNodeCreate {
	cnc.mutation.SetNextID(s)
	return cnc
}

// SetNillableNextID sets the "next_id" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillableNextID(s *string) *ConcurrentNodeCreate {
	if s != nil {
		cnc.SetNextID(*s)
	}
	return cnc
}

// SetState sets the "state" field.
func (cnc *ConcurrentNodeCreate) SetState(i int32) *ConcurrentNodeCreate {
	cnc.mutation.SetState(i)
	return cnc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillableState(i *int32) *ConcurrentNodeCreate {
	if i != nil {
		cnc.SetState(*i)
	}
	return cnc
}

// SetIsDel sets the "is_del" field.
func (cnc *ConcurrentNodeCreate) SetIsDel(i int32) *ConcurrentNodeCreate {
	cnc.mutation.SetIsDel(i)
	return cnc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillableIsDel(i *int32) *ConcurrentNodeCreate {
	if i != nil {
		cnc.SetIsDel(*i)
	}
	return cnc
}

// SetCreateTime sets the "create_time" field.
func (cnc *ConcurrentNodeCreate) SetCreateTime(t time.Time) *ConcurrentNodeCreate {
	cnc.mutation.SetCreateTime(t)
	return cnc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillableCreateTime(t *time.Time) *ConcurrentNodeCreate {
	if t != nil {
		cnc.SetCreateTime(*t)
	}
	return cnc
}

// SetUpdateTime sets the "update_time" field.
func (cnc *ConcurrentNodeCreate) SetUpdateTime(t time.Time) *ConcurrentNodeCreate {
	cnc.mutation.SetUpdateTime(t)
	return cnc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cnc *ConcurrentNodeCreate) SetNillableUpdateTime(t *time.Time) *ConcurrentNodeCreate {
	if t != nil {
		cnc.SetUpdateTime(*t)
	}
	return cnc
}

// SetID sets the "id" field.
func (cnc *ConcurrentNodeCreate) SetID(i int64) *ConcurrentNodeCreate {
	cnc.mutation.SetID(i)
	return cnc
}

// Mutation returns the ConcurrentNodeMutation object of the builder.
func (cnc *ConcurrentNodeCreate) Mutation() *ConcurrentNodeMutation {
	return cnc.mutation
}

// Save creates the ConcurrentNode in the database.
func (cnc *ConcurrentNodeCreate) Save(ctx context.Context) (*ConcurrentNode, error) {
	var (
		err  error
		node *ConcurrentNode
	)
	cnc.defaults()
	if len(cnc.hooks) == 0 {
		if err = cnc.check(); err != nil {
			return nil, err
		}
		node, err = cnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConcurrentNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cnc.check(); err != nil {
				return nil, err
			}
			cnc.mutation = mutation
			if node, err = cnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cnc.hooks) - 1; i >= 0; i-- {
			if cnc.hooks[i] == nil {
				return nil, fmt.Errorf("act: uninitialized hook (forgotten import act/runtime?)")
			}
			mut = cnc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cnc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ConcurrentNode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ConcurrentNodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cnc *ConcurrentNodeCreate) SaveX(ctx context.Context) *ConcurrentNode {
	v, err := cnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cnc *ConcurrentNodeCreate) Exec(ctx context.Context) error {
	_, err := cnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnc *ConcurrentNodeCreate) ExecX(ctx context.Context) {
	if err := cnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cnc *ConcurrentNodeCreate) defaults() {
	if _, ok := cnc.mutation.IsDel(); !ok {
		v := concurrentnode.DefaultIsDel
		cnc.mutation.SetIsDel(v)
	}
	if _, ok := cnc.mutation.CreateTime(); !ok {
		v := concurrentnode.DefaultCreateTime
		cnc.mutation.SetCreateTime(v)
	}
	if _, ok := cnc.mutation.UpdateTime(); !ok {
		v := concurrentnode.DefaultUpdateTime
		cnc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cnc *ConcurrentNodeCreate) check() error {
	if _, ok := cnc.mutation.ProcInstID(); !ok {
		return &ValidationError{Name: "proc_inst_id", err: errors.New(`act: missing required field "ConcurrentNode.proc_inst_id"`)}
	}
	if _, ok := cnc.mutation.ProcDefID(); !ok {
		return &ValidationError{Name: "proc_def_id", err: errors.New(`act: missing required field "ConcurrentNode.proc_def_id"`)}
	}
	if _, ok := cnc.mutation.NodeID(); !ok {
		return &ValidationError{Name: "node_id", err: errors.New(`act: missing required field "ConcurrentNode.node_id"`)}
	}
	if v, ok := cnc.mutation.NodeID(); ok {
		if err := concurrentnode.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`act: validator failed for field "ConcurrentNode.node_id": %w`, err)}
		}
	}
	if v, ok := cnc.mutation.NodeInfo(); ok {
		if err := concurrentnode.NodeInfoValidator(v); err != nil {
			return &ValidationError{Name: "node_info", err: fmt.Errorf(`act: validator failed for field "ConcurrentNode.node_info": %w`, err)}
		}
	}
	if v, ok := cnc.mutation.PrevID(); ok {
		if err := concurrentnode.PrevIDValidator(v); err != nil {
			return &ValidationError{Name: "prev_id", err: fmt.Errorf(`act: validator failed for field "ConcurrentNode.prev_id": %w`, err)}
		}
	}
	if v, ok := cnc.mutation.NextID(); ok {
		if err := concurrentnode.NextIDValidator(v); err != nil {
			return &ValidationError{Name: "next_id", err: fmt.Errorf(`act: validator failed for field "ConcurrentNode.next_id": %w`, err)}
		}
	}
	return nil
}

func (cnc *ConcurrentNodeCreate) sqlSave(ctx context.Context) (*ConcurrentNode, error) {
	_node, _spec := cnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (cnc *ConcurrentNodeCreate) createSpec() (*ConcurrentNode, *sqlgraph.CreateSpec) {
	var (
		_node = &ConcurrentNode{config: cnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: concurrentnode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: concurrentnode.FieldID,
			},
		}
	)
	if id, ok := cnc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cnc.mutation.ProcInstID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: concurrentnode.FieldProcInstID,
		})
		_node.ProcInstID = value
	}
	if value, ok := cnc.mutation.ProcDefID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: concurrentnode.FieldProcDefID,
		})
		_node.ProcDefID = value
	}
	if value, ok := cnc.mutation.NodeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: concurrentnode.FieldNodeID,
		})
		_node.NodeID = value
	}
	if value, ok := cnc.mutation.NodeInfo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: concurrentnode.FieldNodeInfo,
		})
		_node.NodeInfo = value
	}
	if value, ok := cnc.mutation.PrevID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: concurrentnode.FieldPrevID,
		})
		_node.PrevID = value
	}
	if value, ok := cnc.mutation.NextID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: concurrentnode.FieldNextID,
		})
		_node.NextID = value
	}
	if value, ok := cnc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: concurrentnode.FieldState,
		})
		_node.State = value
	}
	if value, ok := cnc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: concurrentnode.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := cnc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: concurrentnode.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cnc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: concurrentnode.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	return _node, _spec
}

// ConcurrentNodeCreateBulk is the builder for creating many ConcurrentNode entities in bulk.
type ConcurrentNodeCreateBulk struct {
	config
	builders []*ConcurrentNodeCreate
}

// Save creates the ConcurrentNode entities in the database.
func (cncb *ConcurrentNodeCreateBulk) Save(ctx context.Context) ([]*ConcurrentNode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cncb.builders))
	nodes := make([]*ConcurrentNode, len(cncb.builders))
	mutators := make([]Mutator, len(cncb.builders))
	for i := range cncb.builders {
		func(i int, root context.Context) {
			builder := cncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConcurrentNodeMutation)
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
					_, err = mutators[i+1].Mutate(root, cncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, cncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cncb *ConcurrentNodeCreateBulk) SaveX(ctx context.Context) []*ConcurrentNode {
	v, err := cncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cncb *ConcurrentNodeCreateBulk) Exec(ctx context.Context) error {
	_, err := cncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cncb *ConcurrentNodeCreateBulk) ExecX(ctx context.Context) {
	if err := cncb.Exec(ctx); err != nil {
		panic(err)
	}
}
