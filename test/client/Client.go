package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {

	var h hello = hello{
		endpoint: "http://localhost:9090/",
	}
	sayHello, err := h.SayHello("yang_zzu")
	if err != nil {
		println(err.Error())
		return
	}
	println(sayHello)
	//打印方法名
	PrintFuncName(h)

}

type HelloService interface {
	SayHello() (string, error)
}

type hello struct {
	endpoint string
}

func (h hello) SayHello(name string) (string, error) {
	client := http.Client{}
	response, err := client.Get(h.endpoint + name)
	//返回内容有 err 首先要判断err
	if err != nil {
		println(err.Error())
		return "", err
	}
	all, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		fmt.Printf("%+v\n", err1)
		return "", err1
	}
	fmt.Printf("%v\n", all)
	//byte 转字符串
	return string(all), nil
}

/**
反射
*/
func PrintFuncName(val interface{}) {
	typeOf := reflect.TypeOf(val)
	num := typeOf.NumMethod()
	for i := 0; i < num; i++ {
		method := typeOf.Method(i)
		fmt.Println(method.Name)
	}
}
