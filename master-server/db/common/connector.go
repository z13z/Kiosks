package common

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

func (connector *DBConnector) GetObjectsFromDb(entity IEntity) *[]interface{} {
	return connector.selectRowsFromDb(entity.GetTableName(), entity.GetFieldNames(), entity.GetFieldValueHolders(), entity.NewEntity)
}

func (connector *DBConnector) selectRowsFromDb(tableName string, fieldNames *[]string, fieldValueHolders *[]interface{},
	entityCreator func() IEntity) *[]interface{} {
	fieldNamesStr := strings.Join(*fieldNames, ", ")
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
		rowObject := entityCreator()
		err = rows.Scan(*fieldValueHolders...)
		rowObject.SetEntityFields(getFieldNamesAndValuesMap(fieldNames, fieldValueHolders))
		if err != nil {
			panic(err)
		}
		resultObjectsArray = append(resultObjectsArray, rowObject)
	}
	return &resultObjectsArray
}

func getFieldNamesAndValuesMap(names *[]string, holders *[]interface{}) map[string]interface{} {
	resultFieldNamesAndValuesMap := make(map[string]interface{})
	for ind, name := range *names {
		resultFieldNamesAndValuesMap[name] = (*holders)[ind]
	}
	return resultFieldNamesAndValuesMap
}
