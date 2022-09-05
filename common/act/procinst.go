// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/procinst"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ProcInst is the model entity for the ProcInst schema.
type ProcInst struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 流程定义id
	ProcDefID int64 `json:"proc_def_id,omitempty"`
	// 发起流程标题
	Title string `json:"title,omitempty"`
	// 流程编号
	Code string `json:"code,omitempty"`
	// 组织id
	TargetID int64 `json:"target_id,omitempty"`
	// 节点id
	NodeID string `json:"node_id,omitempty"`
	// 当前审批任务id
	TaskID int64 `json:"task_id,omitempty"`
	// 发起时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 结束时间
	EndTime time.Time `json:"end_time,omitempty"`
	// 发起人id
	StartUserID int64 `json:"start_user_id,omitempty"`
	// 发起人姓名
	StartUserName string `json:"start_user_name,omitempty"`
	// 流程是否结束,0:未结束,1:已结束
	IsFinished int8 `json:"is_finished,omitempty"`
	// 流程状态,类型为:1待处理、2处理中、3驳回、4已撤回、5未通过、6已通过、7废弃
	State int32 `json:"state,omitempty"`
	// 流程绑定数据id
	DataID int64 `json:"data_id,omitempty"`
	// 是否删除
	IsDel int8 `json:"is_del,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 审批限定时间
	RemainHours int32 `json:"remain_hours,omitempty"`
	// 流程修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProcInst) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case procinst.FieldID, procinst.FieldProcDefID, procinst.FieldTargetID, procinst.FieldTaskID, procinst.FieldStartUserID, procinst.FieldIsFinished, procinst.FieldState, procinst.FieldDataID, procinst.FieldIsDel, procinst.FieldRemainHours:
			values[i] = new(sql.NullInt64)
		case procinst.FieldTitle, procinst.FieldCode, procinst.FieldNodeID, procinst.FieldStartUserName:
			values[i] = new(sql.NullString)
		case procinst.FieldStartTime, procinst.FieldEndTime, procinst.FieldCreateTime, procinst.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProcInst", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProcInst fields.
func (pi *ProcInst) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case procinst.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = value.Int64
		case procinst.FieldProcDefID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proc_def_id", values[i])
			} else if value.Valid {
				pi.ProcDefID = value.Int64
			}
		case procinst.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pi.Title = value.String
			}
		case procinst.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				pi.Code = value.String
			}
		case procinst.FieldTargetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field target_id", values[i])
			} else if value.Valid {
				pi.TargetID = value.Int64
			}
		case procinst.FieldNodeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				pi.NodeID = value.String
			}
		case procinst.FieldTaskID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value.Valid {
				pi.TaskID = value.Int64
			}
		case procinst.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				pi.StartTime = value.Time
			}
		case procinst.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				pi.EndTime = value.Time
			}
		case procinst.FieldStartUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_user_id", values[i])
			} else if value.Valid {
				pi.StartUserID = value.Int64
			}
		case procinst.FieldStartUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_user_name", values[i])
			} else if value.Valid {
				pi.StartUserName = value.String
			}
		case procinst.FieldIsFinished:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_finished", values[i])
			} else if value.Valid {
				pi.IsFinished = int8(value.Int64)
			}
		case procinst.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				pi.State = int32(value.Int64)
			}
		case procinst.FieldDataID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field data_id", values[i])
			} else if value.Valid {
				pi.DataID = value.Int64
			}
		case procinst.FieldIsDel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				pi.IsDel = int8(value.Int64)
			}
		case procinst.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pi.CreateTime = value.Time
			}
		case procinst.FieldRemainHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remain_hours", values[i])
			} else if value.Valid {
				pi.RemainHours = int32(value.Int64)
			}
		case procinst.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pi.UpdateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ProcInst.
// Note that you need to call ProcInst.Unwrap() before calling this method if this ProcInst
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *ProcInst) Update() *ProcInstUpdateOne {
	return (&ProcInstClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the ProcInst entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *ProcInst) Unwrap() *ProcInst {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("act: ProcInst is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *ProcInst) String() string {
	var builder strings.Builder
	builder.WriteString("ProcInst(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("proc_def_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.ProcDefID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pi.Title)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(pi.Code)
	builder.WriteString(", ")
	builder.WriteString("target_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.TargetID))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(pi.NodeID)
	builder.WriteString(", ")
	builder.WriteString("task_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.TaskID))
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(pi.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(pi.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("start_user_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.StartUserID))
	builder.WriteString(", ")
	builder.WriteString("start_user_name=")
	builder.WriteString(pi.StartUserName)
	builder.WriteString(", ")
	builder.WriteString("is_finished=")
	builder.WriteString(fmt.Sprintf("%v", pi.IsFinished))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", pi.State))
	builder.WriteString(", ")
	builder.WriteString("data_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.DataID))
	builder.WriteString(", ")
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", pi.IsDel))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(pi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("remain_hours=")
	builder.WriteString(fmt.Sprintf("%v", pi.RemainHours))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pi.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProcInsts is a parsable slice of ProcInst.
type ProcInsts []*ProcInst

func (pi ProcInsts) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}
