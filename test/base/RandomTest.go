package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//设置随机数的种子，只需要设置一次,种子一样，每次生成的随机数也是一样的
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	}

	//生成 100 以内的随机数
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int() % 100)
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

}
