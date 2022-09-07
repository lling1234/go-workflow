// Code generated by ent, DO NOT EDIT.

package execution

import (
	"time"
)

const (
	// Label holds the string label denoting the execution type in the database.
	Label = "execution"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProcInstID holds the string denoting the proc_inst_id field in the database.
	FieldProcInstID = "proc_inst_id"
	// FieldProcDefID holds the string denoting the proc_def_id field in the database.
	FieldProcDefID = "proc_def_id"
	// FieldNodeInfos holds the string denoting the node_infos field in the database.
	FieldNodeInfos = "node_infos"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the execution in the database.
	Table = "act_execution"
)

// Columns holds all SQL columns for execution fields.
var Columns = []string{
	FieldID,
	FieldProcInstID,
	FieldProcDefID,
	FieldNodeInfos,
	FieldStartTime,
	FieldIsDel,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NodeInfosValidator is a validator for the "node_infos" field. It is called by the builders before save.
	NodeInfosValidator func(string) error
	// DefaultStartTime holds the default value on creation for the "start_time" field.
	DefaultStartTime time.Time
	// DefaultIsDel holds the default value on creation for the "is_del" field.
	DefaultIsDel int32
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime time.Time
)
