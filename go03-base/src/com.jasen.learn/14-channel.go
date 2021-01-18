package main

import (
	"fmt"
	"time"
)

func printHello(ch chan int)  {
	time.Sleep(2*time.Second)
	println("from fun2: hello")

	ch <- 2
}

func main() {
	ch := make(chan int)

	go func() {
		println("from fun1: hello")
		ch <- 1
	}()
    go printHello(ch)
	println("from main: Hello")

	t := <-ch
	fmt.Printf("from main : %d",t)
	<- ch

}
