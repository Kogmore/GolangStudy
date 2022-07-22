package main

import "fmt"

func printMap(cityMap map[string]string) {
	fmt.Println("-------------------------")
	//cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Printf("key = %v , value = %v\n", key, value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap1 := make(map[string]string)

	//添加键值对
	cityMap1["China"] = "BeiJing"
	cityMap1["Japan"] = "Tokyo"
	cityMap1["USA"] = "NewYork"

	//遍历Map
	printMap(cityMap1)
	//删除键值对
	delete(cityMap1, "China")

	//修改
	cityMap1["USA"] = "DC"
	ChangeValue(cityMap1)

	//遍历Map
	printMap(cityMap1)
}
