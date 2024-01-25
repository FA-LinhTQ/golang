package main

import (
	"context"
	"time"

	"example.com/crud"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	initServer()
}

func initServer() {
	h := server.Default(server.WithHostPorts(":8080"))
	registerRoutes(h)
	h.Spin()

}

func registerRoutes(h *server.Hertz) {
	h.GET("/get", func(ctx context.Context, c *app.RequestContext) {
		session := crud.CreateSession()
		crud.GetAndPrint(session)
		crud.CloseSession(session)
		// name := c.Param("name")
		// age := c.Param("age")
		// c.String(consts.StatusOK, name+" "+age)
	})
	h.GET("/create", func(ctx context.Context, c *app.RequestContext) {
		name := c.Param("name")
		age := c.Param("age")
		c.String(consts.StatusOK, name+" "+age)
	})
	h.GET("/update", func(ctx context.Context, c *app.RequestContext) {
		name := c.Param("name")
		age := c.Param("age")
		c.String(consts.StatusOK, name+" "+age)
	})
	h.GET("/delete", func(ctx context.Context, c *app.RequestContext) {
		name := c.Param("name")
		age := c.Param("age")
		c.String(consts.StatusOK, name+" "+age)
	})
	// h.GET("/src/*filepath", func(ctx context.Context, c *app.RequestContext) {
	// 	filepath := c.Param("filepath")
	// 	c.String(consts.StatusOK, filepath)
	// })
	// h.GET("/test/", func(ctx context.Context, c *app.RequestContext) {
	// 	c.String(consts.StatusOK, "/test/")
	// })
	// h.GET("/test", func(ctx context.Context, c *app.RequestContext) {
	// 	c.String(consts.StatusOK, "/test")
	// })
}

type PetInfo struct {
	Name      string
	HeartRate int
	Time      time.Time
}
