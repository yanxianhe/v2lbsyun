package v2lcontext

import (
	"v2lbsyun/v2lmethod"

	"github.com/kataras/iris/v12"
)

/*
* 探针服务，目的项目及依赖探活
* 连接Redis，并获取Redis中key的数量
 */
func Ping(ctx iris.Context) {
	ctx.JSON(v2lmethod.PingGetRedisKeyCount)
}

// 调用百度V1 API 入口
func Webgl(ctx iris.Context) {
	v2lmethod.GetBaiduV1Webgl(ctx)
}

// 获取替换后的 js
func Script(ctx iris.Context) {
	v2lmethod.GetRedisBaiduV1Getscript(ctx)
}

// 获取替换后的 css
func Bmapcss(ctx iris.Context) {
	v2lmethod.GetRedisBaiduV1GetCSS(ctx)
}

// 删除 Redis key
func DeleteRedisKey(ctx iris.Context) {
	v2lmethod.DeleteRedisBaiduV1Key(ctx)
}

// 根据明文法获取密文
func GetEncrypt(ctx iris.Context) {
	v2lmethod.GetCiphertext(ctx)
}
