// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNodeId holds the string denoting the nodeid field in the database.
	FieldNodeId = "node_id"
	// FieldProcInstID holds the string denoting the procinstid field in the database.
	FieldProcInstID = "proc_inst_id"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldClaimTime holds the string denoting the claimtime field in the database.
	FieldClaimTime = "claim_time"
	// FieldIsFinished holds the string denoting the isfinished field in the database.
	FieldIsFinished = "is_finished"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// FieldCreateUser holds the string denoting the createuser field in the database.
	FieldCreateUser = "create_user"
	// FieldUpdateUser holds the string denoting the updateuser field in the database.
	FieldUpdateUser = "update_user"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// Table holds the table name of the task in the database.
	Table = "wflow_task"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldNodeId,
	FieldProcInstID,
	FieldCreateTime,
	FieldClaimTime,
	FieldIsFinished,
	FieldUpdateTime,
	FieldCreateUser,
	FieldUpdateUser,
	FieldVersion,
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
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime time.Time
	// DefaultIsFinished holds the default value on creation for the "isFinished" field.
	DefaultIsFinished uint64
	// DefaultUpdateTime holds the default value on creation for the "updateTime" field.
	DefaultUpdateTime time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uint64
)