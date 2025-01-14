// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-wflow/kernel/ent/identitylink"
	"time"

	"github.com/qkbyte/ent/dialect/sql/sqlgraph"
	"github.com/qkbyte/ent/schema/field"
)

// IdentitylinkCreate is the builder for creating a Identitylink entity.
type IdentitylinkCreate struct {
	config
	mutation *IdentitylinkMutation
	hooks    []Hook
}

// SetUserID sets the "userID" field.
func (ic *IdentitylinkCreate) SetUserID(u uint64) *IdentitylinkCreate {
	ic.mutation.SetUserID(u)
	return ic
}

// SetNillableUserID sets the "userID" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableUserID(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetUserID(*u)
	}
	return ic
}

// SetUserName sets the "userName" field.
func (ic *IdentitylinkCreate) SetUserName(s string) *IdentitylinkCreate {
	ic.mutation.SetUserName(s)
	return ic
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableUserName(s *string) *IdentitylinkCreate {
	if s != nil {
		ic.SetUserName(*s)
	}
	return ic
}

// SetProcInstID sets the "procInstID" field.
func (ic *IdentitylinkCreate) SetProcInstID(u uint64) *IdentitylinkCreate {
	ic.mutation.SetProcInstID(u)
	return ic
}

// SetTargetID sets the "targetID" field.
func (ic *IdentitylinkCreate) SetTargetID(u uint64) *IdentitylinkCreate {
	ic.mutation.SetTargetID(u)
	return ic
}

// SetNillableTargetID sets the "targetID" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableTargetID(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetTargetID(*u)
	}
	return ic
}

// SetStation sets the "station" field.
func (ic *IdentitylinkCreate) SetStation(u uint64) *IdentitylinkCreate {
	ic.mutation.SetStation(u)
	return ic
}

// SetNillableStation sets the "station" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableStation(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetStation(*u)
	}
	return ic
}

// SetComment sets the "comment" field.
func (ic *IdentitylinkCreate) SetComment(s string) *IdentitylinkCreate {
	ic.mutation.SetComment(s)
	return ic
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableComment(s *string) *IdentitylinkCreate {
	if s != nil {
		ic.SetComment(*s)
	}
	return ic
}

// SetTaskID sets the "taskID" field.
func (ic *IdentitylinkCreate) SetTaskID(u uint64) *IdentitylinkCreate {
	ic.mutation.SetTaskID(u)
	return ic
}

// SetResult sets the "result" field.
func (ic *IdentitylinkCreate) SetResult(u uint64) *IdentitylinkCreate {
	ic.mutation.SetResult(u)
	return ic
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableResult(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetResult(*u)
	}
	return ic
}

// SetCreateTime sets the "createTime" field.
func (ic *IdentitylinkCreate) SetCreateTime(t time.Time) *IdentitylinkCreate {
	ic.mutation.SetCreateTime(t)
	return ic
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableCreateTime(t *time.Time) *IdentitylinkCreate {
	if t != nil {
		ic.SetCreateTime(*t)
	}
	return ic
}

// SetIsDeal sets the "isDeal" field.
func (ic *IdentitylinkCreate) SetIsDeal(u uint64) *IdentitylinkCreate {
	ic.mutation.SetIsDeal(u)
	return ic
}

// SetNillableIsDeal sets the "isDeal" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableIsDeal(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetIsDeal(*u)
	}
	return ic
}

// SetUpdateTime sets the "updateTime" field.
func (ic *IdentitylinkCreate) SetUpdateTime(t time.Time) *IdentitylinkCreate {
	ic.mutation.SetUpdateTime(t)
	return ic
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableUpdateTime(t *time.Time) *IdentitylinkCreate {
	if t != nil {
		ic.SetUpdateTime(*t)
	}
	return ic
}

// SetCreateUser sets the "createUser" field.
func (ic *IdentitylinkCreate) SetCreateUser(u uint64) *IdentitylinkCreate {
	ic.mutation.SetCreateUser(u)
	return ic
}

// SetNillableCreateUser sets the "createUser" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableCreateUser(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetCreateUser(*u)
	}
	return ic
}

// SetUpdateUser sets the "updateUser" field.
func (ic *IdentitylinkCreate) SetUpdateUser(u uint64) *IdentitylinkCreate {
	ic.mutation.SetUpdateUser(u)
	return ic
}

// SetNillableUpdateUser sets the "updateUser" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableUpdateUser(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetUpdateUser(*u)
	}
	return ic
}

// SetAttachmentID sets the "attachmentID" field.
func (ic *IdentitylinkCreate) SetAttachmentID(u uint64) *IdentitylinkCreate {
	ic.mutation.SetAttachmentID(u)
	return ic
}

// SetNillableAttachmentID sets the "attachmentID" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableAttachmentID(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetAttachmentID(*u)
	}
	return ic
}

// SetVersion sets the "version" field.
func (ic *IdentitylinkCreate) SetVersion(u uint64) *IdentitylinkCreate {
	ic.mutation.SetVersion(u)
	return ic
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableVersion(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetVersion(*u)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IdentitylinkCreate) SetID(u uint64) *IdentitylinkCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *IdentitylinkCreate) SetNillableID(u *uint64) *IdentitylinkCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// Mutation returns the IdentitylinkMutation object of the builder.
func (ic *IdentitylinkCreate) Mutation() *IdentitylinkMutation {
	return ic.mutation
}

// Save creates the Identitylink in the database.
func (ic *IdentitylinkCreate) Save(ctx context.Context) (*Identitylink, error) {
	var (
		err  error
		node *Identitylink
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IdentitylinkMutation)
			if !ok {
				return nil, fmt.Errorf("意外突变类型 %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: 未初始化挂钩 (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IdentitylinkCreate) SaveX(ctx context.Context) *Identitylink {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IdentitylinkCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IdentitylinkCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IdentitylinkCreate) defaults() {
	if _, ok := ic.mutation.CreateTime(); !ok {
		v := identitylink.DefaultCreateTime
		ic.mutation.SetCreateTime(v)
	}
	if _, ok := ic.mutation.IsDeal(); !ok {
		v := identitylink.DefaultIsDeal
		ic.mutation.SetIsDeal(v)
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		v := identitylink.DefaultUpdateTime
		ic.mutation.SetUpdateTime(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := identitylink.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IdentitylinkCreate) check() error {
	if v, ok := ic.mutation.UserName(); ok {
		if err := identitylink.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "userName", err: fmt.Errorf(`ent: 字段验证失败 "Identitylink.userName": %w`, err)}
		}
	}
	if _, ok := ic.mutation.ProcInstID(); !ok {
		return &ValidationError{Name: "procInstID", err: errors.New(`ent: missing required field "Identitylink.procInstID"`)}
	}
	if v, ok := ic.mutation.Comment(); ok {
		if err := identitylink.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: 字段验证失败 "Identitylink.comment": %w`, err)}
		}
	}
	if _, ok := ic.mutation.TaskID(); !ok {
		return &ValidationError{Name: "taskID", err: errors.New(`ent: missing required field "Identitylink.taskID"`)}
	}
	return nil
}

func (ic *IdentitylinkCreate) sqlSave(ctx context.Context) (*Identitylink, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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

func (ic *IdentitylinkCreate) createSpec() (*Identitylink, *sqlgraph.CreateSpec) {
	var (
		_node = &Identitylink{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: identitylink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: identitylink.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ic.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identitylink.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := ic.mutation.ProcInstID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldProcInstID,
		})
		_node.ProcInstID = value
	}
	if value, ok := ic.mutation.TargetID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldTargetID,
		})
		_node.TargetID = value
	}
	if value, ok := ic.mutation.Station(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldStation,
		})
		_node.Station = value
	}
	if value, ok := ic.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identitylink.FieldComment,
		})
		_node.Comment = value
	}
	if value, ok := ic.mutation.TaskID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldTaskID,
		})
		_node.TaskID = value
	}
	if value, ok := ic.mutation.Result(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldResult,
		})
		_node.Result = value
	}
	if value, ok := ic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: identitylink.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ic.mutation.IsDeal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldIsDeal,
		})
		_node.IsDeal = value
	}
	if value, ok := ic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: identitylink.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ic.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := ic.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := ic.mutation.AttachmentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldAttachmentID,
		})
		_node.AttachmentID = value
	}
	if value, ok := ic.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: identitylink.FieldVersion,
		})
		_node.Version = value
	}
	return _node, _spec
}

// IdentitylinkCreateBulk is the builder for creating many Identitylink entities in bulk.
type IdentitylinkCreateBulk struct {
	config
	builders []*IdentitylinkCreate
}

// Save creates the Identitylink entities in the database.
func (icb *IdentitylinkCreateBulk) Save(ctx context.Context) ([]*Identitylink, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Identitylink, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentitylinkMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IdentitylinkCreateBulk) SaveX(ctx context.Context) []*Identitylink {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IdentitylinkCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IdentitylinkCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
