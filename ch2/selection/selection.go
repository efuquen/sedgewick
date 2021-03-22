package main

import (
	"fmt"
	"sort"

	"github.com/efuquen/sedgewick/stdin"
)

func Sort(a sort.Interface) {
	for i := 0; i < a.Len(); i++ {
		min := i
		for j := i + 1; j < a.Len(); j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Swap(i, min)
	}	
}

func main() {
	a := stdin.ReadAllStrings()
	Sort(sort.StringSlice(a))
	fmt.Printf("%v\n", a)
}
