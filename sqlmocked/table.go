package sqlmocked

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
)

var (
	GloabalExtract Extract = SimpleExtract{}
	DefaultTagName         = "gorm"
)

func init() {
	gofakeit.Seed(time.Now().Unix())
}

func NewRecord[T any](customFncs ...func(v *T)) T {
	record := new(T)
	gofakeit.Struct(record)
	for _, fnc := range customFncs {
		fnc(record)
	}
	return *record
}

// Titles through gorm:"column:xx"
func Titles[T any](record T) []string {
	var res []string
	GloabalExtract.ExtractTypes(reflect.TypeOf(&record), func(sf reflect.StructField) {
		tag := sf.Tag.Get(DefaultTagName)
		prefix := "column:"
		index := strings.Index(tag, prefix)
		if index != -1 {
			res = append(res, tag[index+len(prefix):])
		}
	})
	return res
}

func Values[T any](record T) []driver.Value {
	var res []driver.Value
	GloabalExtract.ExtractVals(reflect.ValueOf(&record), func(v reflect.Value) {
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		res = append(res, v.Interface())
	})
	return res
}

func GormRow[T any](record T) *sqlmock.Rows {
	return sqlmock.NewRows(Titles(record)).AddRow(Values(record)...)
}
