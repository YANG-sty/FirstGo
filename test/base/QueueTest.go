package main

import "fmt"

func main() {
	//队列，先进先出

	var ids = []int{}
	ids = append(ids, 1, 2, 3)
	ids = append(ids, 4)
	fmt.Printf("%T, %#v\n", ids, ids)

	var x int
	fmt.Printf("%T, %#v\n", x, x) //x 在不赋值的时候，默认的初始值为 0
	x = ids[0]
	ids = ids[1:]                 //移除第一个元素
	fmt.Printf("%T, %#v\n", x, x) //1

	x = ids[0]
	ids = ids[1:]                 //移除第一个元素
	fmt.Printf("%T, %#v\n", x, x) //2

	x = ids[0]
	ids = ids[1:]                 //移除第一个元素
	fmt.Printf("%T, %#v\n", x, x) //3

	x = ids[0]
	ids = ids[1:]                 //移除第一个元素
	fmt.Printf("%T, %#v\n", x, x) //4

}
