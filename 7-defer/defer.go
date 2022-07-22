package main

import "fmt"

func main() {
	//写入defer关键字
	defer fmt.Println("main and 1")
	defer fmt.Println("main and 2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
