package main

import "fmt"

/*
	channel与select
*/
func fibonacii(ch1, ch2 chan int) {
	x, y := 1, 1

	for {
		select {
		case ch1 <- x: //如果ch1可写，则该case就会进来
			x = y
			y = x + y
		case <-ch2:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-ch1)
		}
		ch2 <- 0
	}()

	//main go
	fibonacii(ch1, ch2)
}
