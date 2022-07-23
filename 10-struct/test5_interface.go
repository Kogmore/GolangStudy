package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println("arg value is = ", arg)

	//interface{}应该如何区分 此时引用的底层数据类型到底是什么呢？

	//给interface{} 提供了一种“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Printf("arg is type = %T\n", value)
	}
}

type Book1 struct {
	auth string
}

func main() {
	book1 := Book1{auth: "GoLang"}
	myFunc(book1)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
