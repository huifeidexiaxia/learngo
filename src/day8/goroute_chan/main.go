package main

import (
	"fmt"
	"time"
)

/**
chan是阻塞的
注意点： 如果只写，不读，那么只能定义长度个数的，后面就写不进去了

如果只读，不写，也一样，当没有可读的时候，就阻塞了，类似于Java中和阻塞队列
*/
func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("put data:", i)
		time.Sleep(time.Second)
	}
}

func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
		time.Sleep(time.Second)
	}
}

func main() {
	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
}
