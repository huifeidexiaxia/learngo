package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	//close(ch)
	for {
		var b int

		// 如果Ok==false，说明chan 已经关掉了
		b, ok := <-ch
		if ok == false {
			fmt.Println("chan is close")
			break
		}
		fmt.Println(b)
	}

	time.Sleep(1 * time.Second)
}
