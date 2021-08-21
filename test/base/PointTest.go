package main

import "fmt"

func main() {
	//初始值，在不进行赋值的情况下，为 nil
	var pointInt *int;
	var pointString *string;

	fmt.Printf("%T, %#v\n", pointInt, pointInt)
	fmt.Printf("%T, %#v\n", pointString, pointString)

	age := 18
	fmt.Printf("%T, %#v\n", age, age)
	fmt.Printf("%T, %#v\n", &age, &age)

	fmt.Println("----------------------")

	//字符串赋值
	pointString = new(string)
	fmt.Printf("%T, %#v\n", *pointString, *pointString)


	fmt.Println("----------------------")

	//指针赋值
	pointInt = &age
	fmt.Printf("%T, %#v\n", pointInt, pointInt)
	fmt.Printf("%T, %#v\n", age, age)
	fmt.Printf("%T, %#v\n", *pointInt, *pointInt)

	fmt.Println("----------------------")

	*pointInt = 23
	fmt.Printf("%T, %#v\n", age, age)
	fmt.Printf("%T, %#v\n", *pointInt, *pointInt)

	fmt.Println("----------------------")

	// age2 是拷贝了 age 的数据，创建了一个新的内存空间
	age2 := age
	fmt.Printf("%T, %#v\n", age2, age2)

	fmt.Println("----------------------")

	// 修改 age 的值，age2 的内存空间的 数据是不会发生变化的
	age = 25
	fmt.Printf("%T, %#v\n", age, age)
	fmt.Printf("%T, %#v\n", age2, age2)
	fmt.Printf("%T, %#v\n", *pointInt, *pointInt)

	fmt.Println("----------------------")

	//指针的指针
	//ppString 指向 指针pointString的地址
	//指针的值发生改变，指针的指针的值也会发生改变
	//指针的指针 值 发生改变，指针的值也会发生改变
	ppString := &pointString
	**ppString = "yang_zzu"
	fmt.Printf("%T, %#v\n", **ppString, **ppString)
	fmt.Printf("%T, %#v\n", *pointString, *pointString)

	fmt.Println("----------------------")

	*pointString = "yang_"
	fmt.Printf("%T, %#v\n", **ppString, **ppString)
	fmt.Printf("%T, %#v\n", *pointString, *pointString)






}
