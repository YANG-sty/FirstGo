package main

import "fmt"

/*
goto
break
标签
*/
func main() {

	fmt.Println("1")
	goto FIRST //FIRST 表示的 goto 要跳转的位置的标签
	fmt.Println("2")
	fmt.Println("3")
FIRST: //标签后面需要跟 :
	fmt.Println("4")

	//执行结果为
	/*
		1
		4
	*/

	index := 0
	total := 0
START:
	index += 1
	total += index
	if index == 100 {
		goto END
	}
	goto START
END:
	fmt.Println(total)

BREAK:
	for i := 0; i < 5; i++ {
		fmt.Println("i : ", i, "---------------")
		for j := 0; j < 3; j++ {
			fmt.Println("j : ", j)
			if j == 2 {
				break BREAK //直接跳出多级循环
			}
		}
	}

}
