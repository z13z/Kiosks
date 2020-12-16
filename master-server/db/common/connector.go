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

//todo move to external configuration
const dbName = "kiosks"
const dbDriverName = "postgres"
const dbUser = "postgres"
const dbPassword = "postgres"
const dbHost = "localhost"
const dbPort = "5432"

func NewDBConnector() *DBConnector {
	resultConnector := new(DBConnector)
	var err error
	resultConnector.pool, err = sql.Open(dbDriverName, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		panic(err.Error())
	}
	err = resultConnector.pool.Ping()
	if err != nil {
		panic(err.Error())
	}
	return resultConnector
}

func (connector *DBConnector) GetObjectsFromDb(entity IEntity, whereParams *map[string]string, offset, limit int) *[]interface{} {
	wherePart, wherePartValues := parseWherePartAndParamValues(whereParams)
	return connector.selectRowsFromDb(entity.GetTableName(), entity.GetFieldNames(), entity.GetFieldValueHolders(),
		entity.NewEntity, wherePart, wherePartValues, offset, limit)
}

func parseWherePartAndParamValues(params *map[string]string) (*string, *[]interface{}) {
	if params == nil || len(*params) == 0 {
		return nil, &[]interface{}{}
	}
	whereQuery := "WHERE "
	var whereQueryValues []interface{}
	paramNumber := 1
	for column, value := range *params {
		whereQuery += column + fmt.Sprintf(" = $%d, ", paramNumber)
		paramNumber++
		whereQueryValues = append(whereQueryValues, value)
	}
	//remove last ', ' part
	whereQuery = whereQuery[:len(whereQuery)-2]
	return &whereQuery, &whereQueryValues
}

func (connector *DBConnector) selectRowsFromDb(tableName string, fieldNames *[]string, fieldValueHolders *[]interface{},
	entityCreator func() IEntity, wherePart *string, wherePartValues *[]interface{}, offset, limit int) *[]interface{} {
	var wherePartVal string
	if wherePart == nil {
		wherePartVal = ""
	} else {
		wherePartVal = *wherePart
	}
	fieldNamesStr := strings.Join(*fieldNames, ", ")
	rows, err := connector.pool.Query(fmt.Sprintf("SELECT %s FROM %s %s ORDER BY id LIMIT %d OFFSET %d",
		fieldNamesStr, tableName, wherePartVal, limit, offset), *wherePartValues...)
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
