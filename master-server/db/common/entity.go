package common

type IEntity interface {
	SetEntityFields(fields map[string]interface{})
	GetFieldValueHolders() *[]interface{}
	GetTableName() string
	GetFieldNames() *[]string
	NewEntity() IEntity
	GetEditableFieldValueHolders() *[]interface{}
	GetEditableFieldNames() *[]string
	GetId() int64
}
