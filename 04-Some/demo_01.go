package main

import "fmt"

func main() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	//var t1 chan int
	//t1 <- pipe

	//pipe <- 4

	fmt.Println(len(pipe))

}
