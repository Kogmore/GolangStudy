package main

import "fmt"

/*
	知识点二：defer和return 谁先谁后
*/

func deferFunc() int {
	fmt.Println("defer func call end...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func call end...")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
