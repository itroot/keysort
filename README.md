# Keysort - go module to sort slices by key

This module introduce function `keysort.Sort` that enables sort slice by key, like it's done in python `sorted` function.

## Usage

```golang
package main

import (
	"fmt"

	"github.com/itroot/keysort"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	keysort.Sort(slice, func(i int) []interface{} {
		return []interface{}{slice[i]%2 == 1, slice[i]}
	})
	fmt.Println(slice)
}
```

will output `[2 4 1 3 5]`, first all even numbers in sorted order, and then all odd numbers in sorted order

## Other examples
https://stackoverflow.com/a/63373689/1586620


