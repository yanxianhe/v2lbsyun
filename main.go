package main

import (
	"fmt"
	"v2lbsyun/config"
	v2lcontext "v2lbsyun/v2l_context"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	if err := config.InitConfig(); err != nil {
		fmt.Println("Failed to init config:", err)
		return
	}
	app.OnErrorCode(iris.StatusNotFound, v2lcontext.NotFoundHandler)
	v1 := app.Party("/check")
	{
		v1.Get("/ping", v2lcontext.Ping)
	}
	// Simple group: v1
	v2l := app.Party("/v2lbsyun")
	{
		v2l.Get("/webgl/{ak:string}", v2lcontext.Webgl)
		v2l.Get("/getscript/{ak:string}", v2lcontext.Script)
		v2l.Get("/getbmapcss/{ak:string}", v2lcontext.Bmapcss)
		v2l.Post("/clear/{ak:string}", v2lcontext.DeleteRedisKey)
		v2l.Post("/encrypt", v2lcontext.GetEncrypt)
	}
	PORT := config.Cfg.Host.RY_MAP_EXTERNAL_PORT
	app.Listen(fmt.Sprintf(":%d", PORT))
}
