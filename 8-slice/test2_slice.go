package main

import (
	"fmt"
)

func printArray1(myArray []int) { //引用传递
	//_表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
	//myArray = append(myArray, 100)
}

func main() {
	//动态数组其实就是一个指针  指向底层的数组
	myArray := []int{1, 2, 3, 4} //动态数组， 切片 slice

	fmt.Printf("myArray types is %T\n", myArray)

	//传递切片给printArray1
	printArray1(myArray)
	fmt.Println("----------------")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
