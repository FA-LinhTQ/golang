package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type Args struct {
	Query      string   `query:"query"`
	QuerySlice []string `query:"q"`
	Path       string   `path:"path"`
	Header     string   `header:"header"`
	Form       string   `form:"form"`
	Json       string   `json:"json"`
	Vd         int      `query:"vd" vd:"$==0||$==1"`
}

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.POST("v:path/bind", func(ctx context.Context, c *app.RequestContext) {
		var arg Args
		err := c.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}
		fmt.Println(arg)
	})

	h.Spin()
}
