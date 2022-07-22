package main

import "fmt"

//func swap(a int, b int) {
//	var temp int
//	temp = a
//	a = b
//	b = temp
//}

func swap(as *int, bs *int) {
	var temp int
	temp = *as //temp = main::a
	*as = *bs  //main::a = main::b
	*bs = temp //main::b = temp
}

func swap2(pp **int) {
	**pp = 200
}

func main() {
	var a int = 10
	var b int = 20

	//调用swap调换a，b的值
	swap(&a, &b)

	fmt.Println("a值 = ", a, "b值 = ", b)

	var p *int
	p = &a

	fmt.Println("a地址=", &a)
	fmt.Println("p值=", p)

	var pp **int //二级指针
	pp = &p

	fmt.Println("p地址=", &p)
	fmt.Println("pp值=", pp)

	//调用swap2修改a的值
	swap2(&p)
	fmt.Println("最终a的值=", a)
}
