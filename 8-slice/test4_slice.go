package main

import "fmt"

func main() {
	//定义一个slice  numbers
	numbers := make([]int, 3, 5)
	fmt.Printf("numbers len = %d , numbers cap = %d,numbers value = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1，numbers len=4，【0,0,0,1】，cap=5
	numbers = append(numbers, 1)
	fmt.Printf("numbers len = %d , numbers cap = %d,numbers value = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素2，numbers len=5，[0,0,0,1,5]，cap=5
	numbers = append(numbers, 2)
	fmt.Printf("numbers len = %d , numbers cap = %d,numbers value = %v\n", len(numbers), cap(numbers), numbers)

	//向一个cap已经满的slice追加元素,第一次会扩容原有容量的一倍  比如原有为5 则扩容5  第二次扩容则会一直保持为偶数
	numbers = append(numbers, 3)
	fmt.Printf("numbers len = %d , numbers cap = %d,numbers value = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("---------------------------------------")
	//定义一个slice  numbers2  如不设置容量  则容量保持跟长度一致  扩容规律同上
	numbers2 := make([]int, 3)
	fmt.Printf("numbers2 len = %d , numbers2 cap = %d,numbers2 value = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1, 2, 3, 4)
	fmt.Printf("numbers2 len = %d , numbers2 cap = %d,numbers2 value = %v\n", len(numbers2), cap(numbers2), numbers2)
}

/*
切片的长度表示  左右指针的之间的距离  也就是首部指针和尾部指针之间
*/
