package v2lcontext

import (
	"github.com/kataras/iris/v12"
)

func NotFoundHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  "200",
		"message": "Not Found",
		"data":    nil,
	})
}
