package main

import (
	"fmt"
)

func printArray(myArray [5]int) {
	//值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}
	myArray[0] = 369
}

func main() {
	//固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4, 5}
	myArray3 := [5]int{9, 8, 7, 6, 5}

	for i := 0; i < len(myArray1); i++ {
		fmt.Printf("myArray%v = %v\n", i, myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)

	fmt.Println("---------------")

	for index, value := range myArray3 {
		fmt.Println("index = ", index, "value = ", value)
	}
}
