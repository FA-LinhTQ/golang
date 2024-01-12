package main

import (
	"errors"
	"fmt"
)

func pointerFunc() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// type Vertex struct {
// 	// X int
// 	// Y int
// 	X, Y int
// }

// var (
// 	v1 = Vertex{1, 2}
// 	v2 = Vertex{X: 10}
// 	v3 = Vertex{}
// 	p  = &Vertex{1, 2}
// )

var arr = []string{"Linh", "Tran", "Quang"}

func sliceFunc() {

	var s []int
	fmt.Println(s)
	s = append(s, 5)
	fmt.Println(s)

	// a := make([]int, 5)
	// fmt.Println()

	// var s []int
	// fmt.Println(s)
	// if s == nil {
	// 	fmt.Println("nil")
	// }

	// var primes = [6]int{2, 3, 4, 7, 11, 13}
	// fmt.Println(cap(primes))
	// fmt.Println(primes)
	// var s []int = primes[1:4]
	// s[0] = 99
	// fmt.Println(s)
	// fmt.Println(primes)
}

type Vertex struct {
	Lat, Long float64
}

func closureFunc() func(int) int {
	var sum int
	return func(length int) int {
		for i := 0; i < length; i++ {
			sum += i
		}
		return sum
	}
}

// var m map[string]Vertex

type Caculator struct {
	X, Y int
}

func (cal Caculator) sum() int {
	return cal.X + cal.Y
}

func Hello() {
	fmt.Println("hello form greetings")
}

func main() {
	// c := Caculator{1, 2}
	// fmt.Println(c.sum())
	// v := Vertex{1, 2}
	// p := &v
	// p.Lat = 888
	// fmt.Println(v)
	// sum := closureFunc()
	// fmt.Println(sum(5))
	// m := make(map[string]int)
	// m["Age"] = 25
	// m["Age"] = 24
	// m["level"] = 1
	// delete(m, "level")
	// v, ok := m["level"]
	// fmt.Println(m)
	// fmt.Println(v, ok)

	// m := make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// m["Bell Labs 2"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// fmt.Println(m)

	// for i, _ := range arr {
	// 	fmt.Println(i)
	// 	// fmt.Println(_)
	// }
	// sliceFunc()
	// v := Vertex{1, 2}
	// p := v
	// p.X = 10
	// fmt.Println(p)
	// fmt.Println(v)
	// var arr [2]string
	// arr[0] = "Hello"
	// arr[1] = "World"
	// fmt.Println(arr[0])
	// var primes = [6]int{2, 3, 4, 7, 11, 13}
	// fmt.Println(primes)
	// fmt.Println(v1)
	// fmt.Println(v2)
	// fmt.Println(v3)
	// fmt.Println(p)
	// countBits(126)
	// fmt.Println(countBits(126))
	errorFunc()
}

// func countBits(num uint32) int32 {
// 	var count int32
// 	for i := 0; i < 32; i++ {
// 		if bit := num & (1 << i); bit != 0 {
// 			count++
// 		}
// 	}
// 	for _, v := range v {

// 	}
// 	return count
// }

func errorFunc() error {
	return errors.New("Unkown error")
}
