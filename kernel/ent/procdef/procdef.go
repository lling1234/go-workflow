// Code generated by entc, DO NOT EDIT.

package procdef

import (
	"time"
)

const (
	// Label holds the string label denoting the procdef type in the database.
	Label = "procdef"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldResource holds the string denoting the resource field in the database.
	FieldResource = "resource"
	// FieldCreateUser holds the string denoting the createuser field in the database.
	FieldCreateUser = "create_user"
	// FieldCreateUserName holds the string denoting the createusername field in the database.
	FieldCreateUserName = "create_user_name"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldTargetID holds the string denoting the targetid field in the database.
	FieldTargetID = "target_id"
	// FieldFormID holds the string denoting the formid field in the database.
	FieldFormID = "form_id"
	// FieldFormName holds the string denoting the formname field in the database.
	FieldFormName = "form_name"
	// FieldAppID holds the string denoting the appid field in the database.
	FieldAppID = "app_id"
	// FieldAppName holds the string denoting the appname field in the database.
	FieldAppName = "app_name"
	// FieldRemainHours holds the string denoting the remainhours field in the database.
	FieldRemainHours = "remain_hours"
	// FieldIsActive holds the string denoting the isactive field in the database.
	FieldIsActive = "is_active"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// FieldDelTime holds the string denoting the deltime field in the database.
	FieldDelTime = "del_time"
	// FieldDelUser holds the string denoting the deluser field in the database.
	FieldDelUser = "del_user"
	// FieldUpdateUser holds the string denoting the updateuser field in the database.
	FieldUpdateUser = "update_user"
	// Table holds the table name of the procdef in the database.
	Table = "wflow_proc_def"
)

// Columns holds all SQL columns for procdef fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCode,
	FieldVersion,
	FieldResource,
	FieldCreateUser,
	FieldCreateUserName,
	FieldCreateTime,
	FieldTargetID,
	FieldFormID,
	FieldFormName,
	FieldAppID,
	FieldAppName,
	FieldRemainHours,
	FieldIsActive,
	FieldUpdateTime,
	FieldDelTime,
	FieldDelUser,
	FieldUpdateUser,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// ResourceValidator is a validator for the "resource" field. It is called by the builders before save.
	ResourceValidator func(string) error
	// CreateUserNameValidator is a validator for the "createUserName" field. It is called by the builders before save.
	CreateUserNameValidator func(string) error
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime time.Time
	// FormNameValidator is a validator for the "formName" field. It is called by the builders before save.
	FormNameValidator func(string) error
	// AppNameValidator is a validator for the "appName" field. It is called by the builders before save.
	AppNameValidator func(string) error
	// DefaultRemainHours holds the default value on creation for the "remainHours" field.
	DefaultRemainHours uint64
	// DefaultIsActive holds the default value on creation for the "isActive" field.
	DefaultIsActive uint64
	// DefaultUpdateTime holds the default value on creation for the "updateTime" field.
	DefaultUpdateTime time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uint64
)
