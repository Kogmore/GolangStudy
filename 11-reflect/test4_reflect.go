package main

import (
	"fmt"
	"reflect"
)

/*
	反射一个基础类型的 值和类型
*/

func reflectNum(arg interface{}) {
	fmt.Println("type = ", reflect.TypeOf(arg))
	fmt.Println("value = ", reflect.ValueOf(arg))
}

func main() {
	var num = 1.2345

	reflectNum(num)
}
