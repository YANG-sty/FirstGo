package main

import "fmt"

func main() {

	//const修饰的是常量，不能被修改，否则会报错，定义后可以不使用
	const yang = "zzu"

	const (
		A = 123
		B // 不显性赋值的时候，会使用前面变量的数据类型和值
		C
		D = "yang_zzu"
		E
		F
	)
	fmt.Println(A, B, C, D, E, F)

	// 使用常量来定义枚举类型
	const (
		_A = iota // 常量组内第一次 iota，会初始化 0，后面 每次调用会 +1
		_B = iota
		_C = iota
		_D	//实际是省略了 iota 关键字
		_E
		_F
		_G
	)
	fmt.Println(_A, _B, _C, _D, _E, _F, _G)

}
