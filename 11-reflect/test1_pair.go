package main

import "fmt"

/*
interface 断言机制
*/
func main() {
	var a string
	//pair<static type:string , value:abc>
	a = "abc"

	var allType interface{}
	//pair<type:string , value:abc>
	allType = a
	value, ok := allType.(string)
	if !ok {
		fmt.Println("allType is not string type")
	} else {
		fmt.Println("allType is type string,value = ", value)
	}
}
