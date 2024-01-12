package main

import (
	"fmt"
	"sync"
)

func Crawl(url int, m sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	fmt.Println("crawl website: ", url)
}

func main() {
	n := 1000
	maxWorker := 5
	queueCh := make(chan int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		queueCh <- i
	}
	// close(queueCh)

	for i := 0; i < maxWorker; i++ {
		wg.Add(1)
		go func(worker int) {
			for v := range queueCh {
				fmt.Printf("Worker %d runing, crawling website %d \n", worker, v)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
