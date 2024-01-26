package main

import (
	"context"
	"strconv"

	"example.com/crud"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
	h.GET("/find/:id", func(ctx context.Context, c *app.RequestContext) {
		id, _ := strconv.Atoi(c.Param("id"))
		session := crud.CreateSession()
		defer crud.CloseSession(session)
		rows := crud.Find(session, id)
		c.JSON(consts.StatusOK, utils.H{
			"status": true,
			"data":   rows,
		})
	})

	h.POST("/create", func(ctx context.Context, c *app.RequestContext) {
		var pet PetInfo
		c.BindAndValidate(&pet)
		petMap := map[string]interface{}{
			"id":   pet.ID,
			"name": pet.Name,
		}
		session := crud.CreateSession()
		defer crud.CloseSession(session)
		crud.Insert(session, petMap)

		c.JSON(consts.StatusOK, utils.H{
			"status": true,
			"data":   petMap,
		})
	})
	h.POST("/update", func(ctx context.Context, c *app.RequestContext) {
		var pet PetInfo
		c.BindAndValidate(&pet)
		petMap := map[string]interface{}{
			"id":   pet.ID,
			"name": pet.Name,
		}
		session := crud.CreateSession()
		defer crud.CloseSession(session)
		crud.Update(session, petMap)

		c.JSON(consts.StatusOK, utils.H{
			"status": true,
			"data":   petMap,
		})
	})
	h.POST("/delete", func(ctx context.Context, c *app.RequestContext) {
		var pet PetInfo
		c.BindAndValidate(&pet)
		petMap := map[string]interface{}{
			"id":   pet.ID,
			"name": pet.Name,
		}
		session := crud.CreateSession()
		defer crud.CloseSession(session)
		crud.Delete(session, petMap)

		c.JSON(consts.StatusOK, utils.H{
			"status": true,
			"data":   petMap,
		})
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
	ID   int    `json:"id"`
	Name string `json:"name"`
}
