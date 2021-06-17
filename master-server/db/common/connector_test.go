package common

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math"
	"reflect"
	"testing"
)

func TestDBConnector_GetObjectsFromDb(t *testing.T) {
	type fields struct {
		pool *sql.DB
	}
	type args struct {
		entityArg   MockEntity
		whereParams map[string]string
		offset      int
		limit       int
	}

	mockEntitiesToTest := make([]MockEntity, 100)
	mockEntityHoldersToTest := make([]interface{}, 100)
	for i := int64(0); i < 100; i++ {
		mockEntitiesToTest[i] = MockEntity{id: i, name: fmt.Sprintf("test_kiosk_%d", i)}
		mockEntityHoldersToTest[i] = &mockEntitiesToTest[i]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]interface{}
	}{
		{name: "emptyRowTableTest", fields: fields{getSelectRowsDbPool(false, 0, math.MaxInt32, "")},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32}, want: &[]interface{}{}},
		{name: "oneRowTableTest", fields: fields{getSelectRowsDbPool(false, 0, math.MaxInt32, "", mockEntitiesToTest[0])},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32}, want: &[]interface{}{&mockEntitiesToTest[0]}},
		{name: "multipleRowTableTest", fields: fields{getSelectRowsDbPool(false, 0, math.MaxInt32, "", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32}, want: &mockEntityHoldersToTest},
		{name: "multipleRowTableWhereClauseTest", fields: fields{getSelectRowsDbPool(true, 0, math.MaxInt32, "WHERE id = \\$1", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32, whereParams: map[string]string{"id": "0"}}, want: &[]interface{}{&mockEntitiesToTest[0]}},
		{name: "multipleRowTableWithOffsetTest", fields: fields{getSelectRowsDbPool(true, 1, 1, "", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, whereParams: map[string]string{}, offset: 1, limit: 1}, want: &[]interface{}{&mockEntitiesToTest[0]}},
		{name: "whereClauseMultipleColumnsTest", fields: fields{getSelectRowsDbPool(true, 0, math.MaxInt32, "WHERE id = \\$1 AND name = \\$2", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32, whereParams: map[string]string{"id": "0", "name": "k"}}, want: &[]interface{}{&mockEntitiesToTest[0]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := &DBConnector{
				pool: tt.fields.pool,
			}
			if got := connector.GetObjectsFromDb(&tt.args.entityArg, &tt.args.whereParams, tt.args.offset, tt.args.limit); len(*tt.want) != len(*got) || len(*got) != 0 && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObjectsFromDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnector_UpdateObjectInDb(t *testing.T) {
	type fields struct {
		pool *sql.DB
	}
	type args struct {
		entity IEntity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{"basicUpdateQuery", fields{pool: getUpdateRowsDbPool(1, "mock", 1)}, args{entity: &MockEntity{id: 1, name: "mock"}}, 1},
		{"emptyUpdateQuery", fields{pool: getUpdateRowsDbPool(1, "mock", 0)}, args{entity: &MockEntity{id: 1, name: "mock"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := &DBConnector{
				pool: tt.fields.pool,
			}
			if got := connector.UpdateObjectInDb(tt.args.entity); got != tt.want {
				t.Errorf("UpdateObjectInDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnector_InsertObjectInDb(t *testing.T) {
	type fields struct {
		pool *sql.DB
	}
	type args struct {
		entity IEntity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"basicInsertQuery", fields{pool: getInsertRowsDbPool("mock", 1)}, args{entity: &MockEntity{id: 1, name: "mock"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := &DBConnector{
				pool: tt.fields.pool,
			}
			if got := connector.InsertObjectInDb(tt.args.entity); got != tt.want {
				t.Errorf("InsertObjectInDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnector_GetObjectsCountFromDb(t *testing.T) {
	type fields struct {
		pool *sql.DB
	}
	type args struct {
		entity      IEntity
		whereParams *map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "emptyTableRowsCountTest", fields: fields{getRowsCountDbPool(0, "")},
			args: args{entity: &MockEntity{}}, want: 0},
		{name: "oneRowTableRowsCountTest", fields: fields{getRowsCountDbPool(1, "WHERE id = \\$1")},
			args: args{entity: &MockEntity{}, whereParams: &map[string]string{"id": "0"}}, want: 1},
		{name: "multipleRowsTableRowsCountTest", fields: fields{getRowsCountDbPool(21, "WHERE id = \\$1 AND name = \\$2")},
			args: args{entity: &MockEntity{}, whereParams: &map[string]string{"id": "2", "name": "z13"}}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := &DBConnector{
				pool: tt.fields.pool,
			}
			if got := connector.GetObjectsCountFromDb(tt.args.entity, tt.args.whereParams); got != tt.want {
				t.Errorf("GetObjectsCountFromDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getRowsCountDbPool(count int, wherePart string) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	rows := sqlmock.NewRows([]string{"count"})
	rows.AddRow(driver.Value(count))
	mock.ExpectQuery("SELECT COUNT\\(1\\) FROM Mock " + wherePart).WillReturnRows(rows)
	return db
}

func getSelectRowsDbPool(onlyFirstEntry bool, offset, limit int, wherePart string, entities ...MockEntity) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	rows := sqlmock.NewRows([]string{"id", "name"})
	if onlyFirstEntry {
		if entities != nil && len(entities) > 0 {
			rows.AddRow(driver.Value(entities[0].id), driver.Value(entities[0].name))
		}
	} else {
		for _, entity := range entities {
			rows.AddRow(driver.Value(entity.id), driver.Value(entity.name))
		}
	}
	mock.ExpectQuery(fmt.Sprintf("SELECT id, name FROM Mock %s ORDER BY id LIMIT %d OFFSET %d", wherePart, limit, offset)).WillReturnRows(rows)
	return db
}

func getUpdateRowsDbPool(id int64, name string, updated int64) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	mock.ExpectExec("UPDATE Mock SET name = \\$2 WHERE id = \\$1").WithArgs(driver.Value(id), driver.Value(name)).WillReturnResult(sqlmock.NewResult(21, updated))
	return db
}

func getInsertRowsDbPool(name string, updated int64) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	mock.ExpectExec("INSERT INTO Mock\\(name\\) VALUES \\(\\$1\\)").WithArgs(driver.Value(name)).WillReturnResult(sqlmock.NewResult(21, updated))
	return db
}

type MockEntity struct {
	id   int64
	name string
}

func (entity *MockEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	entity.id = *idValue.(*int64)
	entity.name = *nameValue.(*string)
}

func (entity *MockEntity) GetTableName() string {
	return "Mock"
}

func (entity *MockEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name"}
}

func (entity *MockEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&entity.id, &entity.name})
}

func (entity *MockEntity) NewEntity() IEntity {
	return new(MockEntity)
}

func (entity *MockEntity) GetEditableFieldValueHolders() *[]interface{} {
	return &([]interface{}{&entity.name})
}

func (entity *MockEntity) GetEditableFieldNames() *[]string {
	return &[]string{"name"}
}

func (entity *MockEntity) GetId() int64 {
	return entity.id
}
