package main

import "fmt"

func main() {
	//buffered channel (canal con búfer)
	ca := make(chan int, 1)

	ca <- 42

	fmt.Println(<-ca)

}
