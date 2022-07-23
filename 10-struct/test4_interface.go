package main

import "fmt"

/*
	实现Golang中的多态：
	变量具有多种形态，多态特征是通过接口来实现的。
*/

//AnimalIF 定义一个interface类  包含三个方法
type AnimalIF interface { //interface 本质是一个指针
	Sleep(animal string) //动物睡觉
	GetColor() string    //获取动物的颜色
	GetType() string     //获取动物的种类
}

//Cat 定义一个动物
type Cat struct {
	Color string
}

//Sleep Cat实现AnimalIF中方法
func (c *Cat) Sleep(animal string) {
	fmt.Printf("%v is Sleep\n", animal)
}

//GetColor Cat实现AnimalIF中方法
func (c *Cat) GetColor() string {
	return c.Color
}

//GetType Cat实现AnimalIF中方法
func (c *Cat) GetType() string {
	return "Cat"
}

//Dog 定义一个动物
type Dog struct {
	Color string
}

//Sleep Dog实现AnimalIF中方法
func (d *Dog) Sleep(animal string) {
	fmt.Printf("%v is Sleep\n", animal)
}

//GetColor Dog实现AnimalIF中方法
func (d *Dog) GetColor() string {
	return d.Color
}

//GetType Dog实现AnimalIF中方法
func (d *Dog) GetType() string {
	return "Dog"
}

//ShowAnimal 定义一个方法 输出动物的种类、颜色、睡觉
func ShowAnimal(animal AnimalIF) {
	//获取动物种类
	kind := animal.GetType()
	fmt.Println("Animal kind is", kind)
	//知道是什么动物在睡觉
	animal.Sleep(kind)
	//获取动物的颜色
	color := animal.GetColor()
	fmt.Printf("%v Color is %v\n", kind, color)
}

func main() {
	/*
		//接口的数据类型  父类的指针
		var animal AnimalIF
		//animal对象获取了动物猫
		animal = &Cat{Color: "Green"}
		//获取动物种类
		kind := animal.GetType()
		fmt.Println("Animal Type is", kind)
		//输出是什么种类的动物在睡觉
		animal.Sleep(kind) //调用的是Cat的Sleep()方法
		//获取动物的颜色
		color := animal.GetColor()
		fmt.Printf("%v Color is %v\n", tp, color)

		//animal对象获取了动物狗
		animal = &Dog{Color: "Red"}
		//获取动物种类
		kind = animal.GetType()
		fmt.Println("Animal Type is", kind)
		//输出是什么种类的动物在睡觉
		animal.Sleep(kind) //调用的是Cat的Sleep()方法
		//获取动物的颜色
		color = animal.GetColor()
		fmt.Printf("%v Color is %v\n", tp, color)
	*/
	//定义一个cat对象
	cat := &Cat{Color: "Green"}
	ShowAnimal(cat)
	//定义一个cat对象
	dog := &Dog{Color: "Red"}
	ShowAnimal(dog)
}
