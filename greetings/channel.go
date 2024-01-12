package main

import (
	"fmt"
	"sync"
	"time"
)

func wait1Ty(wg *sync.WaitGroup) {
	fmt.Println(123)
	wg.Done()
}
func wait2Ty(wg *sync.WaitGroup) {
	fmt.Println(345)
	wg.Done()
}

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	fmt.Println(<-ch1)
	time.Sleep(time.Second * 3)
	// ch2 := make(chan int)
	// wg := new(sync.WaitGroup)
	// wg.Add(2)
	// go wait1Ty(wg)
	// go wait2Ty(wg)
	// wg.Wait()
	// go func() {
	// 	time.Sleep(time.Second * 5)
	// 	quit <- 0
	// }()

	// for {
	// 	select {
	// 	case x := <-ch1:
	// 		fmt.Println(x)
	// 	case y := <-ch2:
	// 		fmt.Println(y)
	// 		// case <-quit:
	// 		// 	fmt.Println("quit")
	// 		// 	return
	// 	}
	// }

	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)
	// time.Sleep(time.Second * 1)

	// wg := new(sync.WaitGroup)
	// ch := make(chan int)
	// go func() {
	// 	wg.Add(1)
	// 	ch <- 1
	// 	fmt.Println("Success")
	// 	wg.Done()
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// ch := make(chan int)
	// quit := make(chan int)
	// // go sendOnly(ch)
	// // receiveOnly(ch)
	// // select {
	// // case <-ch:
	// // 	fmt.Println(<-ch)
	// // }

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// 	quit <- 0
	// }()

	// // for i := 0; i < 10; i++ {
	// // 	fmt.Println(<-ch)
	// // }
	// for {
	// 	select {
	// 	case x := <-ch:
	// 		fmt.Println(x)
	// 	case <-quit:
	// 		fmt.Println("quit")
	// 		return
	// 	}
	// }

}

// func sendOnly(c chan<- int) {
// 	c <- 6
// }

// func receiveOnly(c <-chan int) {
// 	fmt.Println(<-c)
// }
