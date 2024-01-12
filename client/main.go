package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	runServer()
	go getReq()
	time.Sleep(time.Second * 5)
}

func runServer() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/redirect", func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusMovedPermanently, []byte("/redirect2"))
	})
	h.GET("/redirect2", func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusMovedPermanently, []byte("/redirect3"))
	})
	h.GET("/redirect3", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "pong")
	})
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, "ping:pong")
	})
	h.POST("/get-data", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, "ping:pong")
	})
	h.Spin()
}

func getReq() {
	fmt.Println("test")
	// c, err := client.NewClient()
	// if err != nil {
	// 	return
	// }
	// status, body, err := c.Get(context.Background(), nil, "http://localhost:8080/ping")
	// fmt.Printf("status=%v body=%v err=%v\n", status, string(body), err)
	// status == 200 res.Body() == []byte("pong") err == <nil>
}
