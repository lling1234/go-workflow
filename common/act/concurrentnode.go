// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/concurrentnode"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ConcurrentNode is the model entity for the ConcurrentNode schema.
type ConcurrentNode struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 流程实例id
	ProcInstID int64 `json:"proc_inst_id,omitempty"`
	// 流程定义id
	ProcDefID int64 `json:"proc_def_id,omitempty"`
	// 节点id
	NodeID string `json:"node_id,omitempty"`
	// 节点信息
	NodeInfo string `json:"node_info,omitempty"`
	// 上一节点ID
	PrevID string `json:"prev_id,omitempty"`
	// 下一节点ID
	NextID string `json:"next_id,omitempty"`
	// 节点状态 1、未审批  2、正在审批  3、已审批完成
	State int32 `json:"state,omitempty"`
	// 是否删除,0:未删除,1:已删除
	IsDel int32 `json:"is_del,omitempty"`
	// 流程创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 流程修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ConcurrentNode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case concurrentnode.FieldID, concurrentnode.FieldProcInstID, concurrentnode.FieldProcDefID, concurrentnode.FieldState, concurrentnode.FieldIsDel:
			values[i] = new(sql.NullInt64)
		case concurrentnode.FieldNodeID, concurrentnode.FieldNodeInfo, concurrentnode.FieldPrevID, concurrentnode.FieldNextID:
			values[i] = new(sql.NullString)
		case concurrentnode.FieldCreateTime, concurrentnode.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ConcurrentNode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ConcurrentNode fields.
func (cn *ConcurrentNode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case concurrentnode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cn.ID = int64(value.Int64)
		case concurrentnode.FieldProcInstID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proc_inst_id", values[i])
			} else if value.Valid {
				cn.ProcInstID = value.Int64
			}
		case concurrentnode.FieldProcDefID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proc_def_id", values[i])
			} else if value.Valid {
				cn.ProcDefID = value.Int64
			}
		case concurrentnode.FieldNodeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				cn.NodeID = value.String
			}
		case concurrentnode.FieldNodeInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_info", values[i])
			} else if value.Valid {
				cn.NodeInfo = value.String
			}
		case concurrentnode.FieldPrevID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prev_id", values[i])
			} else if value.Valid {
				cn.PrevID = value.String
			}
		case concurrentnode.FieldNextID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field next_id", values[i])
			} else if value.Valid {
				cn.NextID = value.String
			}
		case concurrentnode.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				cn.State = int32(value.Int64)
			}
		case concurrentnode.FieldIsDel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				cn.IsDel = int32(value.Int64)
			}
		case concurrentnode.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cn.CreateTime = value.Time
			}
		case concurrentnode.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cn.UpdateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ConcurrentNode.
// Note that you need to call ConcurrentNode.Unwrap() before calling this method if this ConcurrentNode
// was returned from a transaction, and the transaction was committed or rolled back.
func (cn *ConcurrentNode) Update() *ConcurrentNodeUpdateOne {
	return (&ConcurrentNodeClient{config: cn.config}).UpdateOne(cn)
}

// Unwrap unwraps the ConcurrentNode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cn *ConcurrentNode) Unwrap() *ConcurrentNode {
	_tx, ok := cn.config.driver.(*txDriver)
	if !ok {
		panic("act: ConcurrentNode is not a transactional entity")
	}
	cn.config.driver = _tx.drv
	return cn
}

// String implements the fmt.Stringer.
func (cn *ConcurrentNode) String() string {
	var builder strings.Builder
	builder.WriteString("ConcurrentNode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cn.ID))
	builder.WriteString("proc_inst_id=")
	builder.WriteString(fmt.Sprintf("%v", cn.ProcInstID))
	builder.WriteString(", ")
	builder.WriteString("proc_def_id=")
	builder.WriteString(fmt.Sprintf("%v", cn.ProcDefID))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(cn.NodeID)
	builder.WriteString(", ")
	builder.WriteString("node_info=")
	builder.WriteString(cn.NodeInfo)
	builder.WriteString(", ")
	builder.WriteString("prev_id=")
	builder.WriteString(cn.PrevID)
	builder.WriteString(", ")
	builder.WriteString("next_id=")
	builder.WriteString(cn.NextID)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", cn.State))
	builder.WriteString(", ")
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", cn.IsDel))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(cn.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(cn.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ConcurrentNodes is a parsable slice of ConcurrentNode.
type ConcurrentNodes []*ConcurrentNode

func (cn ConcurrentNodes) config(cfg config) {
	for _i := range cn {
		cn[_i].config = cfg
	}
}
