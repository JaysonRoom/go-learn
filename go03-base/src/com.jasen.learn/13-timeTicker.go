package main

import (
	"fmt"
	"time"
)

func main() {
	timeTicker := time.NewTicker(2 * time.Second)
	fmt.Println("当前时间为:", time.Now())

	go func() {
		t := timeTicker.C
		fmt.Println("当前时间为:",t)
	}()

	for{
		time.Sleep(1*time.Second)
	}
}
