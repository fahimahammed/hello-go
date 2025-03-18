package concurrency

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func DemoGoroutines() {
	fmt.Println("Go Concurrency")
	go PrintNumbers() // Asynchronous (Non-blocking)

	time.Sleep(time.Second * 3)
	fmt.Println("Main function finished")
}
