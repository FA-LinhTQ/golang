package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Publish bool   `json:"publish" gorm:"column:publish;"`
}

var dns = "host=localhost user=default password=secret dbname=golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func main() {
	initDB()
	serverReady := runserver()
	serverReady.Spin()
}

func initDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
}

func runserver() (h *server.Hertz) {
	h = server.Default(server.WithHostPorts(":8080"))
	registerRoutes(h)
	return
}

func registerRoutes(h *server.Hertz) {
	h.GET("/get-all", func(c context.Context, ctx *app.RequestContext) {
		var posts []Post
		var errMessage string
		if err := DB.Table("posts").Find(&posts).Error; err != nil {
			errMessage = err.Error()
		}

		ctx.JSON(consts.StatusOK, utils.H{
			"status":     len(errMessage) == 0,
			"posts":      posts,
			"errMessage": errMessage,
		})
	})
	h.POST("/get-by-id", func(c context.Context, ctx *app.RequestContext) {
		name := ctx.Query("name")
		fmt.Println(name)
		// var posts Post
		// var errMessage string
		// if err := DB.Table("posts").Find(&posts).Error; err != nil {
		// 	errMessage = err.Error()
		// }

		// ctx.JSON(consts.StatusOK, utils.H{
		// 	"status":     len(errMessage) == 0,
		// 	"posts":      posts,
		// 	"errMessage": errMessage,
		// })
	})
	h.POST("/create", func(c context.Context, ctx *app.RequestContext) {
		var post Post
		var errMessage string
		ctx.BindQuery(&post)

		if err := DB.Table("posts").Create(&post).Error; err != nil {
			errMessage = err.Error()
		}
		ctx.JSON(consts.StatusOK, utils.H{
			"status":     len(errMessage) == 0,
			"post":       post,
			"errMessage": errMessage,
		})
	})
}
