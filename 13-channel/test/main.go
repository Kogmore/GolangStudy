package main

import (
	"GolangStudy/13-channel/test/channel"
	"sync"
)

/*
	生产与消费
	1、定义通道用于生产和消费的通信
	2、生产100条数据 塞入ch通道中 即可关闭通道
	3、消费ch通道中的所有数据  直至通道关闭
*/
//创建线程等待
var w sync.WaitGroup

func main() {
	//定义一个通道ch
	ch := make(chan int, 10)
	//创建计数器
	w.Add(2)
	//生产者
	go channel.Production(ch, &w)
	//消费者
	go channel.Consumption(ch, &w)
	//等待计数器为0
	w.Wait()
}
