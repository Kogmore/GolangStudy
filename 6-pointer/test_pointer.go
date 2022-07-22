package main

import "fmt"

func changeValue(b *int) {

	fmt.Println(b)
	*b = 10
}

func main() {
	a := 1
	fmt.Println(&a)
	changeValue(&a)

	fmt.Println(a)

	var c *int
	fmt.Println("c 的默认值=", c)
}
