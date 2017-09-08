package util

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var infos = make(map[reflect.Type]*structInfo)
var infoLock = &sync.RWMutex{}
var infoKey = reflect.TypeOf(&structInfo{})
var valueKey = reflect.TypeOf(&reflect.Value{})

type structInfoField struct {
	FieldName string
	KeyName   string
	Field     *reflect.StructField
}
type structInfo struct {
	Type   reflect.Type
	Fields []*structInfoField
	setter StructSetter
}

type StructSetter interface {
	FillStruct(ctx context.Context, obj interface{}, m map[string]interface{}) error
}

type StructSetterFunc func(ctx context.Context, obj interface{}, m map[string]interface{}) error

func (self StructSetterFunc) FillStruct(ctx context.Context, obj interface{}, m map[string]interface{}) error {
	return self(ctx, obj, m)
}

func SetStruct(obj interface{}, m map[string]interface{}) error {
	val := reflect.ValueOf(obj)
	t := val.Type()

	if val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
		t = val.Type()
	}

	info, err := getOrCreateInfo(t)
	if err != nil {
		return err
	}

	ctx := context.WithValue(nil, infoKey, info)
	ctx = context.WithValue(ctx, valueKey, val)
	return info.setter.FillStruct(ctx, obj, m)
}

func simpleStructSetter(ctx context.Context, obj interface{}, m map[string]interface{}) error {
	info := ctx.Value(infoKey).(*structInfo)
	val := ctx.Value(valueKey).(reflect.Value)

	for i, f := range info.Fields {
		v, ok := m[f.KeyName]
		if !ok {
			continue
		}
		val := val.Field(i)
		val.Set(reflect.ValueOf(v))
	}

	return nil
}

func getOrCreateSetter(t reflect.Type) (StructSetter, error) {
	i, err := getOrCreateInfo(t)
	if err != nil {
		return nil, err
	}
	return i.setter, nil
}
func getOrCreateInfo(t reflect.Type) (*structInfo, error) {
	info := infos[t]
	if info != nil {
		return info, nil
	}

	infoLock.Lock()
	defer infoLock.Unlock()

	info = infos[t]
	if info != nil {
		return info, nil
	}

	n := t.NumField()
	info = &structInfo{
		Fields: make([]*structInfoField, n),
	}
	for i := 0; i < n; i++ {
		field := t.Field(i)
		f := &structInfoField{
			Field:     &field,
			FieldName: field.Name,
		}
		info.Fields[i] = f
		tag := field.Tag.Get("astgo")
		if tag != "" {
			f.KeyName = tag
		} else {
			f.KeyName = f.FieldName
		}
	}
	info.setter = StructSetterFunc(simpleStructSetter)
	infos[t] = info
	return info, nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
