package mod

import "reflect"

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}
