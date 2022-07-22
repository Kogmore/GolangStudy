package main

import "fmt"

func main() {
	//声明myMap1是一种map类型  kye是string类型 value是string类型
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空的Map")
	}

	//在使用Map前  需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 3)

	myMap1["one"] = "猪八戒"
	myMap1["two"] = "孙悟空"
	myMap1["three"] = "沙悟净"
	myMap1["four"] = "唐僧"

	fmt.Println("myMap1 value = ", myMap1)
	fmt.Printf("myMap1 len = %v\n", len(myMap1))

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[3] = "程序员"
	myMap2[2] = "外卖员"
	myMap2[1] = "清洁员"

	fmt.Println("myMap2 value = ", myMap2)
	fmt.Printf("myMap2 len = %v\n", len(myMap2))

	//第三种声明方式
	myMap3 := map[int]string{
		3: "嘻嘻",
		4: "呵呵",
		2: "？？",
		1: "哦哦",
	}
	fmt.Println("myMap3 value = ", myMap3)
	fmt.Printf("myMap3 len = %v\n", len(myMap3))
}
