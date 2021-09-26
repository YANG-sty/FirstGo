package main

import "fmt"

func main() {
	/*
		nil 切片, 等价余 null
		empty 切片，进行了初始化，但是没有赋值
	*/
	var nilSlice []int
	var emptySice []int = []int{}

	fmt.Printf("%T, %#v\n", nilSlice, nilSlice)
	fmt.Printf("%T, %#v\n", emptySice, emptySice)

	nilSlice = append(nilSlice, 100)
	emptySice = append(emptySice, 200)

	fmt.Printf("%T, %#v\n", nilSlice, nilSlice)
	fmt.Printf("%T, %#v\n", emptySice, emptySice)

}
