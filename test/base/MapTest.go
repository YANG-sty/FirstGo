package main

import "fmt"

func main() {
	//未进行初始化， key string ,value 浮点型
	var scores map[string]float32

	fmt.Printf("%T, %#v\n", scores, scores)

	//空map， 进行初始化，等价于 make
	scores = map[string]float32{}
	fmt.Printf("%T, %#v\n", scores, scores)

	//赋值操作
	scores = map[string]float32{"zhangSan": 60, "liSi": 45, "wangWu": 79}
	fmt.Printf("%T, %#v\n", scores, scores)

	//make 初始化变量 等价于 scores = map[string]float32{}
	scores = make(map[string]float32)
	fmt.Printf("%T, %#v\n", scores, scores)

	fmt.Println("-----------------map长度-------------------")
	scores = map[string]float32{"Y": 80, "A": 81, "N": 90, "G": 91}
	fmt.Println(len(scores))

	fmt.Println("-----------------map操作-------------------")

	fmt.Println("--------查找----------")
	//查找
	Y := scores["Y"]
	fmt.Printf("%T, %#v\n", Y, Y)
	//查找的数据不存在的情况下，返回的数据是该value数据类型的零值（初始值）
	yang := scores["yang"]
	fmt.Printf("%T, %#v\n", yang, yang)

	fmt.Println("--------判断元素是否存在----------")
	f, exit := scores["zzu"]
	//返回的value值
	fmt.Printf("%T, %#v\n", f, f)
	//判断该key值，是否存在
	fmt.Printf("%T, %#v\n", exit, exit)

	fmt.Println("--------修改操作----------")
	//新增
	//修改
	//当，修改的key值，不存在的时候，会进行新增操作
	scores["yang"] = 100
	f2, exit2 := scores["yang"]
	fmt.Printf("%T, %#v\n", f2, f2)
	fmt.Printf("%T, %#v\n", exit2, exit2)

	//删除
	fmt.Println("--------删除操作----------")
	fmt.Println(scores)
	delete(scores, "yang")
	fmt.Println(scores)

	fmt.Println("--------遍历元素----------")
	for key, value := range scores {
		//key 表示 key 值， value 表示元素的值
		fmt.Printf("%T, %#v\n", key, key)
		fmt.Printf("%T, %#v\n", value, value)

	}

	fmt.Println("-----------------nil操作-------------------")

	//nil映射，避免使用，建议进行初始化
	var varNil map[string]float32
	fmt.Println(len(varNil))
	//nil可以进行元素的获取，但是 value 的值 是 零值
	fmt.Println(varNil["Y"])
	//nili映射不能进行赋值操作
	varNil["Y"] = 90
	fmt.Println(varNil["Y"])

}
