package sqlmocked

import "reflect"

type (
	Extract interface {
		ExtractVals(reflect.Value, func(reflect.Value))
		ExtractTypes(reflect.Type, func(reflect.StructField))
	}
)

type SimpleExtract struct{}

func (s SimpleExtract) ExtractVals(ps reflect.Value, fnc func(reflect.Value)) {
	if ps.Kind() == reflect.Pointer {
		ps = ps.Elem()
	}

	size := ps.NumField()
	for i := 0; i < size; i++ {
		fields := ps.Field(i)
		if fields.Kind() == reflect.Struct {
			s.ExtractVals(fields, fnc)
			continue
		}

		fnc(fields)
	}
}

func (s SimpleExtract) ExtractTypes(typ reflect.Type, fnc func(reflect.StructField)) {
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	size := typ.NumField()
	for i := 0; i < size; i++ {
		fields := typ.Field(i)
		if fields.Type.Kind() == reflect.Struct {
			s.ExtractTypes(fields.Type, fnc)
			continue
		}
		fnc(fields)
	}
}
