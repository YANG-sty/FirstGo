package main

import "fmt"

func main() {
	var name = "yang_zzu"
	var age = 18
	//var address = "中国"

	/**
	%T 表示数据类型
	%v 自动调用每种类型的格式化输出
	%#v 数据的json值的值，实际就是在 %v 的基础上添加一些类型属性的说明，一般在调试的时候使用
	%q 显示原始的数据，
	%t 打印boolean类型数据
	%d, 转换为10进制
	%xnd, 输出数据的时候至少占n位，不够的时候 用指定的x的进行代替，默认占位符位空格
	%b, 转换为2进制
	%o, 转换为8进制
	%x, 转换为16进制
	%U, 转换为unicode编码
	%c, 转换为byte类型
	%n.mf  n表示该浮点数小数位最少有几位，m表示该浮点数小数位最多位多少位
	%e 浮点数的科学计数法
	%g 会在 %f, $e 之间自动的进行选择
	*/

	fmt.Printf("%T, %v, %#v\n", name, name, name)
	fmt.Printf("%T, %v, %#v\n", age, age, age)

	var yang = "杨"
	fmt.Printf("%v, %q\n", name, name)
	fmt.Printf("%v, %q\n", yang, yang)

	// int 类型占位符,  指定占位的个数，不够的时候默认使用空格占位，可以指定占位的符号例如0
	fmt.Printf("[%d]\n", 9)     //[9]
	fmt.Printf("[%3d]\n", 9)    //[  9]
	fmt.Printf("[%03d]\n", 9)   //[009]
	fmt.Printf("[%3d]\n", 100)  //[100]
	fmt.Printf("[%3d]\n", 9999) //[9999]

}
