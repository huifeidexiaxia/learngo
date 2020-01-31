/**
*@Author: haoxiongxiao
*@Date: 2018/9/16
*@Description: CREATE GO FILE close_channel
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type T int

func main() {
	dataCh := make(chan T, 1)
	stopCh := make(chan T)
	//notifyCh := make(chan T)
	for i := 0; i < 10000; i++ {
		go func(i int) {

			for {
				value := T(rand.Intn(10000))

				select {
				case m := <-stopCh:
					fmt.Println("接收到停止发送的信号")
					fmt.Println(m)
					return
				case dataCh <- value:

				}
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("1秒后开始接收数据")
	for {
		if d, ok := <-dataCh; ok {
			fmt.Println(d)
			if d == 9999 {
				fmt.Println("当在接收端接收到9999时告诉发送端不要发送了")
				close(stopCh)
				return
			}
		} else {
			break
		}

	}
}
