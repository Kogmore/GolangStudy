package main

import "fmt"

/*
	channel与range
*/
func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	/*
		for {
			//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
			if data, ok := <-c; ok {
				fmt.Println("data = ", data)
			} else {
				fmt.Println("通道c数据消费完毕...")
				break
			}
		}
	*/
	//可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println("data = ", data)
	}

	fmt.Println("main 结束...")
}
