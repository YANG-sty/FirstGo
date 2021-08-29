package main

import (
	"fmt"
)

func main() {

	var id [5]float32
	var name [5]string
	var age [5]int

	fmt.Printf("%T\n", id)
	fmt.Printf("%#v\n", id) // 数据的初始值 不是 nil 而是，数组的数据类型的 零值 [5]float32{0, 0, 0, 0, 0}

	fmt.Printf("%T\n", name)
	fmt.Printf("%#v\n", name) //[5]string{"", "", "", "", ""}

	fmt.Printf("%T\n", age)
	fmt.Printf("%#v\n", age) //[5]int{0, 0, 0, 0, 0}

	fmt.Println("---------------------")

	//方式一
	//字面量赋值，方式一
	id = [5]float32{111111.0, 2222222.0, 3333333.0, 4444444, 55555555}
	name = [5]string{"张三", "李四", "王五", "赵六", "孙七"}
	fmt.Printf("%#v\n", id)   //[5]float32{111111, 2.222222e+06, 3.333333e+06, 4.444444e+06, 5.5555556e+07}
	fmt.Printf("%#v\n", name) //[5]string{"张三", "李四", "王五", "赵六", "孙七"}

	fmt.Println("---------------------")

	//方式二
	//短声明赋值，省略数组的大小，会根据后面的值的个数自动的进行赋值操作
	scores := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", scores) //[5]int
	fmt.Printf("%#v\n", scores)

	fmt.Println("---------------------")

	//方式三
	//赋值，通过数组的索引进行赋值，没有指定的下标索引会使用该数组类型的零值代替
	//[5]int{0: 100} 会产生一个新的数组，该数组赋值给 scores
	scores = [5]int{0: 100}
	fmt.Printf("%T\n", scores)  //[5]int
	fmt.Printf("%#v\n", scores) //[5]int{100, 0, 0, 0, 0}

	fmt.Println("---------------------")

	//关系运算
	phone := [...]string{"A", "B", "C"}
	address := [...]string{"A", "B", "C"}

	fmt.Println(phone == address) // true

	email := [...]string{"A", "B", "B"}
	fmt.Println(phone == email) //false

	fmt.Println("-----------元素操作----------")

	phone[0] = "a"
	fmt.Printf("%#v\n", phone)

	fmt.Println("-----------遍历----------")

	for i := range phone {
		fmt.Print(phone[i])
	}

	fmt.Println()
	for _, value := range phone {
		fmt.Print(value)
	}

	fmt.Println()
	fmt.Println("-----------二维数组----------")

	var dd [3][2]int
	dd[0][0] = 100
	dd[2][1] = 999
	fmt.Printf("%#v\n", dd) //[3][2]int{[2]int{100, 0}, [2]int{0, 0}, [2]int{0, 999}}

	fmt.Println("-----------多维数组----------")

	ddd := [3][2]int{0: [2]int{1, 2}, 1: [2]int{100, 200}, 2: [2]int{99, 999}}
	fmt.Printf("%#v\n", ddd)
	ddd[1][0] = 111
	fmt.Printf("%#v\n", ddd)

	/*dddd := [3][2][2]int{
		0: [2]int{0: [2]int{1,2}, 1: [2]int{1,2}},
		1: [2]int{0: [2]int{1,2}, 1: [2]int{1,2}},
		2: [2]int{0: [2]int{1,2}, 1: [2]int{1,2}}
	}*/

}
