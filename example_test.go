package keysort_test

import (
	"fmt"

	"github.com/itroot/keysort"
)

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

func Example() {
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
