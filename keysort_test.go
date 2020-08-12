package keysort_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/itroot/keysort"
)

func TestSort_Indentical(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		return keysort.Sequence{slice[i]}
	})
	if !reflect.DeepEqual(slice, []int{1, 2, 3, 4, 5}) {
		t.Fatal()
	}
}

func TestSort_String(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		return keysort.Sequence{strconv.Itoa(slice[i])}
	})
	if !reflect.DeepEqual(slice, []int{1, 2, 3, 4, 5}) {
		t.Fatal()
	}
}

func TestSort_MultisortString(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		return keysort.Sequence{int(1), strconv.Itoa(slice[i])}
	})
	if !reflect.DeepEqual(slice, []int{1, 2, 3, 4, 5}) {
		t.Fatal(slice)
	}
}

func TestSort_Reversed(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		return keysort.Sequence{-slice[i]}
	})
	if !reflect.DeepEqual(slice, []int{5, 4, 3, 2, 1}) {
		t.Fatal()
	}
}

func TestSort_EvenOdd(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		return keysort.Sequence{slice[i]%2 == 1, slice[i]}
	})
	if !reflect.DeepEqual(slice, []int{2, 4, 1, 3, 5}) {
		t.Fatal(slice)
	}
}

func TestSort_Desc(t *testing.T) {
	slice := []int{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		return keysort.Sequence{slice[i]%2 == 1, keysort.StringDesc(strconv.Itoa(slice[i]))}
	})
	if !reflect.DeepEqual(slice, []int{4, 2, 5, 3, 1}) {
		t.Fatal()
	}
}

func TestSort_CustomType(t *testing.T) {
	type CustomInt int
	type CustomString string
	slice := []CustomInt{2, 1, 3, 4, 5}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		var cs CustomString
		cs = CustomString(strconv.Itoa(int(slice[i])))
		return keysort.Sequence{slice[i]%2 == 1, keysort.StringDesc(cs)}
	})
	if !reflect.DeepEqual(slice, []CustomInt{4, 2, 5, 3, 1}) {
		t.Fatal(slice)
	}
}
