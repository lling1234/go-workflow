// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/identitylink"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IdentityLinkCreate is the builder for creating a IdentityLink entity.
type IdentityLinkCreate struct {
	config
	mutation *IdentityLinkMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ilc *IdentityLinkCreate) SetUserID(i int64) *IdentityLinkCreate {
	ilc.mutation.SetUserID(i)
	return ilc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableUserID(i *int64) *IdentityLinkCreate {
	if i != nil {
		ilc.SetUserID(*i)
	}
	return ilc
}

// SetUserName sets the "user_name" field.
func (ilc *IdentityLinkCreate) SetUserName(s string) *IdentityLinkCreate {
	ilc.mutation.SetUserName(s)
	return ilc
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableUserName(s *string) *IdentityLinkCreate {
	if s != nil {
		ilc.SetUserName(*s)
	}
	return ilc
}

// SetStep sets the "step" field.
func (ilc *IdentityLinkCreate) SetStep(i int) *IdentityLinkCreate {
	ilc.mutation.SetStep(i)
	return ilc
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableStep(i *int) *IdentityLinkCreate {
	if i != nil {
		ilc.SetStep(*i)
	}
	return ilc
}

// SetProcInstID sets the "proc_inst_id" field.
func (ilc *IdentityLinkCreate) SetProcInstID(i int64) *IdentityLinkCreate {
	ilc.mutation.SetProcInstID(i)
	return ilc
}

// SetNillableProcInstID sets the "proc_inst_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableProcInstID(i *int64) *IdentityLinkCreate {
	if i != nil {
		ilc.SetProcInstID(*i)
	}
	return ilc
}

// SetTargetID sets the "target_id" field.
func (ilc *IdentityLinkCreate) SetTargetID(i int64) *IdentityLinkCreate {
	ilc.mutation.SetTargetID(i)
	return ilc
}

// SetNillableTargetID sets the "target_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableTargetID(i *int64) *IdentityLinkCreate {
	if i != nil {
		ilc.SetTargetID(*i)
	}
	return ilc
}

// SetComment sets the "comment" field.
func (ilc *IdentityLinkCreate) SetComment(s string) *IdentityLinkCreate {
	ilc.mutation.SetComment(s)
	return ilc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableComment(s *string) *IdentityLinkCreate {
	if s != nil {
		ilc.SetComment(*s)
	}
	return ilc
}

// SetTaskID sets the "task_id" field.
func (ilc *IdentityLinkCreate) SetTaskID(i int64) *IdentityLinkCreate {
	ilc.mutation.SetTaskID(i)
	return ilc
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableTaskID(i *int64) *IdentityLinkCreate {
	if i != nil {
		ilc.SetTaskID(*i)
	}
	return ilc
}

// SetResult sets the "result" field.
func (ilc *IdentityLinkCreate) SetResult(i int) *IdentityLinkCreate {
	ilc.mutation.SetResult(i)
	return ilc
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableResult(i *int) *IdentityLinkCreate {
	if i != nil {
		ilc.SetResult(*i)
	}
	return ilc
}

// SetCreateTime sets the "create_time" field.
func (ilc *IdentityLinkCreate) SetCreateTime(t time.Time) *IdentityLinkCreate {
	ilc.mutation.SetCreateTime(t)
	return ilc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableCreateTime(t *time.Time) *IdentityLinkCreate {
	if t != nil {
		ilc.SetCreateTime(*t)
	}
	return ilc
}

// SetIsDel sets the "is_del" field.
func (ilc *IdentityLinkCreate) SetIsDel(i int8) *IdentityLinkCreate {
	ilc.mutation.SetIsDel(i)
	return ilc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableIsDel(i *int8) *IdentityLinkCreate {
	if i != nil {
		ilc.SetIsDel(*i)
	}
	return ilc
}

// SetIsDeal sets the "is_deal" field.
func (ilc *IdentityLinkCreate) SetIsDeal(i int8) *IdentityLinkCreate {
	ilc.mutation.SetIsDeal(i)
	return ilc
}

// SetNillableIsDeal sets the "is_deal" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableIsDeal(i *int8) *IdentityLinkCreate {
	if i != nil {
		ilc.SetIsDeal(*i)
	}
	return ilc
}

// Mutation returns the IdentityLinkMutation object of the builder.
func (ilc *IdentityLinkCreate) Mutation() *IdentityLinkMutation {
	return ilc.mutation
}

// Save creates the IdentityLink in the database.
func (ilc *IdentityLinkCreate) Save(ctx context.Context) (*IdentityLink, error) {
	var (
		err  error
		node *IdentityLink
	)
	ilc.defaults()
	if len(ilc.hooks) == 0 {
		if err = ilc.check(); err != nil {
			return nil, err
		}
		node, err = ilc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IdentityLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ilc.check(); err != nil {
				return nil, err
			}
			ilc.mutation = mutation
			if node, err = ilc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ilc.hooks) - 1; i >= 0; i-- {
			if ilc.hooks[i] == nil {
				return nil, fmt.Errorf("act: uninitialized hook (forgotten import act/runtime?)")
			}
			mut = ilc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ilc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*IdentityLink)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from IdentityLinkMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ilc *IdentityLinkCreate) SaveX(ctx context.Context) *IdentityLink {
	v, err := ilc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilc *IdentityLinkCreate) Exec(ctx context.Context) error {
	_, err := ilc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilc *IdentityLinkCreate) ExecX(ctx context.Context) {
	if err := ilc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ilc *IdentityLinkCreate) defaults() {
	if _, ok := ilc.mutation.CreateTime(); !ok {
		v := identitylink.DefaultCreateTime
		ilc.mutation.SetCreateTime(v)
	}
	if _, ok := ilc.mutation.IsDel(); !ok {
		v := identitylink.DefaultIsDel
		ilc.mutation.SetIsDel(v)
	}
	if _, ok := ilc.mutation.IsDeal(); !ok {
		v := identitylink.DefaultIsDeal
		ilc.mutation.SetIsDeal(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ilc *IdentityLinkCreate) check() error {
	if v, ok := ilc.mutation.UserName(); ok {
		if err := identitylink.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`act: validator failed for field "IdentityLink.user_name": %w`, err)}
		}
	}
	if v, ok := ilc.mutation.Comment(); ok {
		if err := identitylink.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`act: validator failed for field "IdentityLink.comment": %w`, err)}
		}
	}
	return nil
}

func (ilc *IdentityLinkCreate) sqlSave(ctx context.Context) (*IdentityLink, error) {
	_node, _spec := ilc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ilc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ilc *IdentityLinkCreate) createSpec() (*IdentityLink, *sqlgraph.CreateSpec) {
	var (
		_node = &IdentityLink{config: ilc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: identitylink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: identitylink.FieldID,
			},
		}
	)
	if value, ok := ilc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: identitylink.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ilc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identitylink.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := ilc.mutation.Step(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: identitylink.FieldStep,
		})
		_node.Step = value
	}
	if value, ok := ilc.mutation.ProcInstID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: identitylink.FieldProcInstID,
		})
		_node.ProcInstID = value
	}
	if value, ok := ilc.mutation.TargetID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: identitylink.FieldTargetID,
		})
		_node.TargetID = value
	}
	if value, ok := ilc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identitylink.FieldComment,
		})
		_node.Comment = value
	}
	if value, ok := ilc.mutation.TaskID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: identitylink.FieldTaskID,
		})
		_node.TaskID = value
	}
	if value, ok := ilc.mutation.Result(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: identitylink.FieldResult,
		})
		_node.Result = value
	}
	if value, ok := ilc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: identitylink.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ilc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: identitylink.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := ilc.mutation.IsDeal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: identitylink.FieldIsDeal,
		})
		_node.IsDeal = value
	}
	return _node, _spec
}

// IdentityLinkCreateBulk is the builder for creating many IdentityLink entities in bulk.
type IdentityLinkCreateBulk struct {
	config
	builders []*IdentityLinkCreate
}

// Save creates the IdentityLink entities in the database.
func (ilcb *IdentityLinkCreateBulk) Save(ctx context.Context) ([]*IdentityLink, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ilcb.builders))
	nodes := make([]*IdentityLink, len(ilcb.builders))
	mutators := make([]Mutator, len(ilcb.builders))
	for i := range ilcb.builders {
		func(i int, root context.Context) {
			builder := ilcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentityLinkMutation)
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
					_, err = mutators[i+1].Mutate(root, ilcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ilcb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ilcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ilcb *IdentityLinkCreateBulk) SaveX(ctx context.Context) []*IdentityLink {
	v, err := ilcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilcb *IdentityLinkCreateBulk) Exec(ctx context.Context) error {
	_, err := ilcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilcb *IdentityLinkCreateBulk) ExecX(ctx context.Context) {
	if err := ilcb.Exec(ctx); err != nil {
		panic(err)
	}
}