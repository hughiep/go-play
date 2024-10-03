package main

import "fmt"

func asyncTask() {
	println("asyncTask")
}

func main() {
	// Create channel
	ch := make(chan int)
	chReceive := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		a := <-ch
		chReceive <- a
	}()

	// Wait for channel
	b := <-chReceive
	fmt.Println(b)
}
