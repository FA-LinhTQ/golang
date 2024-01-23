package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"log"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	c, _ := client.NewClient(client.WithResponseBodyStream(true))
	req := &protocol.Request{}
	resp := &protocol.Response{}
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("https://cdn.shopify.com/s/files/1/0558/5046/7427/products/Main_b13ad453-477c-4ed1-9b43-81f3345adfd6.jpg?v=1702545967")
	err := c.Do(context.Background(), req, resp)
	if err != nil {
		return
	}
	bodyStream := resp.BodyStream()
	fmt.Println(bodyStream)
	p := make([]byte, resp.Header.ContentLength()/2)
	_, err = bodyStream.Read(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	// left, _ := ioutil.ReadAll(bodyStream)
	// fmt.Println(string(p), string(left))
	// s := server.Default(server.WithHostPorts(":8080"))
	// s.GET("/resize", func(c context.Context, ctx *app.RequestContext) {
	// 	url, hasUrl := ctx.GetQuery("url")
	// 	if !hasUrl {
	// 		ctx.AbortWithMsg("missing image url", consts.StatusOK)
	// 	}
	// 	handleResizeImage(url)
	// })
	// s.Spin()
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
