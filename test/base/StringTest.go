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
	fmt.Println("name:", name) //name: yang_zzu

	var age = 18
	fmt.Println(age) //18

	var phone string
	fmt.Println(phone)

	// 在函数内使用，不能在包级别使用
	address := "中国"
	fmt.Println(address) //中国

	var (
		spuName  = "yang_zzu"
		spuAge   = 18
		spuPhone string
	)

	x, y := "1", 1

	fmt.Println(spuName, spuAge, spuPhone, x, y) //yang_zzu 18  1 1

	fmt.Println("\a") //响铃，运行到这里后会发出响声

	fmt.Print("\n") //换行

	var yuanSheng = `yang_zzu \n yang_zzu`        //不会进行换行
	fmt.Printf("%T, %#v\n", yuanSheng, yuanSheng) //string, "yang_zzu \\n yang_zzu"

	yuanSheng += " 18"
	fmt.Println(yuanSheng) //yang_zzu \n yang_zzu 18

	var noYuanSheng = "yang_zzu\nyang_zzu"            //不会进行换行
	fmt.Printf("%T, %#v\n", noYuanSheng, noYuanSheng) //string, "yang_zzu\nyang_zzu"

	// 字符串的比较会，根据char类型进行挨个比较
	fmt.Println("abc" > "abcde") //false
	fmt.Println("abc" > "bac")   //false

	//字符串切片取值,字符串中如果存在中文，获取的值是字节的数据，不是单个字符
	//字符串全部是由ascii组成的，通过切片获取的数据是，一一对应的
	fmt.Println(yuanSheng[0])        //121
	fmt.Printf("%q\n", yuanSheng[0]) //'y'
	fmt.Printf("%c\n", 121)          //y
	//字符串切片
	fmt.Println(yuanSheng[1:4]) //ang
	var yang = "杨"
	fmt.Println("-------------")
	fmt.Println(yang[0]) //230 , yang变量的第一个字节输出的 char 值为 230
	fmt.Println("-------------")
	fmt.Println(yang[0:3]) //杨
	fmt.Println("-------------")
	fmt.Printf("%c\n", 230) //æ ， 230表示的char为
	fmt.Printf("%U\n", '杨') //U+6768

	//字符串的长度， len 实际计算的是占用的字节的大小，而不是字符串的长度，
	//全是ascii组成的字符串，表示的是字符串的长度
	fmt.Println(len(yuanSheng)) //23
	fmt.Println(len(yang))      //3

}
