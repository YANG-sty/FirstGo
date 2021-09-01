package main

import (
	"fmt"
	"sort"
)

func main() {

	var ids = []int{3, 2, 31, 5, 1, 0}
	sort.Ints(ids)
	fmt.Printf("%T, %#v\n", ids, ids)
	//逆序
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	fmt.Printf("%T, %#v\n", ids, ids)

	var names = []string{"c", "d", "a", "z", "e"}
	sort.Strings(names)
	fmt.Printf("%T, %#v\n", names, names)

}
