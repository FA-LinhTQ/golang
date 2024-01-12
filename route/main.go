package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	normalServer()
	// serverWithHandleMethodNotAllow()
}

func normalServer() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8888"))
	// h := server.Default(server.WithRedirectTrailingSlash(false)) // server with no redirect slash
	registerRoutes(h)
	regiterRoutesWithName(h)
	h.Spin()
}

func serverWithHandleMethodNotAllow() {
	h := server.Default(server.WithHandleMethodNotAllowed(true))
	h.POST("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	// set NoRoute handler
	h.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "no route")
	})
	// set NoMethod handler
	h.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "no method")
	})

	h.Spin()
}

func regiterRoutesWithName(h *server.Hertz) {
	h.GETEX("ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, app.GetHandlerName(c.Handler()))
	}, "ping_handler")
	h.POSTEX("pong", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, app.GetHandlerName(c.Handler()))
	}, "pong_handler")
}

func registerRoutes(h *server.Hertz) {
	h.GET("/user/:name/:age", func(ctx context.Context, c *app.RequestContext) {
		name := c.Param("name")
		age := c.Param("age")
		c.String(consts.StatusOK, name+" "+age)
	})
	h.GET("/src/*filepath", func(ctx context.Context, c *app.RequestContext) {
		filepath := c.Param("filepath")
		c.String(consts.StatusOK, filepath)
	})
	h.GET("/test/", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "/test/")
	})
	// h.GET("/test", func(ctx context.Context, c *app.RequestContext) {
	// 	c.String(consts.StatusOK, "/test")
	// })
}
