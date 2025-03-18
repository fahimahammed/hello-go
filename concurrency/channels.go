package concurrency

import (
	"fmt"
	"time"
)

func SendData(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "Hello from SendData"
}

func DemoChannels() {
	fmt.Println("Go Channels Practice")
	ch := make(chan string)

	go SendData(ch)

	message := <-ch
	fmt.Println(message)
}
