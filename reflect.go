package types

import "reflect"

// Kind is reflect.Kind
type Kind = reflect.Kind

// Type is reflect.Type
type Type = reflect.Type

// Value is reflect.Value
type Value = reflect.Value

// KindSlice is reflect.Slice
const KindSlice = reflect.Slice

// KindMap is reflect.Map
const KindMap = reflect.Map

// Get returns reflect.TypeOf
func Get(any Any) Type { return reflect.TypeOf(any) }

// KindOf returns Get(any).Kind()
func KindOf(any Any) Kind { return Get(any).Kind() }

// Reflect returns reflect.ValueOf
func Reflect(any Any) Value { return reflect.ValueOf(any) }

// ReflectMap returns reflected Map, or error if any is not a map
func ReflectMap(any Any) (v Map, err error) {
	if val := Reflect(any); val.Kind() != KindMap {
		err = NewErr("not map: " + val.Kind().String())
	} else {
		v = make(Map)
		for _, k := range val.MapKeys() {
			v[k.Interface()] = val.MapIndex(k).Interface()
		}
	}
	return
}

// ReflectSlice returns a Slice from an arg
func ReflectSlice(any Any) (v Slice, err error) {
	if val := Reflect(any); val.Kind() != KindSlice {
		err = NewErr("not slice: " + val.Kind().String())
	} else {
		len := val.Len()
		v = make(Slice, len)
		for i := 0; i < len; i++ {
			v[i] = val.Index(i).Interface()
		}
	}
	return
}
