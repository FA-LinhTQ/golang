package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)

		ch2 <- 2
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
