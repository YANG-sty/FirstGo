package main

import "fmt"

func main() {
	//初始值，在不进行赋值的情况下，为 nil
	var pointInt *int
	var pointString *string

	fmt.Printf("%T, %#v\n", pointInt, pointInt)       //*int, (*int)(nil)
	fmt.Printf("%T, %#v\n", pointString, pointString) //*string, (*string)(nil)

	age := 18
	fmt.Printf("%T, %#v\n", age, age)   //int, 18
	fmt.Printf("%T, %#v\n", &age, &age) //*int, (*int)(0xc000132020)

	fmt.Println("----------------------")

	//字符串赋值
	pointString = new(string)
	fmt.Printf("%T, %#v\n", *pointString, *pointString) //string, ""

	fmt.Println("----------------------")

	//指针赋值
	pointInt = &age
	fmt.Printf("%T, %#v\n", pointInt, pointInt)   //*int, (*int)(0xc000132020)
	fmt.Printf("%T, %#v\n", age, age)             //int, 18
	fmt.Printf("%T, %#v\n", *pointInt, *pointInt) //int, 18

	fmt.Println("----------------------")

	*pointInt = 23
	fmt.Printf("%T, %#v\n", age, age)             ////int, 23
	fmt.Printf("%T, %#v\n", *pointInt, *pointInt) ////int, 23

	fmt.Println("----------------------")

	// age2 是拷贝了 age 的数据，创建了一个新的内存空间
	age2 := age
	fmt.Printf("%T, %#v\n", age2, age2) ////int, 23

	fmt.Println("----------------------")

	// 修改 age 的值，age2 内存空间的数据是不会发生变化的
	age = 25
	fmt.Printf("%T, %#v\n", age, age)             //int, 25
	fmt.Printf("%T, %#v\n", age2, age2)           //int, 23
	fmt.Printf("%T, %#v\n", *pointInt, *pointInt) //int, 25

	fmt.Println("----------------------")

	//指针的指针
	//ppString 指向 指针pointString的地址
	//指针的值发生改变，指针的指针的值也会发生改变
	//指针的指针 值 发生改变，指针的值也会发生改变
	ppString := &pointString
	**ppString = "yang_zzu"
	fmt.Printf("%T, %#v\n", **ppString, **ppString)     //string, "yang_zzu"
	fmt.Printf("%T, %#v\n", *pointString, *pointString) //string, "yang_zzu"

	fmt.Println("----------------------")

	*pointString = "yang_"
	fmt.Printf("%T, %#v\n", **ppString, **ppString)     //string, "yang_"
	fmt.Printf("%T, %#v\n", *pointString, *pointString) //string, "yang_"

}
