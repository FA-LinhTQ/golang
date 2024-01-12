package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeX2 struct {
	m sync.Mutex
	v map[int]int
}

func (s *SafeX2) Assign(key int) {
	s.m.Lock()
	s.v[key] = key
	s.m.Unlock()
}

func (s *SafeX2) Value(key int) int {
	s.m.Lock()
	defer s.m.Unlock()
	fmt.Println(s)
	fmt.Println(key)
	return s.v[key]
}

func main() {
	s := SafeX2{v: make(map[int]int)}
	for i := 0; i < 5; i++ {
		go s.Assign(i)
	}

	fmt.Println(s.Value(2))
	time.Sleep(time.Second)
}
