package main

import "fmt"

func main() {
	//一、声明slice1是一个切片，并且初始化，默认值是1,2,3。长度是3
	slice1 := []int{1, 2, 3}
	//1 获取slice1的长度和值
	fmt.Printf("slice1 len = %d , slice1 value = %v\n", len(slice1), slice1)
	fmt.Println("------------------------------------------------------------")
	//二、声明slice2是一个切片，但是并没有给slice分配空间
	var slice2 []int
	//2 获取slice2的长度和值
	fmt.Printf("slice2 len = %d , slice2 value = %v\n", len(slice2), slice2)
	//判断一个slice2是否为0
	if slice2 == nil {
		fmt.Println("slice2是一个空切片")
	} else {
		fmt.Printf("slice2不是一个空切片，他的值是%v", slice2)
	}
	fmt.Println("------------------------------------------------------------")
	//三、声明slice3是一个切片。同时给slice分配空间，3个空间，初始化值是0
	var slice3 []int
	slice3 = make([]int, 3) //开辟三个空间，默认值是0
	slice3[0] = 100         //给slice3下标为0的值  修改为100
	//3 获取slice3的长度和值
	fmt.Printf("slice3 len = %d , slice3 value = %v\n", len(slice3), slice3)
	fmt.Println("------------------------------------------------------------")
	//四、声明slice4是一个切片。同时给slice分配空间，3个空间，初始化值是0，通过:=推导出来slice1是一个切片
	slice4 := make([]int, 3)
	//4 获取slice4的长度和值
	fmt.Printf("slice4 type of = %T , slice4 len = %d , slice4 value = %v", slice4, len(slice4), slice4)
}
