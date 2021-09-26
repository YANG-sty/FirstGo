package main

import "fmt"

func main() {
	//
	var sourceString string = "123abcxyzuvw杨456abc"
	var target map[rune]int = map[rune]int{}

	for _, value := range sourceString {
		//过滤掉，数字 和 中文
		if value >= 'a' && value < 'z' || value >= 'A' && value <= 'Z' {
			//map集合，不存在的元素的初始值，是该类型数据的零值
			//target[value] = target[value] + 1
			target[value]++
		}
	}

	fmt.Println("------------------打印出来的是数字unicode编码------------------")
	fmt.Println(target)

	fmt.Println("------------------打印出字符串------------------")

	for key, value := range target {
		fmt.Printf("%c : %d\n", key, value)
	}

}
