package tipos

import "reflect"

func ShowType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
