package channel

import (
	"fmt"
	"sync"
)

//Production 生产者
func Production(ch chan int, w *sync.WaitGroup) {
	//生产完毕  go程可以关闭
	defer w.Done()
	fmt.Println("数据开始生产...")
	for i := 0; i < 100; i++ {
		fmt.Printf("生产了第%v条数据\n", i)
		ch <- i
	}
	//生产完毕 关闭通道
	close(ch)
	fmt.Println("数据生产完毕...生产go程结束...")
}
