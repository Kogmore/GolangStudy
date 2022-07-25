package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		//用go创建承载一个形参为空，返回值为空的函数
		go func() {
			defer fmt.Println("A.defer")

			func() {
				defer fmt.Println("b.defer")
				//退出当前的goroutine
				runtime.Goexit()	//终止当前的goroutine
				fmt.Println("B")
			}()

			fmt.Println("A")
		}()
	*/

	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)
	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
