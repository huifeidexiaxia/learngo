package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		ch <- i
	}
	// 为什么这里不关闭会报错？
	//close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	time.Sleep(1 * 1000 * 10 * time.Microsecond)
}
