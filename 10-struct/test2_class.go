package main

/*
	实现Golang中的封装：
	把抽象的字段和对字段的操作封装在一起，数据被保护在内部。程序的其他包只能通过被授权的方式才能对其进行操作
*/

import "fmt"

//Hero 定义一个结构体
type Hero struct { //如果类名首字母大写，表示其他包也能访问
	//如果属性首字母大写，表示该属性是外部能够访问的，否则的话只有类的内部能够访问
	Name  string
	Ad    int
	Level int
}

/*
//Show 定义一个属于Hero的成员方法 获取Hero的所有属性
func (h Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.Level)
}

func (h Hero) GetName() string {
	return h.Name
}

func (h Hero) SetName(newName string) {
	//h是调用该方法的对象的一个副本（拷贝）
	h.Name = newName
}
*/

//Show 定义一个属于Hero的成员方法 获取Hero的所有属性
func (h *Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.Level)
}

//GetName 定义一个属于Hero的成员方法 获取Hero的Name属性
func (h *Hero) GetName() string {
	return h.Name
}

//SetName 定义一个属于Hero的成员方法 设置Hero的Name属性
func (h *Hero) SetName(newName string) {
	h.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{
		Name:  "zhang3",
		Ad:    100,
		Level: 1,
	}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
