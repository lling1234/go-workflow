// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/procdef"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProcDefCreate is the builder for creating a ProcDef entity.
type ProcDefCreate struct {
	config
	mutation *ProcDefMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pdc *ProcDefCreate) SetName(s string) *ProcDefCreate {
	pdc.mutation.SetName(s)
	return pdc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableName(s *string) *ProcDefCreate {
	if s != nil {
		pdc.SetName(*s)
	}
	return pdc
}

// SetCode sets the "code" field.
func (pdc *ProcDefCreate) SetCode(s string) *ProcDefCreate {
	pdc.mutation.SetCode(s)
	return pdc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableCode(s *string) *ProcDefCreate {
	if s != nil {
		pdc.SetCode(*s)
	}
	return pdc
}

// SetVersion sets the "version" field.
func (pdc *ProcDefCreate) SetVersion(i int) *ProcDefCreate {
	pdc.mutation.SetVersion(i)
	return pdc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableVersion(i *int) *ProcDefCreate {
	if i != nil {
		pdc.SetVersion(*i)
	}
	return pdc
}

// SetResource sets the "resource" field.
func (pdc *ProcDefCreate) SetResource(s string) *ProcDefCreate {
	pdc.mutation.SetResource(s)
	return pdc
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableResource(s *string) *ProcDefCreate {
	if s != nil {
		pdc.SetResource(*s)
	}
	return pdc
}

// SetCreateUserID sets the "create_user_id" field.
func (pdc *ProcDefCreate) SetCreateUserID(i int64) *ProcDefCreate {
	pdc.mutation.SetCreateUserID(i)
	return pdc
}

// SetNillableCreateUserID sets the "create_user_id" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableCreateUserID(i *int64) *ProcDefCreate {
	if i != nil {
		pdc.SetCreateUserID(*i)
	}
	return pdc
}

// SetCreateUserName sets the "create_user_name" field.
func (pdc *ProcDefCreate) SetCreateUserName(s string) *ProcDefCreate {
	pdc.mutation.SetCreateUserName(s)
	return pdc
}

// SetNillableCreateUserName sets the "create_user_name" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableCreateUserName(s *string) *ProcDefCreate {
	if s != nil {
		pdc.SetCreateUserName(*s)
	}
	return pdc
}

// SetCreateTime sets the "create_time" field.
func (pdc *ProcDefCreate) SetCreateTime(t time.Time) *ProcDefCreate {
	pdc.mutation.SetCreateTime(t)
	return pdc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableCreateTime(t *time.Time) *ProcDefCreate {
	if t != nil {
		pdc.SetCreateTime(*t)
	}
	return pdc
}

// SetTargetID sets the "target_id" field.
func (pdc *ProcDefCreate) SetTargetID(i int64) *ProcDefCreate {
	pdc.mutation.SetTargetID(i)
	return pdc
}

// SetNillableTargetID sets the "target_id" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableTargetID(i *int64) *ProcDefCreate {
	if i != nil {
		pdc.SetTargetID(*i)
	}
	return pdc
}

// SetFormID sets the "form_id" field.
func (pdc *ProcDefCreate) SetFormID(s string) *ProcDefCreate {
	pdc.mutation.SetFormID(s)
	return pdc
}

// SetNillableFormID sets the "form_id" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableFormID(s *string) *ProcDefCreate {
	if s != nil {
		pdc.SetFormID(*s)
	}
	return pdc
}

// SetFormName sets the "form_name" field.
func (pdc *ProcDefCreate) SetFormName(s string) *ProcDefCreate {
	pdc.mutation.SetFormName(s)
	return pdc
}

// SetNillableFormName sets the "form_name" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableFormName(s *string) *ProcDefCreate {
	if s != nil {
		pdc.SetFormName(*s)
	}
	return pdc
}

// SetRemainHours sets the "remain_hours" field.
func (pdc *ProcDefCreate) SetRemainHours(i int) *ProcDefCreate {
	pdc.mutation.SetRemainHours(i)
	return pdc
}

// SetNillableRemainHours sets the "remain_hours" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableRemainHours(i *int) *ProcDefCreate {
	if i != nil {
		pdc.SetRemainHours(*i)
	}
	return pdc
}

// SetIsDel sets the "is_del" field.
func (pdc *ProcDefCreate) SetIsDel(i int8) *ProcDefCreate {
	pdc.mutation.SetIsDel(i)
	return pdc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableIsDel(i *int8) *ProcDefCreate {
	if i != nil {
		pdc.SetIsDel(*i)
	}
	return pdc
}

// SetIsActive sets the "is_active" field.
func (pdc *ProcDefCreate) SetIsActive(i int8) *ProcDefCreate {
	pdc.mutation.SetIsActive(i)
	return pdc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableIsActive(i *int8) *ProcDefCreate {
	if i != nil {
		pdc.SetIsActive(*i)
	}
	return pdc
}

// SetUpdateTime sets the "update_time" field.
func (pdc *ProcDefCreate) SetUpdateTime(t time.Time) *ProcDefCreate {
	pdc.mutation.SetUpdateTime(t)
	return pdc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pdc *ProcDefCreate) SetNillableUpdateTime(t *time.Time) *ProcDefCreate {
	if t != nil {
		pdc.SetUpdateTime(*t)
	}
	return pdc
}

// Mutation returns the ProcDefMutation object of the builder.
func (pdc *ProcDefCreate) Mutation() *ProcDefMutation {
	return pdc.mutation
}

// Save creates the ProcDef in the database.
func (pdc *ProcDefCreate) Save(ctx context.Context) (*ProcDef, error) {
	var (
		err  error
		node *ProcDef
	)
	pdc.defaults()
	if len(pdc.hooks) == 0 {
		if err = pdc.check(); err != nil {
			return nil, err
		}
		node, err = pdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProcDefMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pdc.check(); err != nil {
				return nil, err
			}
			pdc.mutation = mutation
			if node, err = pdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pdc.hooks) - 1; i >= 0; i-- {
			if pdc.hooks[i] == nil {
				return nil, fmt.Errorf("act: uninitialized hook (forgotten import act/runtime?)")
			}
			mut = pdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProcDef)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProcDefMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pdc *ProcDefCreate) SaveX(ctx context.Context) *ProcDef {
	v, err := pdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdc *ProcDefCreate) Exec(ctx context.Context) error {
	_, err := pdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdc *ProcDefCreate) ExecX(ctx context.Context) {
	if err := pdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdc *ProcDefCreate) defaults() {
	if _, ok := pdc.mutation.Version(); !ok {
		v := procdef.DefaultVersion
		pdc.mutation.SetVersion(v)
	}
	if _, ok := pdc.mutation.CreateTime(); !ok {
		v := procdef.DefaultCreateTime
		pdc.mutation.SetCreateTime(v)
	}
	if _, ok := pdc.mutation.RemainHours(); !ok {
		v := procdef.DefaultRemainHours
		pdc.mutation.SetRemainHours(v)
	}
	if _, ok := pdc.mutation.IsDel(); !ok {
		v := procdef.DefaultIsDel
		pdc.mutation.SetIsDel(v)
	}
	if _, ok := pdc.mutation.IsActive(); !ok {
		v := procdef.DefaultIsActive
		pdc.mutation.SetIsActive(v)
	}
	if _, ok := pdc.mutation.UpdateTime(); !ok {
		v := procdef.DefaultUpdateTime
		pdc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdc *ProcDefCreate) check() error {
	if v, ok := pdc.mutation.Name(); ok {
		if err := procdef.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`act: validator failed for field "ProcDef.name": %w`, err)}
		}
	}
	if v, ok := pdc.mutation.Code(); ok {
		if err := procdef.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`act: validator failed for field "ProcDef.code": %w`, err)}
		}
	}
	if v, ok := pdc.mutation.Resource(); ok {
		if err := procdef.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf(`act: validator failed for field "ProcDef.resource": %w`, err)}
		}
	}
	if v, ok := pdc.mutation.CreateUserName(); ok {
		if err := procdef.CreateUserNameValidator(v); err != nil {
			return &ValidationError{Name: "create_user_name", err: fmt.Errorf(`act: validator failed for field "ProcDef.create_user_name": %w`, err)}
		}
	}
	if v, ok := pdc.mutation.FormID(); ok {
		if err := procdef.FormIDValidator(v); err != nil {
			return &ValidationError{Name: "form_id", err: fmt.Errorf(`act: validator failed for field "ProcDef.form_id": %w`, err)}
		}
	}
	if v, ok := pdc.mutation.FormName(); ok {
		if err := procdef.FormNameValidator(v); err != nil {
			return &ValidationError{Name: "form_name", err: fmt.Errorf(`act: validator failed for field "ProcDef.form_name": %w`, err)}
		}
	}
	return nil
}

func (pdc *ProcDefCreate) sqlSave(ctx context.Context) (*ProcDef, error) {
	_node, _spec := pdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pdc *ProcDefCreate) createSpec() (*ProcDef, *sqlgraph.CreateSpec) {
	var (
		_node = &ProcDef{config: pdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: procdef.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: procdef.FieldID,
			},
		}
	)
	if value, ok := pdc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procdef.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pdc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procdef.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := pdc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: procdef.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := pdc.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procdef.FieldResource,
		})
		_node.Resource = value
	}
	if value, ok := pdc.mutation.CreateUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: procdef.FieldCreateUserID,
		})
		_node.CreateUserID = value
	}
	if value, ok := pdc.mutation.CreateUserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procdef.FieldCreateUserName,
		})
		_node.CreateUserName = value
	}
	if value, ok := pdc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: procdef.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pdc.mutation.TargetID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: procdef.FieldTargetID,
		})
		_node.TargetID = value
	}
	if value, ok := pdc.mutation.FormID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procdef.FieldFormID,
		})
		_node.FormID = value
	}
	if value, ok := pdc.mutation.FormName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procdef.FieldFormName,
		})
		_node.FormName = value
	}
	if value, ok := pdc.mutation.RemainHours(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: procdef.FieldRemainHours,
		})
		_node.RemainHours = value
	}
	if value, ok := pdc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: procdef.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := pdc.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: procdef.FieldIsActive,
		})
		_node.IsActive = value
	}
	if value, ok := pdc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: procdef.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	return _node, _spec
}

// ProcDefCreateBulk is the builder for creating many ProcDef entities in bulk.
type ProcDefCreateBulk struct {
	config
	builders []*ProcDefCreate
}

// Save creates the ProcDef entities in the database.
func (pdcb *ProcDefCreateBulk) Save(ctx context.Context) ([]*ProcDef, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pdcb.builders))
	nodes := make([]*ProcDef, len(pdcb.builders))
	mutators := make([]Mutator, len(pdcb.builders))
	for i := range pdcb.builders {
		func(i int, root context.Context) {
			builder := pdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProcDefMutation)
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
					_, err = mutators[i+1].Mutate(root, pdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pdcb *ProcDefCreateBulk) SaveX(ctx context.Context) []*ProcDef {
	v, err := pdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdcb *ProcDefCreateBulk) Exec(ctx context.Context) error {
	_, err := pdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdcb *ProcDefCreateBulk) ExecX(ctx context.Context) {
	if err := pdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
