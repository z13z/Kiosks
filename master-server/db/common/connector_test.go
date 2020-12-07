package common

import (
	"database/sql"
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
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]interface{}
	}{{name: "emptTableTest", fields: fields{getEmptyDbPool()}, args: args{args: MockEntity{}}, want: &[]interface{}{}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := &DBConnector{
				pool: tt.fields.pool,
			}
			if got := connector.GetObjectsFromDb(&tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObjectsFromDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getEmptyDbPool() *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("problem with db mocking")
	}
	mock.ExpectQuery("SELECT id, name FROM Mock").WillReturnRows()
	return db
}

type MockEntity struct {
	id   int64
	name string
}

func (entity *MockEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	entity.id = idValue.(int64)
	entity.name = nameValue.(string)
}

func (entity *MockEntity) GetTableName() string {
	return "Mock"
}

func (entity *MockEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name"}
}

func (entity *MockEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{entity.id, entity.name})
}

func (entity *MockEntity) NewEntity() IEntity {
	return new(MockEntity)
}
