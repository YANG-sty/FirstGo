package main

import "fmt"

//切片

func main() {

	var id []int                    //nil 表示，切片没有初始化
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int(nil) 切片的零值

	id = []int{}                    //空切片，表示，切片已经初始化
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{}

	id = []int{100, 200, 300}
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{100, 200, 300}

	id = []int{0: 100, 10: 999}
	fmt.Printf("%#v\n", id) //[]int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 999}

	fmt.Println("-----------make函数,赋值-------------")

	//申请5个int类型元素的切片
	id = make([]int, 5)
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{0, 0, 0, 0, 0}

	//申请一个，容量是10的数组，切片元素的个数为0
	//id数据能操作的元素的个数为 0 个，数组的长度为10，当超过10 个元素的时候，数组进行扩容
	id = make([]int, 0, 10)
	//fmt.Printf("%T, %#v\n", id[0], id[0]) //切片元素的个数为0，id[0] 报错，数组越界
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{}

	fmt.Println("-----------make函数, 查看数组长度，和容量-------------")

	id = make([]int, 5) //不指定数组的容量的时候，默认和数组的长度一样
	id[4] = 40
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{0, 0, 0, 0, 40}

	fmt.Println(len(id), cap(id)) //5 5

	fmt.Println("-----------make函数, 添加元素-------------")

	//数组切片个数，等于数组的容量的时候，会进行扩容操作
	id = append(id, 50)
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{0, 0, 0, 0, 40, 50}
	fmt.Println(len(id), cap(id))   //6 10

	id = make([]int, 5, 6)
	//赋值
	id[4] = 4
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{0, 0, 0, 0, 4}
	fmt.Println(len(id), cap(id))   //5 6
	//增加元素
	id = append(id, 5)
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{0, 0, 0, 0, 4, 5}
	fmt.Println(len(id), cap(id))   //6 6
	//增加元素
	id = append(id, 6)
	fmt.Printf("%T, %#v\n", id, id) //[]int, []int{0, 0, 0, 0, 4, 5, 6}
	fmt.Println(len(id), cap(id))   //7 12

	fmt.Println("-----------遍历-------------")

	for _, value := range id {
		fmt.Print(value)
	}
	fmt.Println()

	fmt.Println("-----------copy切片之间的赋值-------------")

	aSlice := []int{1, 2, 3}
	bSlice := []int{4, 5, 6, 7}
	//复制的时候，当数据源的长度比目的 长的时候，多余的数据不会进行复制
	//复制的时候，源数据长度少，目的数据比源数据多的部分保持原样
	copy(aSlice, bSlice)
	fmt.Printf("%#v, %#v\n", aSlice, bSlice) //[]int{4, 5, 6}, []int{4, 5, 6, 7}

	fmt.Println("-----------切片操作-------------")

	//字符串切片操作
	var name string = "qwerty"
	ert := name[2:5] //范围 [start, end)
	fmt.Println(ert)

	//数组切片
	ids := []int{0, 1, 2, 3, 4, 5}
	// start <= end <= cap
	// 新生成的数组的 new_cap = cap - start
	// 新生成的数据，和 原来的数组 使用的是 同一个内存空间，当 idChild 的元素发生改变的时候，原数组的元素也会发生改变
	idChild := ids[1:3] //范围 [start, end)
	fmt.Printf("%T, %#v\n", idChild, idChild)
	fmt.Println(len(ids), cap(ids))
	fmt.Println(len(idChild), cap(idChild))

	fmt.Println("-----------切片操作,公用内存空间-------------")

	fmt.Printf("%T, %#v\n", ids, ids)
	fmt.Printf("%T, %#v\n", idChild, idChild)
	// 新生成的数据，和 原来的数组 使用的是 同一个内存空间，当 idChild 的元素发生改变的时候，原数组的元素也会发生改变
	idChild = append(idChild, 300)
	fmt.Printf("%T, %#v\n", ids, ids)
	fmt.Printf("%T, %#v\n", idChild, idChild)

	fmt.Println("-----------切片操作,指定数组的容量,公用内存空间-------------")

	//指定切片的容量
	// start <= end <= new_cap  new_cap = end - start
	//idChild 的容量是2，数组的长度也是2，在新增一个元素之后，数组容量不够，数组进行扩容，
	// 扩容会产生新的内存空间，idChild添加的新的元素是在新的内存空间，数据不会对原数据造成影响
	idChild = ids[1:3:3] // start:end:new_cap
	fmt.Printf("%T, %#v\n", ids, ids)
	fmt.Printf("%T, %#v\n", idChild, idChild)
	fmt.Println(len(idChild), cap(idChild))
	//添加新元素，触发扩容
	idChild = append(idChild, 1000)
	fmt.Printf("%T, %#v\n", ids, ids)
	fmt.Printf("%T, %#v\n", idChild, idChild)
	fmt.Println(len(idChild), cap(idChild))

	fmt.Println("-----------切片操作,移除元素-------------")

	fmt.Printf("%T, %#v\n", ids, ids) //[]int, []int{0, 1, 2, 300, 4, 5}
	//移除元素，相当于对数组进行切片操作
	//ids[start, end], start==0 的时候可以省略不写，end==len(ids) 的时候可以省略不写
	//表示 从索引下标为1的元素开始 到 数组最后 进行切片
	//移除第一个元素
	ids = ids[1:]
	fmt.Printf("%T, %#v\n", ids, ids) //[]int, []int{1, 2, 300, 4, 5}

	//移除最后一个元素
	ids = ids[:len(ids)-1]
	fmt.Printf("%T, %#v\n", ids, ids) //[]int, []int{1, 2, 300, 4}

	//移除中间元素
	/**
	ids: 1 2 300 4
	切片1：2 300 4	目的
	切片2：300 4 	源
	copy: 300 4 4
	因为copy 的结果和 ids 使用的是同一个内存空间，所以
	ids: 1 300 4 4
	*/
	copy(ids[1:], ids[2:])
	fmt.Printf("%T, %#v\n", ids, ids) //[]int, []int{1, 300, 4, 4}
	//移除最后一个重复的元素
	ids = ids[:len(ids)-1]
	fmt.Printf("%T, %#v\n", ids, ids) //[]int, []int{1, 300, 4}
	//移除第几个元素，就使用该元素的下表进行copy操作，然后移除最后一个元素

}
