package main

import (
	"fmt"
	"runtime"
)

func cobaBuff() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)
	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

}
