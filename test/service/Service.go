package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Path: ", request.URL.Path)
	_, err := fmt.Fprintln(writer, "It works!")
	if err != nil {
		return
	}
	_, err1 := fmt.Fprintln(writer, request.URL.Path)
	if err1 != nil {
		return
	}
	// 获取url路径从下标为1开始到最后的数据
	_, err2 := fmt.Fprintln(writer, request.URL.Path[1:])
	if err2 != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}
