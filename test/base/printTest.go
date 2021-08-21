package main

import "fmt"

func main() {
	var name = "yang_zzu"
	var age = 18
	//var address = "中国"

	/**
	%T 表示数据类型
	%v 数据值
	%#v 数据的json值的
	 */

	fmt.Printf("%T, %v, %#v\n", name, name, name)
	fmt.Printf("%T, %v, %#v\n", age, age, age)
}
