package validator

import (
	"fmt"
	"reflect"
)

var (
	typeFmtStringer = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

func isTypeFmtStringer(value reflect.Value) bool {
	// Get a type object of the pointer on the object represented by the parameter
	// and see if it implements
	return reflect.PtrTo(value.Type()).Implements(typeFmtStringer)
}
