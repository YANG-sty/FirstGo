package main

import (
	"errors"
	"fmt"
)

func main() {
	values := []interface{}{0, 1, 2, 3, 5}
	newValues, err := Add(values, 4, 4)
	if err != nil {
		println(err.Error())
		return
	}
	for i, value := range newValues {
		fmt.Printf("下标: %d, 值: %d\n", i, value)
	}

}

func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {
	if index < 0 || index > len(values) {
		return nil, errors.New("index 非法")
	}
	res := []interface{}{}
	//先放好 0 1 2 3
	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}

	//插入的数据放入
	res = append(res, val)

	//放入插入位置后面的数据
	for i := index; i < len(values); i++ {
		res = append(res, values[i])
	}
	return res, nil
}
