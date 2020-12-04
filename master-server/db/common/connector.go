package common

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"reflect"
	"strings"
)

type DBConnector struct {
	pool *sql.DB
}

const databaseName = "kiosks"
const dbDriverName = "postgres"

func NewDBConnector(user, password, host, port string) *DBConnector {
	resultConnector := new(DBConnector)
	var err error
	resultConnector.pool, err = sql.Open(dbDriverName, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, databaseName))
	if err != nil {
		panic(err.Error())
	}
	err = resultConnector.pool.Ping()
	if err != nil {
		panic(err.Error())
	}
	return resultConnector
}

func (connector *DBConnector) GetObjectsFromDb(tp reflect.Type) []interface{} {
	fieldNames := make([]string, 0, tp.NumField())
	fieldTypes := make([]reflect.Kind, 0, tp.NumField())
	for i := 0; i < tp.NumField(); i++ {
		fieldNames[i] = tp.Field(i).Name
		fieldTypes[i] = tp.Field(i).Type.Kind()
	}
	return connector.selectRowsFromDb(tp.Name(), fieldNames, fieldTypes, tp)
}

func (connector *DBConnector) selectRowsFromDb(tableName string, fieldNames []string, types []reflect.Kind, tp reflect.Type) []interface{} {
	fieldNamesStr := strings.Join(fieldNames, ", ")
	fieldNamesStr = fieldNamesStr[:len(fieldNamesStr)-2]
	rows, err := connector.pool.Query(fmt.Sprintf("SELECT %s FROM %s", fieldNamesStr, tableName))
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	var resultObjectsArray []interface{}
	for rows.Next() {
		valueHolders := getFieldsForFieldTypes(types)
		rowObject := reflect.New(tp)
		err = rows.Scan(valueHolders...)
		if err != nil {
			panic(err.Error())
		}
		setFieldsToRowObject(rowObject, valueHolders)
		resultObjectsArray = append(resultObjectsArray, rowObject)
	}
	return resultObjectsArray
}

func setFieldsToRowObject(object reflect.Value, fieldValues []interface{}) {
	for i, value := range fieldValues {
		field := object.Field(i)
		switch field.Kind() {
		case reflect.Bool:
			field.SetBool(value.(bool))
		case reflect.String:
			field.SetString(value.(string))
		case reflect.Float64, reflect.Float32:
			field.SetFloat(value.(float64))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(value.(int64))
		default:
			panic(fmt.Sprintf("invalid field type {%s} for entity field set", field.Kind().String()))
		}
	}
}

func getFieldsForFieldTypes(types []reflect.Kind) []interface{} {
	resultFieldsSlice := make([]interface{}, 0, len(types))
	for i, tp := range types {
		resultFieldsSlice[i] = getVariableForReflectKind(tp)
	}
	return resultFieldsSlice
}

func getVariableForReflectKind(kind reflect.Kind) interface{} {
	switch kind {
	case reflect.Int:
		var result int
		return &result
	case reflect.String:
		var result string
		return &result
	case reflect.Int8:
		var result int8
		return &result
	case reflect.Int16:
		var result int16
		return &result
	case reflect.Int32:
		var result int32
		return &result
	case reflect.Int64:
		var result int64
		return &result
	case reflect.Float32:
		var result float32
		return &result
	case reflect.Float64:
		var result float64
		return &result
	case reflect.Bool:
		var result bool
		return &result
	}
	panic(fmt.Sprintf("unexpected type {%s}", kind.String()))
}
