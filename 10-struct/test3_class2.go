package main

/*
	实现Golang中的继承：
	继承可以解决代码复用，当结构体中存在相同的属性和方法时，可以从这些结构体中抽象出结构体，其他的结构体中不需要重新定义这些相同的属性和方法。
	如果一个结构体中嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。
*/

import "fmt"

//Human 定义一个结构体
type Human struct {
	Name string
	Sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (h *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

//===========================

//SuperMan 定义一个结构体 继承Human类
type SuperMan struct {
	Human //SuperMan类继承了Human的方法
	level int
}

//Eat 重定义父类的方法Eat（）
func (s *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//Fly 子类的新方法
func (s *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

//Print 输出SuperMan的所有属性的值
func (s *SuperMan) Print() {
	fmt.Println("SuperMan.Name = ", s.Name)
	fmt.Println("SuperMan.Sex = ", s.Sex)
	fmt.Println("SuperMan.level = ", s.level)
}

func main() {
	h := Human{
		Name: "zhang3",
		Sex:  "female",
	}

	h.Eat()
	h.Walk()

	//定义一个子类对象
	//s := SuperMan{
	//	Human: Human{Name: "li4", Sex: "female"},
	//	level: 10,
	//}
	var s SuperMan
	s.Name = "li4"
	s.Sex = "male"
	s.level = 88

	s.Eat()  //子类重写父类的方法
	s.Walk() //父类的方法
	s.Fly()  //子类的方法
	s.Print()
}
