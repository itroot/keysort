package keysort_test

import (
	"fmt"
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

type Apple struct {
	taste int
	size  int
	good  bool
}

type Orange struct {
	taste int
	size  int
	good  bool
}

func ExampleApplesAndOranges() {
	slice := []interface{}{&Apple{1, 1, true}, &Orange{5, 0, false}}
	keysort.Sort(slice, func(i int) keysort.Sortable {
		switch v := slice[i].(type) {
		case *Apple:
			return keysort.Sequence{keysort.BoolDesc(v.good), v.taste, v.size}
		case *Orange:
			return keysort.Sequence{keysort.BoolDesc(v.good), v.taste, v.size}
		default:
			panic("Unknown type")
		}
	})
	for _, fruit := range slice {
		fmt.Printf("%#v\n", fruit)
	}
	// Output:
	// &keysort_test.Apple{taste:1, size:1, good:true}
	// &keysort_test.Orange{taste:5, size:0, good:false}
}
