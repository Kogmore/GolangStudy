package main

import (
	"fmt"
)

func main() {
	//定义一个channel
	ch := make(chan int)

	go func() {
		defer fmt.Println("goroutine 运行结束")

		fmt.Println("goroutine 正在运行...")

		ch <- 666 //将666 发送给ch
	}()

	num := <-ch //从ch中接受数据 并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 运行结束...")
}
