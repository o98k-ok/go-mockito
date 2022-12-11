package sqlmocked

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestTable_Titles(t *testing.T) {
	table := NewTable[TestStruct]()
	record := table.NewRecord()

	expect := []string{"name", "age", "job"}
	got := table.Titles(record)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("table.Titles test faile, got %v, expect %v", got, expect)
	}

	nested := NewTable[Nested]()
	nestRecord := nested.NewRecord()
	expect = []string{"job"}
	got = nested.Titles(nestRecord)
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
	expect := []driver.Value{demo.Name, demo.Empty, demo.Age, demo.Nested.Job}

	got := table.Values(demo)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("TestTable_Values failed, got %v, expect %v", got, expect)
	}
}
