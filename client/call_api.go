package main

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app/client"
)

func performRequest(ch chan UsersPage) {
	var userData UsersPage
	c, err := client.NewClient()
	if err != nil {
		return
	}
	_, body, _ := c.Get(context.Background(), nil, "https://reqres.in/api/users?page=2")
	json.Unmarshal(body, &userData)
	ch <- userData
}

// func main() {
// 	h := server.New(server.WithHostPorts(":8080"))
// 	h.GET("/get-data", func(c context.Context, ctx *app.RequestContext) {
// 		ch := make(chan UsersPage)
// 		go performRequest(ch)
// 		data := <-ch
// 		fmt.Println(data)
// 		ctx.JSON(consts.StatusOK, utils.H{
// 			"page":       data.Page,
// 			"per_page":   data.PerPage,
// 			"total":      data.Total,
// 			"total_page": data.TotalPages,
// 			"data":       data.Data,
// 			"support":    data.Support,
// 		})
// 	})

// 	h.Spin()
// }

type UsersPage struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []User  `json:"data"`
	Support    Support `json:"support"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Support struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}
