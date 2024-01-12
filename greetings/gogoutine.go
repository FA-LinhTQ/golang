package main

import (
	"fmt"
	"time"
)

func sayMyName(name string) {
	// for i := 0; i < 5; i++ {
	fmt.Printf("Hello %s\n", name)
	// }
}

func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			ch <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// ch := make(chan int)
	// var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// go sum(arr[:len(arr)/2], ch)
	// go sum(arr[len(arr)/2:], ch)

	// x, y := <-ch, <-ch
	// fmt.Println(x)
	// fmt.Println(y)
}
