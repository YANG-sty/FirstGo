package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入内容：")
	fmt.Scan(&name)
	fmt.Print("输入的内容：", name)
}
