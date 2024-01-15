package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	s := server.Default(server.WithHostPorts(":8080"))
	s.GET("/resize", func(c context.Context, ctx *app.RequestContext) {
		url, hasUrl := ctx.GetQuery("url")
		if !hasUrl {
			ctx.AbortWithMsg("missing image url", consts.StatusOK)
		}
		handleResizeImage(url)
	})
	s.Spin()
}

func handleResizeImage(url string) {
	body, _ := fetchImage(url)
	img, a, d := image.Decode(bytes.NewReader(body))
	// ctx.SetContentType("image/jpeg") // Adjust for different formats
	// ctx.SetBody(img)
	fmt.Println(img)
	fmt.Println(a)
	fmt.Println(d)
}

func fetchImage(url string) ([]byte, error) {
	client, _ := client.NewClient()
	_, body, err := client.Get(context.Background(), nil, url)
	if err != nil {
		log.Println(err)
	}
	return body, err
}
