package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getValue func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getValue = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getValue = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getValue(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {

	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
