package main

import "fmt"

func main() {
	//定义一个slice  numbers
	numbers := make([]int, 3)

	numbers = append(numbers, 1, 5, 6, 7)

	fmt.Printf("numbers len = %d , numbers cap = %d,numbers value = %v\n", len(numbers), cap(numbers), numbers)
}
