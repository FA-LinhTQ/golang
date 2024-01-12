package main

import "fmt"

type Vertex struct {
	X, Y int
}

// func pointerFunc() {

// }

func arrFunc() {
	// var arr [2]string
	// arr[0] = "hello"
	// arr[1] = "world"
	// fmt.Println(arr)
	var primes = []int{2, 3, 4, 7, 11, 13}
	s := primes[1:3]
	fmt.Println(primes)
	fmt.Println(s)
	s[0] = 0
	fmt.Println(primes)
	fmt.Println(s)
}

func main() {
	arrFunc()
	// pointerFunc()
	// a := Vertex{}
	// a1 := Vertex{Y: 3}
	// a2 := Vertex{1, 2}
	// p := &Vertex{1, 4}
	// fmt.Println(a)
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(p)
	// *&p.X = 7
	// fmt.Println(p)

}
