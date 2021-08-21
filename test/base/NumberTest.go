package main

import "fmt"

func main() {
	// 不同的数据类型是不能进行运算的
	var age int = 18 // 默认是64位
	var age8 int8 = 18
	var age32 int32 = 18

	fmt.Println(int(age32) + age)

	fmt.Printf("%T, %#v, %d\n", age, age, age)
	fmt.Printf("%T, %#v, %d\n", age8, age8, age8)
	fmt.Printf("%T, %#v, %d\n", age32, age32, age32)

	fmt.Println(077)   // 8进制
	fmt.Println(0X123) // 16进制
	fmt.Println(78)    // 10进制
	fmt.Println(111)   // 默认是10进制

	/*
			0111
		^	0010
			0101 =》 5
		不同为1，相同为0
	*/
	fmt.Println(7 ^ 2)

	/*
			0111
		&^	0010
			0101 => 5
		不同为1，相同为0
			1111
		&^	1000
			0111 => 7
	*/
	fmt.Println(15 &^ 8)

	var (
		A            byte = 'A'
		Aint         byte = 65
		unicodePoint rune = '中' //码点 输出的时候，和 字节一样会转换成 int
	)
	fmt.Println(A == Aint)
	fmt.Println(unicodePoint)

	/**
	%d, 转换为10进制
	%b, 转换为2进制
	%o, 转换为8进制
	%x, 转换为16进制
	%U, 转换为unicode编码
	%c, 转换为byte类型
	*/
	fmt.Printf("%d, %b, %o, %x, %U, %c, %c\n", 'A', 15, 15, 15, '中', 65, 'A')

	var price float32 = 1
	fmt.Printf("%#v\n", price)
	priceHeight := 1e3
	fmt.Printf("%#v\n", priceHeight)
	var priceLow float32 = 1e3
	fmt.Printf("%#v\n", priceLow)

	var (
		f1 = 1.1
		f2 = 2.23
		f3 = 3.21
	)
	f3++
	fmt.Println(f1 + f2)
	fmt.Println(f3)

}
