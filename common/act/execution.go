// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/execution"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Execution is the model entity for the Execution schema.
type Execution struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 流程实例id
	ProcInstID int64 `json:"proc_inst_id,omitempty"`
	// 流程定义id
	ProcDefID int64 `json:"proc_def_id,omitempty"`
	// 节点审批执行顺序
	NodeInfos string `json:"node_infos,omitempty"`
	// 是否审批中
	IsActive int8 `json:"is_active,omitempty"`
	// 开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 是否删除,0:未删除,1:已删除
	IsDel int8 `json:"is_del,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Execution) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case execution.FieldID, execution.FieldProcInstID, execution.FieldProcDefID, execution.FieldIsActive, execution.FieldIsDel:
			values[i] = new(sql.NullInt64)
		case execution.FieldNodeInfos:
			values[i] = new(sql.NullString)
		case execution.FieldStartTime, execution.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Execution", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Execution fields.
func (e *Execution) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case execution.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case execution.FieldProcInstID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proc_inst_id", values[i])
			} else if value.Valid {
				e.ProcInstID = value.Int64
			}
		case execution.FieldProcDefID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proc_def_id", values[i])
			} else if value.Valid {
				e.ProcDefID = value.Int64
			}
		case execution.FieldNodeInfos:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_infos", values[i])
			} else if value.Valid {
				e.NodeInfos = value.String
			}
		case execution.FieldIsActive:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				e.IsActive = int8(value.Int64)
			}
		case execution.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				e.StartTime = value.Time
			}
		case execution.FieldIsDel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				e.IsDel = int8(value.Int64)
			}
		case execution.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Execution.
// Note that you need to call Execution.Unwrap() before calling this method if this Execution
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Execution) Update() *ExecutionUpdateOne {
	return (&ExecutionClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Execution entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Execution) Unwrap() *Execution {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("act: Execution is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Execution) String() string {
	var builder strings.Builder
	builder.WriteString("Execution(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("proc_inst_id=")
	builder.WriteString(fmt.Sprintf("%v", e.ProcInstID))
	builder.WriteString(", ")
	builder.WriteString("proc_def_id=")
	builder.WriteString(fmt.Sprintf("%v", e.ProcDefID))
	builder.WriteString(", ")
	builder.WriteString("node_infos=")
	builder.WriteString(e.NodeInfos)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", e.IsActive))
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(e.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", e.IsDel))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Executions is a parsable slice of Execution.
type Executions []*Execution

func (e Executions) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}