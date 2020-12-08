package common

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func TestDBConnector_GetObjectsFromDb(t *testing.T) {
	type fields struct {
		pool *sql.DB
	}
	type args struct {
		args MockEntity
	}

	mockEntitiesToTest := make([]MockEntity, 100)
	mockEntityHoldersToTest := make([]interface{}, 100)
	for i := int64(0); i < 100; i++ {
		mockEntitiesToTest[i] = MockEntity{id: i, name: ""}
		mockEntityHoldersToTest[i] = &mockEntitiesToTest[i]
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]interface{}
	}{
		{name: "EmptyRowTableTest", fields: fields{getRowsDbPool()}, args: args{args: MockEntity{}}, want: &[]interface{}{}},
		{name: "oneRowTableTest", fields: fields{getRowsDbPool(mockEntitiesToTest[0])}, args: args{args: MockEntity{}}, want: &[]interface{}{&mockEntitiesToTest[0]}},
		{name: "multipleRowTableTest", fields: fields{getRowsDbPool(mockEntitiesToTest...)}, args: args{args: MockEntity{}}, want: &mockEntityHoldersToTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := &DBConnector{
				pool: tt.fields.pool,
			}
			if got := connector.GetObjectsFromDb(&tt.args.args); len(*tt.want) != len(*got) || len(*got) != 0 && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObjectsFromDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getRowsDbPool(entities ...MockEntity) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	var rows *sqlmock.Rows
	rows = sqlmock.NewRows([]string{"id", "name"})
	for _, entity := range entities {
		rows.AddRow(driver.Value(entity.id), driver.Value(entity.name))
	}
	mock.ExpectQuery("SELECT id, name FROM Mock").WillReturnRows(rows)
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
