// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/predicate"
	"act/common/act/task"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetNodeID sets the "node_id" field.
func (tu *TaskUpdate) SetNodeID(s string) *TaskUpdate {
	tu.mutation.SetNodeID(s)
	return tu
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableNodeID(s *string) *TaskUpdate {
	if s != nil {
		tu.SetNodeID(*s)
	}
	return tu
}

// ClearNodeID clears the value of the "node_id" field.
func (tu *TaskUpdate) ClearNodeID() *TaskUpdate {
	tu.mutation.ClearNodeID()
	return tu
}

// SetLevel sets the "level" field.
func (tu *TaskUpdate) SetLevel(i int32) *TaskUpdate {
	tu.mutation.ResetLevel()
	tu.mutation.SetLevel(i)
	return tu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableLevel(i *int32) *TaskUpdate {
	if i != nil {
		tu.SetLevel(*i)
	}
	return tu
}

// AddLevel adds i to the "level" field.
func (tu *TaskUpdate) AddLevel(i int32) *TaskUpdate {
	tu.mutation.AddLevel(i)
	return tu
}

// ClearLevel clears the value of the "level" field.
func (tu *TaskUpdate) ClearLevel() *TaskUpdate {
	tu.mutation.ClearLevel()
	return tu
}

// SetStep sets the "step" field.
func (tu *TaskUpdate) SetStep(i int32) *TaskUpdate {
	tu.mutation.ResetStep()
	tu.mutation.SetStep(i)
	return tu
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStep(i *int32) *TaskUpdate {
	if i != nil {
		tu.SetStep(*i)
	}
	return tu
}

// AddStep adds i to the "step" field.
func (tu *TaskUpdate) AddStep(i int32) *TaskUpdate {
	tu.mutation.AddStep(i)
	return tu
}

// ClearStep clears the value of the "step" field.
func (tu *TaskUpdate) ClearStep() *TaskUpdate {
	tu.mutation.ClearStep()
	return tu
}

// SetProcInstID sets the "proc_inst_id" field.
func (tu *TaskUpdate) SetProcInstID(i int64) *TaskUpdate {
	tu.mutation.ResetProcInstID()
	tu.mutation.SetProcInstID(i)
	return tu
}

// AddProcInstID adds i to the "proc_inst_id" field.
func (tu *TaskUpdate) AddProcInstID(i int64) *TaskUpdate {
	tu.mutation.AddProcInstID(i)
	return tu
}

// SetCreateTime sets the "create_time" field.
func (tu *TaskUpdate) SetCreateTime(t time.Time) *TaskUpdate {
	tu.mutation.SetCreateTime(t)
	return tu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreateTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreateTime(*t)
	}
	return tu
}

// ClearCreateTime clears the value of the "create_time" field.
func (tu *TaskUpdate) ClearCreateTime() *TaskUpdate {
	tu.mutation.ClearCreateTime()
	return tu
}

// SetClaimTime sets the "claim_time" field.
func (tu *TaskUpdate) SetClaimTime(t time.Time) *TaskUpdate {
	tu.mutation.SetClaimTime(t)
	return tu
}

// SetNillableClaimTime sets the "claim_time" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableClaimTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetClaimTime(*t)
	}
	return tu
}

// ClearClaimTime clears the value of the "claim_time" field.
func (tu *TaskUpdate) ClearClaimTime() *TaskUpdate {
	tu.mutation.ClearClaimTime()
	return tu
}

// SetMemberApprover sets the "member_approver" field.
func (tu *TaskUpdate) SetMemberApprover(s string) *TaskUpdate {
	tu.mutation.SetMemberApprover(s)
	return tu
}

// SetNillableMemberApprover sets the "member_approver" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableMemberApprover(s *string) *TaskUpdate {
	if s != nil {
		tu.SetMemberApprover(*s)
	}
	return tu
}

// ClearMemberApprover clears the value of the "member_approver" field.
func (tu *TaskUpdate) ClearMemberApprover() *TaskUpdate {
	tu.mutation.ClearMemberApprover()
	return tu
}

// SetAgreeApprover sets the "agree_approver" field.
func (tu *TaskUpdate) SetAgreeApprover(s string) *TaskUpdate {
	tu.mutation.SetAgreeApprover(s)
	return tu
}

// SetNillableAgreeApprover sets the "agree_approver" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableAgreeApprover(s *string) *TaskUpdate {
	if s != nil {
		tu.SetAgreeApprover(*s)
	}
	return tu
}

// ClearAgreeApprover clears the value of the "agree_approver" field.
func (tu *TaskUpdate) ClearAgreeApprover() *TaskUpdate {
	tu.mutation.ClearAgreeApprover()
	return tu
}

// SetIsFinished sets the "is_finished" field.
func (tu *TaskUpdate) SetIsFinished(i int8) *TaskUpdate {
	tu.mutation.ResetIsFinished()
	tu.mutation.SetIsFinished(i)
	return tu
}

// SetNillableIsFinished sets the "is_finished" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableIsFinished(i *int8) *TaskUpdate {
	if i != nil {
		tu.SetIsFinished(*i)
	}
	return tu
}

// AddIsFinished adds i to the "is_finished" field.
func (tu *TaskUpdate) AddIsFinished(i int8) *TaskUpdate {
	tu.mutation.AddIsFinished(i)
	return tu
}

// ClearIsFinished clears the value of the "is_finished" field.
func (tu *TaskUpdate) ClearIsFinished() *TaskUpdate {
	tu.mutation.ClearIsFinished()
	return tu
}

// SetMode sets the "mode" field.
func (tu *TaskUpdate) SetMode(t task.Mode) *TaskUpdate {
	tu.mutation.SetMode(t)
	return tu
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableMode(t *task.Mode) *TaskUpdate {
	if t != nil {
		tu.SetMode(*t)
	}
	return tu
}

// ClearMode clears the value of the "mode" field.
func (tu *TaskUpdate) ClearMode() *TaskUpdate {
	tu.mutation.ClearMode()
	return tu
}

// SetDataID sets the "data_id" field.
func (tu *TaskUpdate) SetDataID(i int64) *TaskUpdate {
	tu.mutation.ResetDataID()
	tu.mutation.SetDataID(i)
	return tu
}

// SetNillableDataID sets the "data_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDataID(i *int64) *TaskUpdate {
	if i != nil {
		tu.SetDataID(*i)
	}
	return tu
}

// AddDataID adds i to the "data_id" field.
func (tu *TaskUpdate) AddDataID(i int64) *TaskUpdate {
	tu.mutation.AddDataID(i)
	return tu
}

// ClearDataID clears the value of the "data_id" field.
func (tu *TaskUpdate) ClearDataID() *TaskUpdate {
	tu.mutation.ClearDataID()
	return tu
}

// SetIsDel sets the "is_del" field.
func (tu *TaskUpdate) SetIsDel(i int8) *TaskUpdate {
	tu.mutation.ResetIsDel()
	tu.mutation.SetIsDel(i)
	return tu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableIsDel(i *int8) *TaskUpdate {
	if i != nil {
		tu.SetIsDel(*i)
	}
	return tu
}

// AddIsDel adds i to the "is_del" field.
func (tu *TaskUpdate) AddIsDel(i int8) *TaskUpdate {
	tu.mutation.AddIsDel(i)
	return tu
}

// ClearIsDel clears the value of the "is_del" field.
func (tu *TaskUpdate) ClearIsDel() *TaskUpdate {
	tu.mutation.ClearIsDel()
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TaskUpdate) SetUpdateTime(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableUpdateTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetUpdateTime(*t)
	}
	return tu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (tu *TaskUpdate) ClearUpdateTime() *TaskUpdate {
	tu.mutation.ClearUpdateTime()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("act: uninitialized hook (forgotten import act/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.NodeID(); ok {
		if err := task.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`act: validator failed for field "Task.node_id": %w`, err)}
		}
	}
	if v, ok := tu.mutation.MemberApprover(); ok {
		if err := task.MemberApproverValidator(v); err != nil {
			return &ValidationError{Name: "member_approver", err: fmt.Errorf(`act: validator failed for field "Task.member_approver": %w`, err)}
		}
	}
	if v, ok := tu.mutation.AgreeApprover(); ok {
		if err := task.AgreeApproverValidator(v); err != nil {
			return &ValidationError{Name: "agree_approver", err: fmt.Errorf(`act: validator failed for field "Task.agree_approver": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Mode(); ok {
		if err := task.ModeValidator(v); err != nil {
			return &ValidationError{Name: "mode", err: fmt.Errorf(`act: validator failed for field "Task.mode": %w`, err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.NodeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldNodeID,
		})
	}
	if tu.mutation.NodeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldNodeID,
		})
	}
	if value, ok := tu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldLevel,
		})
	}
	if value, ok := tu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldLevel,
		})
	}
	if tu.mutation.LevelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: task.FieldLevel,
		})
	}
	if value, ok := tu.mutation.Step(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldStep,
		})
	}
	if value, ok := tu.mutation.AddedStep(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldStep,
		})
	}
	if tu.mutation.StepCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: task.FieldStep,
		})
	}
	if value, ok := tu.mutation.ProcInstID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldProcInstID,
		})
	}
	if value, ok := tu.mutation.AddedProcInstID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldProcInstID,
		})
	}
	if value, ok := tu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreateTime,
		})
	}
	if tu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldCreateTime,
		})
	}
	if value, ok := tu.mutation.ClaimTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldClaimTime,
		})
	}
	if tu.mutation.ClaimTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldClaimTime,
		})
	}
	if value, ok := tu.mutation.MemberApprover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldMemberApprover,
		})
	}
	if tu.mutation.MemberApproverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldMemberApprover,
		})
	}
	if value, ok := tu.mutation.AgreeApprover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldAgreeApprover,
		})
	}
	if tu.mutation.AgreeApproverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldAgreeApprover,
		})
	}
	if value, ok := tu.mutation.IsFinished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsFinished,
		})
	}
	if value, ok := tu.mutation.AddedIsFinished(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsFinished,
		})
	}
	if tu.mutation.IsFinishedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: task.FieldIsFinished,
		})
	}
	if value, ok := tu.mutation.Mode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldMode,
		})
	}
	if tu.mutation.ModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: task.FieldMode,
		})
	}
	if value, ok := tu.mutation.DataID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldDataID,
		})
	}
	if value, ok := tu.mutation.AddedDataID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldDataID,
		})
	}
	if tu.mutation.DataIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: task.FieldDataID,
		})
	}
	if value, ok := tu.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsDel,
		})
	}
	if value, ok := tu.mutation.AddedIsDel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsDel,
		})
	}
	if tu.mutation.IsDelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: task.FieldIsDel,
		})
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
	}
	if tu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldUpdateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetNodeID sets the "node_id" field.
func (tuo *TaskUpdateOne) SetNodeID(s string) *TaskUpdateOne {
	tuo.mutation.SetNodeID(s)
	return tuo
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableNodeID(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetNodeID(*s)
	}
	return tuo
}

// ClearNodeID clears the value of the "node_id" field.
func (tuo *TaskUpdateOne) ClearNodeID() *TaskUpdateOne {
	tuo.mutation.ClearNodeID()
	return tuo
}

// SetLevel sets the "level" field.
func (tuo *TaskUpdateOne) SetLevel(i int32) *TaskUpdateOne {
	tuo.mutation.ResetLevel()
	tuo.mutation.SetLevel(i)
	return tuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableLevel(i *int32) *TaskUpdateOne {
	if i != nil {
		tuo.SetLevel(*i)
	}
	return tuo
}

// AddLevel adds i to the "level" field.
func (tuo *TaskUpdateOne) AddLevel(i int32) *TaskUpdateOne {
	tuo.mutation.AddLevel(i)
	return tuo
}

// ClearLevel clears the value of the "level" field.
func (tuo *TaskUpdateOne) ClearLevel() *TaskUpdateOne {
	tuo.mutation.ClearLevel()
	return tuo
}

// SetStep sets the "step" field.
func (tuo *TaskUpdateOne) SetStep(i int32) *TaskUpdateOne {
	tuo.mutation.ResetStep()
	tuo.mutation.SetStep(i)
	return tuo
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStep(i *int32) *TaskUpdateOne {
	if i != nil {
		tuo.SetStep(*i)
	}
	return tuo
}

// AddStep adds i to the "step" field.
func (tuo *TaskUpdateOne) AddStep(i int32) *TaskUpdateOne {
	tuo.mutation.AddStep(i)
	return tuo
}

// ClearStep clears the value of the "step" field.
func (tuo *TaskUpdateOne) ClearStep() *TaskUpdateOne {
	tuo.mutation.ClearStep()
	return tuo
}

// SetProcInstID sets the "proc_inst_id" field.
func (tuo *TaskUpdateOne) SetProcInstID(i int64) *TaskUpdateOne {
	tuo.mutation.ResetProcInstID()
	tuo.mutation.SetProcInstID(i)
	return tuo
}

// AddProcInstID adds i to the "proc_inst_id" field.
func (tuo *TaskUpdateOne) AddProcInstID(i int64) *TaskUpdateOne {
	tuo.mutation.AddProcInstID(i)
	return tuo
}

// SetCreateTime sets the "create_time" field.
func (tuo *TaskUpdateOne) SetCreateTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreateTime(t)
	return tuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreateTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreateTime(*t)
	}
	return tuo
}

// ClearCreateTime clears the value of the "create_time" field.
func (tuo *TaskUpdateOne) ClearCreateTime() *TaskUpdateOne {
	tuo.mutation.ClearCreateTime()
	return tuo
}

// SetClaimTime sets the "claim_time" field.
func (tuo *TaskUpdateOne) SetClaimTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetClaimTime(t)
	return tuo
}

// SetNillableClaimTime sets the "claim_time" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableClaimTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetClaimTime(*t)
	}
	return tuo
}

// ClearClaimTime clears the value of the "claim_time" field.
func (tuo *TaskUpdateOne) ClearClaimTime() *TaskUpdateOne {
	tuo.mutation.ClearClaimTime()
	return tuo
}

// SetMemberApprover sets the "member_approver" field.
func (tuo *TaskUpdateOne) SetMemberApprover(s string) *TaskUpdateOne {
	tuo.mutation.SetMemberApprover(s)
	return tuo
}

// SetNillableMemberApprover sets the "member_approver" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableMemberApprover(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetMemberApprover(*s)
	}
	return tuo
}

// ClearMemberApprover clears the value of the "member_approver" field.
func (tuo *TaskUpdateOne) ClearMemberApprover() *TaskUpdateOne {
	tuo.mutation.ClearMemberApprover()
	return tuo
}

// SetAgreeApprover sets the "agree_approver" field.
func (tuo *TaskUpdateOne) SetAgreeApprover(s string) *TaskUpdateOne {
	tuo.mutation.SetAgreeApprover(s)
	return tuo
}

// SetNillableAgreeApprover sets the "agree_approver" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableAgreeApprover(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetAgreeApprover(*s)
	}
	return tuo
}

// ClearAgreeApprover clears the value of the "agree_approver" field.
func (tuo *TaskUpdateOne) ClearAgreeApprover() *TaskUpdateOne {
	tuo.mutation.ClearAgreeApprover()
	return tuo
}

// SetIsFinished sets the "is_finished" field.
func (tuo *TaskUpdateOne) SetIsFinished(i int8) *TaskUpdateOne {
	tuo.mutation.ResetIsFinished()
	tuo.mutation.SetIsFinished(i)
	return tuo
}

// SetNillableIsFinished sets the "is_finished" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableIsFinished(i *int8) *TaskUpdateOne {
	if i != nil {
		tuo.SetIsFinished(*i)
	}
	return tuo
}

// AddIsFinished adds i to the "is_finished" field.
func (tuo *TaskUpdateOne) AddIsFinished(i int8) *TaskUpdateOne {
	tuo.mutation.AddIsFinished(i)
	return tuo
}

// ClearIsFinished clears the value of the "is_finished" field.
func (tuo *TaskUpdateOne) ClearIsFinished() *TaskUpdateOne {
	tuo.mutation.ClearIsFinished()
	return tuo
}

// SetMode sets the "mode" field.
func (tuo *TaskUpdateOne) SetMode(t task.Mode) *TaskUpdateOne {
	tuo.mutation.SetMode(t)
	return tuo
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableMode(t *task.Mode) *TaskUpdateOne {
	if t != nil {
		tuo.SetMode(*t)
	}
	return tuo
}

// ClearMode clears the value of the "mode" field.
func (tuo *TaskUpdateOne) ClearMode() *TaskUpdateOne {
	tuo.mutation.ClearMode()
	return tuo
}

// SetDataID sets the "data_id" field.
func (tuo *TaskUpdateOne) SetDataID(i int64) *TaskUpdateOne {
	tuo.mutation.ResetDataID()
	tuo.mutation.SetDataID(i)
	return tuo
}

// SetNillableDataID sets the "data_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDataID(i *int64) *TaskUpdateOne {
	if i != nil {
		tuo.SetDataID(*i)
	}
	return tuo
}

// AddDataID adds i to the "data_id" field.
func (tuo *TaskUpdateOne) AddDataID(i int64) *TaskUpdateOne {
	tuo.mutation.AddDataID(i)
	return tuo
}

// ClearDataID clears the value of the "data_id" field.
func (tuo *TaskUpdateOne) ClearDataID() *TaskUpdateOne {
	tuo.mutation.ClearDataID()
	return tuo
}

// SetIsDel sets the "is_del" field.
func (tuo *TaskUpdateOne) SetIsDel(i int8) *TaskUpdateOne {
	tuo.mutation.ResetIsDel()
	tuo.mutation.SetIsDel(i)
	return tuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableIsDel(i *int8) *TaskUpdateOne {
	if i != nil {
		tuo.SetIsDel(*i)
	}
	return tuo
}

// AddIsDel adds i to the "is_del" field.
func (tuo *TaskUpdateOne) AddIsDel(i int8) *TaskUpdateOne {
	tuo.mutation.AddIsDel(i)
	return tuo
}

// ClearIsDel clears the value of the "is_del" field.
func (tuo *TaskUpdateOne) ClearIsDel() *TaskUpdateOne {
	tuo.mutation.ClearIsDel()
	return tuo
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TaskUpdateOne) SetUpdateTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableUpdateTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetUpdateTime(*t)
	}
	return tuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (tuo *TaskUpdateOne) ClearUpdateTime() *TaskUpdateOne {
	tuo.mutation.ClearUpdateTime()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("act: uninitialized hook (forgotten import act/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.NodeID(); ok {
		if err := task.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`act: validator failed for field "Task.node_id": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.MemberApprover(); ok {
		if err := task.MemberApproverValidator(v); err != nil {
			return &ValidationError{Name: "member_approver", err: fmt.Errorf(`act: validator failed for field "Task.member_approver": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.AgreeApprover(); ok {
		if err := task.AgreeApproverValidator(v); err != nil {
			return &ValidationError{Name: "agree_approver", err: fmt.Errorf(`act: validator failed for field "Task.agree_approver": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Mode(); ok {
		if err := task.ModeValidator(v); err != nil {
			return &ValidationError{Name: "mode", err: fmt.Errorf(`act: validator failed for field "Task.mode": %w`, err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`act: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("act: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.NodeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldNodeID,
		})
	}
	if tuo.mutation.NodeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldNodeID,
		})
	}
	if value, ok := tuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldLevel,
		})
	}
	if value, ok := tuo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldLevel,
		})
	}
	if tuo.mutation.LevelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: task.FieldLevel,
		})
	}
	if value, ok := tuo.mutation.Step(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldStep,
		})
	}
	if value, ok := tuo.mutation.AddedStep(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: task.FieldStep,
		})
	}
	if tuo.mutation.StepCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: task.FieldStep,
		})
	}
	if value, ok := tuo.mutation.ProcInstID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldProcInstID,
		})
	}
	if value, ok := tuo.mutation.AddedProcInstID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldProcInstID,
		})
	}
	if value, ok := tuo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreateTime,
		})
	}
	if tuo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldCreateTime,
		})
	}
	if value, ok := tuo.mutation.ClaimTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldClaimTime,
		})
	}
	if tuo.mutation.ClaimTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldClaimTime,
		})
	}
	if value, ok := tuo.mutation.MemberApprover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldMemberApprover,
		})
	}
	if tuo.mutation.MemberApproverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldMemberApprover,
		})
	}
	if value, ok := tuo.mutation.AgreeApprover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldAgreeApprover,
		})
	}
	if tuo.mutation.AgreeApproverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldAgreeApprover,
		})
	}
	if value, ok := tuo.mutation.IsFinished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsFinished,
		})
	}
	if value, ok := tuo.mutation.AddedIsFinished(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsFinished,
		})
	}
	if tuo.mutation.IsFinishedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: task.FieldIsFinished,
		})
	}
	if value, ok := tuo.mutation.Mode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldMode,
		})
	}
	if tuo.mutation.ModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: task.FieldMode,
		})
	}
	if value, ok := tuo.mutation.DataID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldDataID,
		})
	}
	if value, ok := tuo.mutation.AddedDataID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: task.FieldDataID,
		})
	}
	if tuo.mutation.DataIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: task.FieldDataID,
		})
	}
	if value, ok := tuo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsDel,
		})
	}
	if value, ok := tuo.mutation.AddedIsDel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldIsDel,
		})
	}
	if tuo.mutation.IsDelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: task.FieldIsDel,
		})
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
	}
	if tuo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldUpdateTime,
		})
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
