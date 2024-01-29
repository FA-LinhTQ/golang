package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

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
	h.POST("/generate-url", func(c context.Context, ctx *app.RequestContext) {
		var dataMapping struct {
			Social string `json:"social"`
		}
		data, err := ctx.Body()
		if err != nil {
			return
		}

		if err := json.Unmarshal(data, &dataMapping); err != nil {
			return
		}

		fakeShopID := "1350"

		authUrl := generateUrl(dataMapping.Social, fakeShopID) // return string url

		ctx.JSON(consts.StatusOK, utils.H{
			"status": true,
			"data": utils.H{
				"url": authUrl,
			},
		})
	})
}

func generateUrl(social, shopId string) string {
	// satate := map[string]string{
	// 	"shop_id":   shopId,
	// 	"auth_type": social,
	// }

	permissions := []string{"pages_show_list", "instagram_manage_insights", "instagram_basic, pages_read_engagement, business_management"} // Thay bằng các quyền cần thiết

	// Tạo các tham số truy vấn
	params := url.Values{}
	params.Set("client_id", "1261181917604110")
	params.Set("redirect_uri", "http://localhost/social/auth-handle/facebook")
	params.Set("state", "linhdeptry")
	params.Set("display", "popup")
	params.Set("response_type", "code")
	params.Set("scope", strings.Join(permissions, ","))

	url := fmt.Sprintf("https://facebook.com/dialog/oauth?%s", params.Encode())

	return url
}
