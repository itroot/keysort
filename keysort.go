package keysort

import (
	"reflect"
	"sort"
)

// Sort will sort slice by comparing key function values
func Sort(slice interface{}, key func(i int) []interface{}) {
	sort.Slice(slice, func(i, j int) bool {
		return less(key(i), key(j))
	})
}

func less(lhs []interface{}, rhs []interface{}) bool {
	if len(lhs) == 0 {
		return true
	}
	if len(rhs) == 0 {
		return false
	}
	if lhs[0] == rhs[0] {
		return less(lhs[1:], rhs[1:])
	}
	if reflect.TypeOf(lhs[0]) == reflect.TypeOf(rhs[0]) {
		switch reflect.TypeOf(lhs[0]).Kind() {
		case reflect.String:
			return lhs[0].(string) < rhs[0].(string)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(lhs[0]).Uint() < reflect.ValueOf(rhs[0]).Uint()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(lhs[0]).Int() < reflect.ValueOf(rhs[0]).Int()
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(lhs[0]).Float() < reflect.ValueOf(rhs[0]).Float()
		case reflect.Bool:
			return !lhs[0].(bool) && rhs[0].(bool)
		default:
			panic("Unsupported type")
		}
	} else {
		return reflect.TypeOf(lhs[0]).String() < reflect.TypeOf(rhs[0]).String()
	}
}
