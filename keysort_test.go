package keysort_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/itroot/keysort"
)

func TestSort_Indentical(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) []interface{} {
		return []interface{}{slice[i]}
	})
	if !reflect.DeepEqual(slice, []int{1, 2, 3, 4, 5}) {
		t.Fatal()
	}
}

func TestSort_String(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) []interface{} {
		return []interface{}{strconv.Itoa(slice[i])}
	})
	if !reflect.DeepEqual(slice, []int{1, 2, 3, 4, 5}) {
		t.Fatal()
	}
}

func TestSort_MultisortString(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) []interface{} {
		return []interface{}{int(1), strconv.Itoa(slice[i])}
	})
	if !reflect.DeepEqual(slice, []int{1, 2, 3, 4, 5}) {
		t.Fatal()
	}
}

func TestSort_Reversed(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) []interface{} {
		return []interface{}{-slice[i]}
	})
	if !reflect.DeepEqual(slice, []int{5, 4, 3, 2, 1}) {
		t.Fatal()
	}
}

func TestSort_EvenOdd(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) []interface{} {
		return []interface{}{slice[i]%2 == 1, slice[i]}
	})
	if !reflect.DeepEqual(slice, []int{2, 4, 1, 3, 5}) {
		t.Fatal()
	}
}
