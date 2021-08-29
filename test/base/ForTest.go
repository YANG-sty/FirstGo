package main

import "fmt"

func main() {

	var (
		id  = 0
		sum = 0
	)
	for {
		sum = id + sum
		id++
		if id > 100 {
			break
		}
	}
	fmt.Println(sum)

	fmt.Println("--------------------")

	var index int = 0
	for index < 10 {
		fmt.Println(index)
		index++
	}

	fmt.Println("--------------------")

	var content string = "ABCDEF"

	//方式一
	for i := 0; i < len(content); i++ {
		fmt.Printf("%c\n", content[i])
	}
	fmt.Println("--------------------")
	//等同于方式一
	for i := range content {
		fmt.Println(i)
	}
	fmt.Println("--------------------")
	//推荐使用这中方式
	for i, i2 := range content {
		fmt.Printf("%#v, %c\n", i, i2)
	}
	fmt.Println("--------------------")
	//推荐使用这种方式，进行遍历，因为
	var msg string = "我是yang_zzu"
	for i, value := range msg {
		fmt.Printf("%#v, %#v, %U, %q\n", i, value, value, value)
	}
	fmt.Println("--------------------")
	// 当不用某个变量的时候 可以使用 _ 代替
	for _, value := range msg {
		fmt.Printf("%q", value)
	}
	fmt.Println()
	fmt.Println("--------------------")
	//这种方式获得的脚标，表示的是字符串切片的位置，通过位置去获得中文是不正确的，中文占用的不是一个字节，英文占用的是一个字节
	for i := range msg {
		fmt.Printf("%#v, %#v, %U, %q\n", i, msg[i], msg[i], msg[i])
	}

}
