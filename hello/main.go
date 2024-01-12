// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/cloudwego/hertz/pkg/app/client"
// 	"github.com/cloudwego/hertz/pkg/app/server"
// 	"github.com/cloudwego/hertz/pkg/protocol"
// )

// func performRequest() {
// 	c, _ := client.NewClient()
// 	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
// 	req.SetRequestURI("https://api.sampleapis.com/coffee/hot")

// 	req.SetMethod("GET")
// 	_ = c.Do(context.Background(), req, resp)
// 	fmt.Printf("get response: %s\n", resp.Body()) // status == 200 resp.Body() == []byte("hello hertz")
// }

// // func main() {
// // 	h := server.New(server.WithHostPorts(":8888"))
// // 	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
// // 		ctx.JSON(consts.StatusOK, "hello hertz")
// // 	})
// // 	go performRequest()
// // 	h.Spin()
// // }

// func main() {
// 	// c, _ := client.NewClient()
// 	// req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
// 	// req.SetRequestURI("https://api.sampleapis.com/coffee/hot")

// 	// req.SetMethod("GET")
// 	// _ = c.Do(context.Background(), req, resp)
// 	// fmt.Printf("get response: %s\n", resp.Body())

// 	// h := server.New()
// 	// h.SetCustomSignalWaiter(func(err chan error) error {
// 	// 	return nil
// 	// })
// 	// h.Spin()

// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 20
	maxWorkers := 5
	wg := new(sync.WaitGroup)
	queueChan := make(chan int)
	wg.Add(maxWorkers)
	for i := 1; i <= maxWorkers; i++ {

		go func(count int) {
			for v := range queueChan {
				time.Sleep(time.Second)
				fmt.Printf("Worker %d is crawling web url %d \n ", count, v)
			}
			wg.Done()
		}(i)
	}
	for i := 1; i <= n; i++ {
		queueChan <- i
	}
	// Phải close để biết ngừng nhận dữ liệu
	close(queueChan)
	wg.Wait()
}
