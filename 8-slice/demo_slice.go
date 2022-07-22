package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := []int{}
	fmt.Printf("main slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("main slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
	slice = append(slice, 6)
	fmt.Printf("main slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
	slice = append(slice, 7)
	fmt.Printf("main slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
	ap(slice)
	fmt.Println("---------------------")
	fmt.Printf("main slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
}

func ap(slice []int) {
	fmt.Printf("ap slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
	slice[0] = 10
	slice = append(slice, 10)
	fmt.Printf("ap slice:%v &slice:%p unsafe.Pointer:%p  slice:%p len:%d cap:%v\n", slice, &slice, unsafe.Pointer(&slice), slice, len(slice), cap(slice))
}
