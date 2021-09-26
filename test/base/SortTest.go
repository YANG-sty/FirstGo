package main

import (
	"fmt"
	"sort"
)

func main() {

	var ids = []int{3, 2, 31, 5, 1, 0}
	//正序
	sort.Ints(ids)
	fmt.Printf("%T, %#v\n", ids, ids)
	//逆序
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	fmt.Printf("%T, %#v\n", ids, ids)

	var names = []string{"c", "d", "a", "z", "e"}
	sort.Strings(names)
	fmt.Printf("%T, %#v\n", names, names)

	var nums = []int{1, 2, 3, 5, 10, 8}
	var serach = 30
	sort.Ints(nums)
	/**
	二分查找查找 sort.SearchInts
	1. 数据必须是有序的
	2. 正序排序
	*/
	index := sort.SearchInts(nums, serach)
	fmt.Println(index)
	fmt.Println(ids[index] == serach)

}
