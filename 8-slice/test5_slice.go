package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len = 3,cap = 3

	//[0,1]
	s1 := s[0:2] //[1,2]  不要下标为2的值
	fmt.Println("s1=", s1)
	ss2 := s[0:] //[1,2,3]  包括本身 以及之后的所有元素值
	fmt.Println("ss2=", ss2)
	sss3 := s[:2] //[1,2] 不包括本身 只截取下标前面的所有元素值
	fmt.Println("sss3=", sss3)

	s1[0] = 100
	//截取操作属于引用传递  底层用的是一个数组
	fmt.Println("s=", s)
	fmt.Println("s1=", s1)
	fmt.Println("ss2=", ss2)
	fmt.Println("sss3=", sss3)

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2=[0,0,0]

	//将s中的值拷贝到s2中
	copy(s2, s)
	//修改s的值
	s[1] = 200
	//s2值未改变  copy关键字属于值拷贝
	fmt.Println(s)
	fmt.Println(s2)
}
