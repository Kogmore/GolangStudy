package main

import (
	"fmt"
	"reflect"
)

/*
反射一个复杂类型（结构体）的值、类型、成员方法
*/

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is Called...")
	fmt.Printf("user = %v\n", u)
}

func main() {
	user := User{
		Id:   1,
		Name: "Aceld",
		Age:  18,
	}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input is Type = ", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input is Value = ", inputValue)
	//通过type获取里面的字段
	//1、获取interface的reflect.TypeOf，通过Type得到NumField，进行遍历
	//2、得到每个field，数据类型
	//3、通过field有一个Interface()方法得到对应的Value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fmt.Printf("field%v = %v\n", i, field)
		value := inputValue.Field(i).Interface()
		fmt.Printf("value%v = %v\n", i, value)
	}

	//通过type获取里面的方法，调用
	//注意获取方法名的时候  如果成员方法他是指针类型  0 = inputType.NumMethod()
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}
