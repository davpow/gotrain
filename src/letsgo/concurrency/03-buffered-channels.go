/**
Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.
 */
package main

import "fmt"

func main() {
	var powr = []int{1, 2, 4, 8, 16, 32, 64, 128}
	indexChan := make(chan int, len(powr))
	valueChan := make(chan int, len(powr))

	for i,v := range powr{
		indexChan <- i
		valueChan <- v
	}

	for range powr{
		fmt.Println(fmt.Sprintf("Index=%v, Value=%v", <-indexChan, <-valueChan))
	}
}
