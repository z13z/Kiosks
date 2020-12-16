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
		{name: "emptyRowTableTest", fields: fields{getRowsDbPool(false, 0, math.MaxInt32, "")},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32}, want: &[]interface{}{}},
		{name: "oneRowTableTest", fields: fields{getRowsDbPool(false, 0, math.MaxInt32, "", mockEntitiesToTest[0])},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32}, want: &[]interface{}{&mockEntitiesToTest[0]}},
		{name: "multipleRowTableTest", fields: fields{getRowsDbPool(false, 0, math.MaxInt32, "", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32}, want: &mockEntityHoldersToTest},
		{name: "multipleRowTableWhereClauseTest", fields: fields{getRowsDbPool(true, 0, math.MaxInt32, "WHERE id = \\$1", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, limit: math.MaxInt32, whereParams: map[string]string{"id": "0"}}, want: &[]interface{}{&mockEntitiesToTest[0]}},
		{name: "multipleRowTableWithOffsetTest", fields: fields{getRowsDbPool(true, 1, 1, "", mockEntitiesToTest...)},
			args: args{entityArg: MockEntity{}, whereParams: map[string]string{}, offset: 1, limit: 1}, want: &[]interface{}{&mockEntitiesToTest[0]}},
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

func getRowsDbPool(onlyFirstEntry bool, offset, limit int, wherePart string, entities ...MockEntity) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	var rows *sqlmock.Rows
	rows = sqlmock.NewRows([]string{"id", "name"})
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
