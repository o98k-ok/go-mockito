package sqlmocked

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestTable_Titles(t *testing.T) {
	var table SqlMockRepo = NewTable[TestStruct]()
	expect := []string{"name", "age", "job"}
	got := table.Titles()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("table.Titles test faile, got %v, expect %v", got, expect)
	}

	nested := NewTable[Nested]()
	expect = []string{"job"}
	got = nested.Titles()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("table.Titles test failed, got %v, expect %v", got, expect)
	}
}

func TestTable_Values(t *testing.T) {
	table := NewTable[TestStruct]()
	demo := TestStruct{
		Name:  gofakeit.Name(),
		Empty: gofakeit.UUID(),
		Nested: Nested{
			Job: gofakeit.JobTitle(),
		},
	}
	table.Struct = demo
	expect := []driver.Value{demo.Name, demo.Empty, demo.Age, demo.Nested.Job}

	got := table.Values()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("TestTable_Values failed, got %v, expect %v", got, expect)
	}
}

func TestTable_Fresh(t *testing.T) {
	type Status struct {
		Id    *uint32 `gorm:"column:id" json:"-" fake:"skip"`
		Name  *string `gorm:"column:name" json:"name" fake:"{firstname}"`
		Code  *string `gorm:"column:code" json:"code" fake:"{uuid}"`
		Token *string `gorm:"column:token" json:"token" fake:"{uuid}"`
	}
	table := NewTable[Status]()
	table.Fresh()

	dat, err := json.Marshal(table.Struct)
	if err != nil {
		t.Errorf("test table data fresh failed %v", err)
	}
	t.Log(string(dat))
}
