package main

import "fmt"

//声明一种新的数据类型 myint，是int 的一个别名
type myint int

//Book 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	//传递一个Book的副本
	book.title = "法外狂徒"
	book.auth = "张三"
}

func changeBook2(book *Book) {
	//指针传递
	book.title = "法外狂徒"
	book.auth = "张三"
}
func main() {
	/*
		var a myint = 10
		fmt.Printf("type of a = %T\n", a)
	*/
	var book1 Book
	book1.title = "GOLang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
