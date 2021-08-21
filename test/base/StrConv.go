package main

import (
	"fmt"
	"strconv"
)

func main() {
	var(
		intValue = 18
		float64Value = 18.00
		stringValue = "18"
	)
	fmt.Printf("%T, %#v\n", string(intValue), string(intValue))
	fmt.Printf("%T, %#v\n", int(float64Value), int(float64Value))
	fmt.Printf("%T, %#v\n", float64(intValue), float64(intValue))

	//字符串转int
	v, err := strconv.Atoi(stringValue)
	fmt.Println(v, err)
	fmt.Printf("%T, %#v\n", v, v)

	//字符串转float
	f, err := strconv.ParseFloat(stringValue, 64)
	fmt.Println(f, err)
	fmt.Printf("%T, %#v\n", f, f)

	//int转字符串
	fmt.Printf("%T, %#v\n", strconv.Itoa(65), strconv.Itoa(65))

	//float转字符串，保留指定的精度
	ff := strconv.FormatFloat(float64Value, 'f', 10, 64)
	fmt.Printf("%T, %#v\n", ff, ff)

}
