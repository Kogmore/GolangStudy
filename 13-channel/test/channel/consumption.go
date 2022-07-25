package channel

import (
	"fmt"
	"sync"
)

//Consumption 消费者
func Consumption(ch chan int, w *sync.WaitGroup) {
	//数据消费完毕 go程可以关闭
	defer w.Done()
	fmt.Println("开始消费数据...")
	for data := range ch {
		fmt.Printf("消费了第%v条数据，data = %v\n", data, data)
	}
	fmt.Println("数据消费完成...完结撒花！！！")
}
