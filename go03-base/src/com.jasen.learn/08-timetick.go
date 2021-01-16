package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		<-ticker.C
		println("bb")
	}
}
