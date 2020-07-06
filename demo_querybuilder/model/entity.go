package model

//import (
//	"fmt"
//	"reflect"
//)
//// convertMulti will convert any single model to pointer of []model
//func convertMulti(v reflect.Value) reflect.Value {
//	vi := reflect.MakeSlice(reflect.SliceOf(v.Type()), 1, 1)
//	vi.Index(0).Set(v)
//	vv := reflect.New(vi.Type())
//	vv.Elem().Set(vi)
//	return vv
//}
//
//type entity struct {
//	name       string
//	typeOf     reflect.Type
//	isMultiPtr bool
//	slice      reflect.Value
//	//fields     map[string]Column
//	//columns    []Column
//}
//
//// TODO: check primary key must present
//func newEntity(it interface{}) (*entity, error) {
//	v := reflect.ValueOf(it)
//	if v.Kind() != reflect.Ptr {
//		return nil, fmt.Errorf("goloquent: model is not addressable")
//	}
//
//
//	t := v.Type().Elem()
//	switch t.Kind() {
//	case reflect.Slice, reflect.Array:
//		t = t.Elem()
//		if t.Kind() == reflect.Ptr {
//			isMultiPtr = true
//			t = t.Elem()
//		}
//		if t.Kind() != reflect.Struct {
//			return nil, fmt.Errorf("goloquent: invalid entity data type : %v, it should be struct", t)
//		}
//	case reflect.Struct:
//		isMultiPtr = true
//		v = convertMulti(v)
//	default:
//		return nil, fmt.Errorf("goloquent: invalid entity data type : %v, it should be struct", t)
//	}


	//fields := make(map[string]Column)
	//cols := getColumns(nil, codec)
	//for _, c := range cols {
	//	fields[c.Name()] = c
	//}
	//
	//if _, hasKey := fields[keyFieldName]; !hasKey {
	//	return nil, fmt.Errorf("goloquent: entity %v doesn't has primary key property", t)
	//}
	//
	//return &entity{
	//	name:       t.Name(),
	//	typeOf:     t,
	//	isMultiPtr: isMultiPtr,
	//	codec:      codec,
	//	slice:      v,
	//	fields:     fields,
	//	columns:    cols,
	//}, nil
//}
