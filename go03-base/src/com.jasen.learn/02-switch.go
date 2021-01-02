package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is a weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour()<12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatamI := func(i interface{}) {
		switch t:= i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am an int")
		default:
			fmt.Printf("Don't know type %T\n",t)
		}
	}
	whatamI(true)
	whatamI(1)
	whatamI("hey")
}
