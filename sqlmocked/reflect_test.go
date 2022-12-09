package sqlmocked

import (
	"reflect"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestSimpleExtract_ExtractTypes(t *testing.T) {
	// get all tags with nested struct
	var extract Extract = SimpleExtract{}
	var got []string
	var expect = []string{"name", "age", "job"}
	extract.ExtractTypes(reflect.TypeOf(TestStruct{}), func(sf reflect.StructField) {
		tag := sf.Tag.Get("gorm")
		prefix := "column:"
		index := strings.Index(tag, prefix)
		if index != -1 {
			got = append(got, tag[index+len(prefix):])
		}
	})
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("test get all tags failed, got %v", got)
	}

	// get all names
	got = []string{}
	expect = []string{"Name", "Empty", "Age", "Job"}
	extract.ExtractTypes(reflect.TypeOf(TestStruct{}), func(sf reflect.StructField) {
		got = append(got, sf.Name)
	})
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("test get all names failed, got %v", got)
	}
}

func TestSimpleExtract_ExtractVals(t *testing.T) {
	var extract Extract = SimpleExtract{}
	demo := TestStruct{
		Name:  gofakeit.Name(),
		Empty: gofakeit.UUID(),
		Nested: Nested{
			Job: gofakeit.JobTitle(),
		},
	}

	var size int
	extract.ExtractVals(reflect.ValueOf(demo), func(v reflect.Value) {
		if v.Kind() == reflect.String && len(v.String()) != 0 {
			size += 1
		}
	})
	if size != 3 {
		t.Errorf("test get all vals failed, got %d", size)
	}
}
