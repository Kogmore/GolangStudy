package main

import (
	"fmt"
	"reflect"
)

/*
结构体标签的作用，以及反射机构体的标签
*/

type resume struct {
	Name string `info:"name" doc:"我的名字"` //对结构体进行一个标签说明
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		fmt.Println("info:", tagInfo)

		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("doc:", tagDoc)
	}
}

func main() {
	var re resume

	findTag(&re)
}
