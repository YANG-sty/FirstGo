package main

import "fmt"

func main() {

	var flag string
	fmt.Println("请输入（y or n)：")
	fmt.Scan(&flag)

	if flag == "y" || flag == "Y" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// switch 的方式，每个 case 在匹配到后，不用使用break ，后面的内容默认不再进行匹配操作
	switch flag {
	case "y":
		fmt.Println("2 YES")
	case "Y":
		fmt.Println("2 YES")
	case "n":
		fmt.Println("2 NO")
	case "N":
		fmt.Println("2 NO")
	default:
		fmt.Println("错误输入❌")
	}

	switch {
	case flag == "y" || flag == "Y":
		fmt.Println("3 YES")
	case flag == "n" || flag == "N":
		fmt.Println("3 NO")
	default:
		fmt.Println("3 输入错误🙅")
	}

}
