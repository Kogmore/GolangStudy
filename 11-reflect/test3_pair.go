package main

import "fmt"

/*
详细介绍pair的作用
*/

type Reader interface {
	ReaderBook()
}

type Writer interface {
	WriterBook()
}

//Book 具体类型
type Book struct {
}

func (b *Book) ReaderBook() {
	fmt.Println("Read a Book")
}

func (b *Book) WriterBook() {
	fmt.Println("Write a Book")
}
func main() {
	//book:<type:Book , value:&Book{}>
	book := &Book{}
	//r:<type: , value:>
	var r Reader
	//book:<type:Book , value:&Book{}>
	r = book

	r.ReaderBook()

	var w Writer
	//w:<type:Book , value:&Book{}>
	w = r.(Writer) //此处的断言为何能成功？因为w 和 r具体的type是一致的

	w.WriterBook()
}
