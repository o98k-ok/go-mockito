package sqlmocked

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestTable_Titles(t *testing.T) {
	record := NewRecord[TestStruct]()

	expect := []string{"name", "age", "job"}
	got := Titles(record)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("table.Titles test faile, got %v, expect %v", got, expect)
	}

	nestRecord := NewRecord[Nested]()
	expect = []string{"job"}
	got = Titles(nestRecord)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("table.Titles test failed, got %v, expect %v", got, expect)
	}
}

func TestTable_Values(t *testing.T) {
	demo := TestStruct{
		Name:  gofakeit.Name(),
		Empty: gofakeit.UUID(),
		Nested: Nested{
			Job: gofakeit.JobTitle(),
		},
	}
	expect := []driver.Value{demo.Name, demo.Empty, demo.Age, demo.Nested.Job}

	got := Values(demo)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("TestTable_Values failed, got %v, expect %v", got, expect)
	}
}

func TestNewRecord(t *testing.T) {
	age, name := "100", "o98k-ok"
	record := NewRecord(func(v *TestStruct) {
		v.Age = age
	}, func(v *TestStruct) {
		v.Name = name
	})
	if age != record.Age || name != record.Name {
		t.Errorf("TestNewRecord failed, got %v", record)
	}
}
