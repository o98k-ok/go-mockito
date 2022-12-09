package sqlmocked

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
)

type Table[T any] struct {
	Struct  *T
	srcType string
	extract Extract
}

func NewTable[T any]() *Table[T] {
	res := &Table[T]{
		Struct:  new(T),
		srcType: "gorm",
		extract: SimpleExtract{},
	}
	res.Fresh()
	return res
}

// Titles through gorm:"column:xx"
func (t *Table[T]) Titles() []string {
	var res []string
	t.extract.ExtractTypes(reflect.TypeOf(t.Struct), func(sf reflect.StructField) {
		tag := sf.Tag.Get(t.srcType)
		prefix := "column:"
		index := strings.Index(tag, prefix)
		if index != -1 {
			res = append(res, tag[index+len(prefix):])
		}
	})
	return res
}

func (t *Table[T]) Values() []driver.Value {
	var res []driver.Value
	t.extract.ExtractVals(reflect.ValueOf(t.Struct), func(v reflect.Value) {
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		res = append(res, v.Interface())
	})
	return res
}

func (t *Table[T]) Fresh() {
	gofakeit.Seed(time.Now().Unix())
	gofakeit.Struct(t.Struct)
}

func (t *Table[T]) Rows() *sqlmock.Rows {
	return sqlmock.NewRows(t.Titles()).AddRow(t.Values()...)
}
