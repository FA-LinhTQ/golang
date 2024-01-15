package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	runServer()
}

func runServer() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.POST("/file", func(c context.Context, ctx *app.RequestContext) {
		avatarFile, err := ctx.FormFile("avatar") // avatarFile.Filename == "abc.jpg", err == nil
		ctx.SaveUploadedFile(avatarFile, avatarFile.Filename)
		fmt.Println(avatarFile.Filename)
		fmt.Println(err)
	})
	h.GET("/user", func(c context.Context, ctx *app.RequestContext) {
		ctx.Request.Header.Add("Content-Type", "application/json; charset=utf-8")
		ctx.Request.Header.Add("hertz1", "value1")
		ctx.Set("heh", "haha")
		fmt.Println(ctx.Get("heh"))
		// ctx.Request.Header.Add("hertz1", "value2")
		// ctx.Request.Header.Set("hertz3", "value4")
		// fmt.Println(ctx.Request.Header.GetAll("hertz1"))
		// fmt.Println(ctx.Request.Header.Get("hertz3"))
		// fmt.Println(string(ctx.Request.Header.ContentType()))
		// fmt.Println(ctx.Request.Header.String())
		// fmt.Println(string(ctx.Method()))
		// fmt.Println(string(ctx.ContentType()))
		fmt.Println(string(ctx.GetHeader("hertz1")))

		// args := ctx.QueryArgs()
		// args.Set("foo", "bar")
		// fmt.Println(args)
		// name := ctx.Query("name")
		// age := ctx.Query("age")
		name := ctx.DefaultQuery("name", "Linhdz")
		age := ctx.DefaultQuery("age", "24")
		fmt.Println(name)
		fmt.Println(age)
		ctx.JSON(consts.StatusOK, string(ctx.Host()))
	})
	h.GET("/next", func(c context.Context, ctx *app.RequestContext) {
		// fmt.Println(ctx.HandlerName())
		// fmt.Println(ctx.GetIndex())
		fmt.Println(ctx.ClientIP())
		// ctx.Abort()
		// isAborted := ctx.IsAborted()
		// fmt.Println(isAborted)
		// ctx.Next(c)
		// v := ctx.GetString("version")
		// fmt.Println(v)
	}, func(c context.Context, ctx *app.RequestContext) {
		fmt.Println(ctx.GetIndex())
		ctx.Set("version", "v1")
	})
	h.GET("/ip", func(c context.Context, ctx *app.RequestContext) {
		customIp := func(ctx *app.RequestContext) string {
			return "127.0.0.1"
		}
		ctx.SetClientIPFunc(customIp)
		fmt.Println(ctx.ClientIP())
	})
	h.Spin()
}
