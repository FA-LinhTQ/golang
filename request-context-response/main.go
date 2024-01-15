package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	runServer()
}

func runServer() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/user", func(c context.Context, ctx *app.RequestContext) {
		ctx.Write([]byte(`{"foo":"bar"}`))
		ctx.SetContentType("application/json; charset=utf-8")
	})
	h.GET("/cookie", func(c context.Context, ctx *app.RequestContext) {
		ctx.SetCookie("user", "hertz", 1, "/", "localhost", protocol.CookieSameSiteLaxMode, true, true)
		cookie := ctx.Response.Header.Get("Set-Cookie")
		// ctx.Write([]byte(`{"foo":"bar"}`))
		// ctx.SetContentType("application/json; charset=utf-8")
		fmt.Println(cookie)
	})
	h.GET("/abort", func(c context.Context, ctx *app.RequestContext) {
		// ctx.AbortWithStatus(consts.StatusOK)
		// ctx.AbortWithError(consts.StatusOK, errors.New("hertz error"))
		// ctx.AbortWithMsg("loi roi fen oiw", consts.StatusOK)
		ctx.AbortWithStatusJSON(consts.StatusOK, utils.H{
			"foo":  "bar",
			"html": "<b>",
		})
	}, func(c context.Context, ctx *app.RequestContext) {
		// will not execute
	})
	h.GET("/render", func(c context.Context, ctx *app.RequestContext) {
		data := Data{"Linh", 24}
		ctx.JSON(consts.StatusOK, data)
		// ctx.IndentedJSON(consts.StatusOK, data)
	})
	h.GET("/pureJson", func(ctx context.Context, c *app.RequestContext) {
		c.PureJSON(consts.StatusOK, utils.H{
			"html": "<p> Hello World </p>",
		})
	})
	h.GET("/file", func(c context.Context, ctx *app.RequestContext) {
		ctx.File("https://cdn-stag.alireviews.io/uploads/76375032127/gwvyFeDkAiJBTUFo6i3WU0ObA6qvw4tonsulZxhu.png")
		// ctx.FileAttachment("./test.jpg", "hehe.jpg")
	})
	h.Spin()
}

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
