package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendDataTm(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("time out . no activities under 5 seconds", time.Second)
			break loop
		}
	}
}
