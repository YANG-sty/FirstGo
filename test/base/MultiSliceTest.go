package main

import "fmt"

func main() {

	//多维切片
	multi := [][]int{}

	fmt.Printf("%T, %#v\n", multi, multi) //[][]int, [][]int{}

	//切片第一个元素赋值
	multi = append(multi, []int{1, 2, 3})
	//切片第二个元素赋值
	multi = append(multi, []int{7, 8, 9})

	fmt.Printf("%T, %#v\n", multi, multi) //[][]int, [][]int{[]int{1, 2, 3}, []int{7, 8, 9}}

	//第二个切片的，第三个元素
	i12 := multi[1][2]
	fmt.Printf("%T, %#v\n", i12, i12) //int, 9

	//切片赋值
	fmt.Println("------------------------切片赋值---------------------")
	//第一个切片的，第一个元素赋值为 100
	multi[0][0] = 100
	fmt.Printf("%T, %#v\n", multi, multi) //[][]int, [][]int{[]int{100, 2, 3}, []int{7, 8, 9}}

	//第二哥切片，增加一个元素200
	multi[1] = append(multi[1], 200)
	fmt.Printf("%T, %#v\n", multi, multi) //[][]int, [][]int{[]int{100, 2, 3}, []int{7, 8, 9, 200}}

}
