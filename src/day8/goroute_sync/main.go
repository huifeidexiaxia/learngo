package main

import (
	"fmt"
)

/**
之前写的例子都 是通过 sleep来让程序执行完，这个有个很大的弊端就是不并不清楚具体要等待多长时间

go语言中建议使用通信来实现通知
*/
func calc(taskChan chan int, resChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			// 判断是否是素数（只能被1与自身整除）
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resChan <- v
		}
	}
	// 这样做是有问题的，因为只有当8个全部完成才应该退出，写在这里有可能有的还没做完，就关闭了，肯定会有问题的
	//close(resChan)
	fmt.Println("exit")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}

		close(intChan)
	}()

	// 启动8个线程，同时来计算，提高效率，
	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	//等待所有计算的goroutine全部退出
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Println("wait goroute ", i, " exited")
		}
		close(resultChan)
	}()

	for _ = range resultChan {
		//fmt.Println(v)
	}
	fmt.Println("dddddd")
}
