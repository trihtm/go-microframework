package support

import (
	"reflect"
)

func IsPointer(v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return NonPointerError{rv.Type()}
	}

	return nil
}

type NonPointerError struct {
	Type reflect.Type
}

func (e NonPointerError) Error() string {
	return "non-pointer object: " + e.Type.String() + ", object should be a pointer."
}
