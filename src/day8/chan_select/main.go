package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		var i int
		for {
			time.Sleep(time.Second * 2)
			ch <- i
			time.Sleep(time.Second * 3)
			ch2 <- i * i

			i++
		}
	}()

	for i := 0; i < 3; i++ {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}

}
