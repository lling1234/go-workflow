// Code generated by entc, DO NOT EDIT.

package procinst

import (
	"time"
)

const (
	// Label holds the string label denoting the procinst type in the database.
	Label = "procinst"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProcDefID holds the string denoting the procdefid field in the database.
	FieldProcDefID = "proc_def_id"
	// FieldRefID holds the string denoting the refid field in the database.
	FieldRefID = "ref_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldTargetID holds the string denoting the targetid field in the database.
	FieldTargetID = "target_id"
	// FieldResource holds the string denoting the resource field in the database.
	FieldResource = "resource"
	// FieldNodeID holds the string denoting the nodeid field in the database.
	FieldNodeID = "node_id"
	// FieldTaskID holds the string denoting the taskid field in the database.
	FieldTaskID = "task_id"
	// FieldConNodeIDs holds the string denoting the connodeids field in the database.
	FieldConNodeIDs = "con_node_ids"
	// FieldConTaskIDs holds the string denoting the contaskids field in the database.
	FieldConTaskIDs = "con_task_ids"
	// FieldIsFinished holds the string denoting the isfinished field in the database.
	FieldIsFinished = "is_finished"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldDataID holds the string denoting the dataid field in the database.
	FieldDataID = "data_id"
	// FieldUpdateUser holds the string denoting the updateuser field in the database.
	FieldUpdateUser = "update_user"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldFinishTime holds the string denoting the finishtime field in the database.
	FieldFinishTime = "finish_time"
	// FieldCreateUser holds the string denoting the createuser field in the database.
	FieldCreateUser = "create_user"
	// FieldCreateUsername holds the string denoting the createusername field in the database.
	FieldCreateUsername = "create_username"
	// FieldRemainHours holds the string denoting the remainhours field in the database.
	FieldRemainHours = "remain_hours"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// FieldFlowType holds the string denoting the flowtype field in the database.
	FieldFlowType = "flow_type"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldDelTime holds the string denoting the deltime field in the database.
	FieldDelTime = "del_time"
	// FieldDelUser holds the string denoting the deluser field in the database.
	FieldDelUser = "del_user"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// Table holds the table name of the procinst in the database.
	Table = "wflow_proc_inst"
)

// Columns holds all SQL columns for procinst fields.
var Columns = []string{
	FieldID,
	FieldProcDefID,
	FieldRefID,
	FieldTitle,
	FieldCode,
	FieldTargetID,
	FieldResource,
	FieldNodeID,
	FieldTaskID,
	FieldConNodeIDs,
	FieldConTaskIDs,
	FieldIsFinished,
	FieldState,
	FieldDataID,
	FieldUpdateUser,
	FieldCreateTime,
	FieldFinishTime,
	FieldCreateUser,
	FieldCreateUsername,
	FieldRemainHours,
	FieldUpdateTime,
	FieldFlowType,
	FieldRemark,
	FieldDelTime,
	FieldDelUser,
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// ResourceValidator is a validator for the "resource" field. It is called by the builders before save.
	ResourceValidator func(string) error
	// NodeIDValidator is a validator for the "nodeID" field. It is called by the builders before save.
	NodeIDValidator func(string) error
	// ConNodeIDsValidator is a validator for the "conNodeIDs" field. It is called by the builders before save.
	ConNodeIDsValidator func(string) error
	// ConTaskIDsValidator is a validator for the "conTaskIDs" field. It is called by the builders before save.
	ConTaskIDsValidator func(string) error
	// DefaultIsFinished holds the default value on creation for the "isFinished" field.
	DefaultIsFinished uint64
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime time.Time
	// DefaultFinishTime holds the default value on creation for the "finishTime" field.
	DefaultFinishTime time.Time
	// CreateUsernameValidator is a validator for the "createUsername" field. It is called by the builders before save.
	CreateUsernameValidator func(string) error
	// DefaultUpdateTime holds the default value on creation for the "updateTime" field.
	DefaultUpdateTime time.Time
	// DefaultFlowType holds the default value on creation for the "flowType" field.
	DefaultFlowType uint64
	// RemarkValidator is a validator for the "remark" field. It is called by the builders before save.
	RemarkValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uint64
)