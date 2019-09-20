package main

import "fmt"

/*
* channel direction "bisa ngirim dan terima" (ch chan string)
*
*
 */

//hanya bisa menerima (ch chan <- string)
func sendMessageCh(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)

}

//hanya bisa menerima (ch <- chan string)
func printMessageCh(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
