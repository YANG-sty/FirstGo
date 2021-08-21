package main

import "fmt"

// 函数外定义的变量，可以不用进行引用
var version string = "1.0.0"

func main() {

	// 函数内定义的变量，必须要被使用，否则会报错

	//const修饰的是常量，不能被修改，否则会报错，定义后可以不使用
	const id = 123
	fmt.Println(id)

	var name string = "yang_zzu"
	fmt.Println("name:", name)

	var age = 18
	fmt.Println(age)

	var phone string
	fmt.Println(phone)

	// 在函数内使用，不能在包级别使用
	address := "中国"
	fmt.Println(address)

	var (
		spuName  = "yang_zzu"
		spuAge   = 18
		spuPhone string
	)

	x, y := "1", 1

	fmt.Println(spuName, spuAge, spuPhone, x, y)


	fmt.Println("\a") //响铃，运行到这里后会发出响声

	fmt.Print("\n") //换行

	var yuanSheng = `yang_zzu \n yang_zzu` //不会进行换行
	fmt.Printf("%T, %#v\n", yuanSheng, yuanSheng)

	yuanSheng += " 18"
	fmt.Println(yuanSheng)

	var noYuanSheng = "yang_zzu\nyang_zzu" //不会进行换行
	fmt.Printf("%T, %#v\n", noYuanSheng, noYuanSheng)

	// 字符串的比较会，根据char类型进行挨个比较
	fmt.Println("abc" > "abcde")
	fmt.Println("abc" > "bac")

	//字符串切片取值,字符串中如果存在中文，获取的值是字节的数据，不是单个字符
	//字符串全部是由ascii组成的，通过切片获取的数据是，一一对应的
	fmt.Println(yuanSheng[0])
	fmt.Printf("%c\n", 121)
	fmt.Println(yuanSheng[1:4])
	var yang = "杨"
	fmt.Println(yang[0])
	fmt.Printf("%c\n", 232)
	fmt.Printf("%U\n", '杨')

	//字符串的长度， len 实际计算的是占用的字节的大小，而不是字符串的长度，
	//全是ascii组成的字符串，表示的是字符串的长度
	fmt.Println(len(yuanSheng))
	fmt.Println(len(yang))






}
