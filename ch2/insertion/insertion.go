package main

import (
	"fmt"
	"sort"

	"github.com/efuquen/sedgewick/stdin"
)

func Sort(a sort.Interface) {
	for i := 1; i < a.Len(); i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j, j-1)
		}
	}
}

func main() {
	a := stdin.ReadAllStrings()
	Sort(sort.StringSlice(a))
	fmt.Printf("%v\n", a)
}
