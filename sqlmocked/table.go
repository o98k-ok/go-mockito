package sqlmocked

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(time.Now().Unix())
}

type Table[T any] struct {
	srcType string
	extract Extract
}

func NewTable[T any]() *Table[T] {
	res := &Table[T]{
		srcType: "gorm",
		extract: SimpleExtract{},
	}
	return res
}

func (t *Table[T]) NewRecord() T {
	record := new(T)
	gofakeit.Struct(record)
	return *record
}

// Titles through gorm:"column:xx"
func (t *Table[T]) Titles(record T) []string {
	var res []string
	t.extract.ExtractTypes(reflect.TypeOf(&record), func(sf reflect.StructField) {
		tag := sf.Tag.Get(t.srcType)
		prefix := "column:"
		index := strings.Index(tag, prefix)
		if index != -1 {
			res = append(res, tag[index+len(prefix):])
		}
	})
	return res
}

func (t *Table[T]) Values(record T) []driver.Value {
	var res []driver.Value
	t.extract.ExtractVals(reflect.ValueOf(&record), func(v reflect.Value) {
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		res = append(res, v.Interface())
	})
	return res
}

func (t *Table[T]) Row(record T) *sqlmock.Rows {
	return sqlmock.NewRows(t.Titles(record)).AddRow(t.Values(record)...)
}
