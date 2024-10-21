package v2lmethod

import (
	"fmt"
	"regexp"
	"strings"
	"v2lbsyun/config"
	"v2lbsyun/public/client"
	"v2lbsyun/public/utils"
	"v2lbsyun/v2lmethod/v2lstructs"

	"github.com/kataras/iris/v12"
)

func isValidAK(ak string) bool {
	// 简单的正则表达式验证
	match, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, ak)
	return match
}

// 替换js 获取 js
func getscript(curl string, ak string) {
	body, err := client.SendHTTPSGetRequest(curl)
	if err != nil {
		fmt.Printf("请求百度地图接口失败，错误信息：%s \n", curl)
	}
	redisClient := client.InitRedis()
	// 将结果存入 Redis
	client.SetValue(redisClient, utils.RedisKeyBaiduScript(ak), string(body), config.Cfg.Redis.RY_MAP_REDIS_DURATION)
	// v2l 转化script
	v2l_body := utils.ReplaceDomains(body)
	client.SetValue(redisClient, utils.RedisKeyV2lScript(ak), string(v2l_body), config.Cfg.Redis.RY_MAP_REDIS_DURATION)
}

// 替换css 获取css
func getcss(curl string, ak string) {
	body, err := client.SendHTTPSGetRequest(curl)
	if err != nil {
		fmt.Printf("请求百度地图接口失败，错误信息：%s \n", curl)
	}
	redisClient := client.InitRedis()
	// 将结果存入 Redis
	client.SetValue(redisClient, utils.RedisKeyBaiduCSS(ak), string(body), config.Cfg.Redis.RY_MAP_REDIS_DURATION)
	// v2l 转化css
	v2l_body := utils.ReplaceDomains(body)
	client.SetValue(redisClient, utils.RedisKeyV2lCSS(ak), string(v2l_body), config.Cfg.Redis.RY_MAP_REDIS_DURATION)
}

// 替换webgl中的url
func replaceWebgl(webgl string, ak string) string {
	// 替换字符串中的 "http://" 为 "https://"
	pattern := regexp.MustCompile(`https?://[^\s"']+`)
	urls := pattern.FindAllString(webgl, -1)
	for _, url := range urls {
		if strings.Contains(url, "getscript") {
			getscript(url, ak)
			webgl = strings.Replace(webgl, url, utils.LocalScript(ak), 1)
		} else if strings.Contains(url, ".css") {
			getcss(url, ak)
			webgl = strings.Replace(webgl, url, utils.LocalCSS(ak), 1)
		}
	}
	redisClient := client.InitRedis()
	// v2l 转化后的结果  Redis
	client.SetValue(redisClient, utils.RedisKeyV2lWebgl(ak), string(webgl), config.Cfg.Redis.RY_MAP_REDIS_DURATION)
	return webgl
}

// 清除redis缓存
func clearRedisBaiduV1Key(ak string) {
	redisClient := client.InitRedis()
	client.DelValue(redisClient, utils.RedisKeyBaiduWebgl(ak))
	client.DelValue(redisClient, utils.RedisKeyBaiduScript(ak))
	client.DelValue(redisClient, utils.RedisKeyBaiduCSS(ak))

	client.DelValue(redisClient, utils.RedisKeyV2lWebgl(ak))
	client.DelValue(redisClient, utils.RedisKeyV2lScript(ak))
	client.DelValue(redisClient, utils.RedisKeyV2lCSS(ak))
}

// 检查redis缓存是否存在
func v2lCheck(ak string) string {
	redisClient := client.InitRedis()
	v2l_webgl, _ := client.GetValue(redisClient, utils.RedisKeyV2lWebgl(ak))
	v2l_script, _ := client.GetValue(redisClient, utils.RedisKeyV2lScript(ak))
	v2l_css, _ := client.GetValue(redisClient, utils.RedisKeyV2lCSS(ak))

	if v2l_webgl == "" || v2l_script == "" || v2l_css == "" {
		return ""
	}
	return v2l_webgl
}

// 在拼装数据 入口函数
func v2l_M_Webgl(ak string) string {
	redisClient := client.InitRedis()
	baidu_url := "" // 定义请求百度地图接口的url
	baidu_url = utils.GetUirBaiduWebgl(ak)
	// baidu_url = fmt.Sprintf("%s%s", config.Cfg.Baidu.RY_MAP_BAIDU_API_URL, ak)
	body, err := client.SendHTTPSGetRequest(baidu_url)
	if err != nil {
		fmt.Printf("请求百度地图接口失败，错误信息：%s \n", baidu_url)
	}
	// 将结果存入 Redis
	client.SetValue(redisClient, utils.RedisKeyBaiduWebgl(ak), string(body), config.Cfg.Redis.RY_MAP_REDIS_DURATION)
	v2l_body := replaceWebgl(body, ak)
	return v2l_body
}

func PingGetRedisKeyCount2() v2lstructs.ResponseData {
	redisClient := client.InitRedis()
	keyCount := client.KeyCount(redisClient, "*")
	someData := map[string]string{"data": fmt.Sprintf("%d", keyCount)}
	responseData := v2lstructs.ResponseData{}
	responseData.Status = 200
	if keyCount == -1 {
		responseData.Message = "Redis连接失败"
		responseData.Data = nil
	} else {
		responseData.Message = "Redis连接成功"
		responseData.Data = someData
	}
	return responseData
}

func GetBaiduV1Webgl(ctx iris.Context) {
	ak := ctx.Params().Get("ak")
	v2l_flg := ctx.Params().Get("v2l_flg")
	if !isValidAK(ak) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid AK")
		return
	}

	v2l_body := v2lCheck(ak)
	if v2l_body == "" {
		clearRedisBaiduV1Key(ak)
		v2l_body = v2l_M_Webgl(ak)
	}
	if v2l_flg != "" {
		redisClient := client.InitRedis()
		if v2l_flg == "getscript" {
			body, _ := client.GetValue(redisClient, utils.RedisKeyV2lScript(ak))
			v2l_body = body
		}
		if v2l_flg == ".css" {
			body, _ := client.GetValue(redisClient, utils.RedisKeyBaiduCSS(ak))
			v2l_body = body
		}
	}
	ctx.ContentType("application/javascript")
	ctx.WriteString(v2l_body)
}

// 获取redis缓存的js
func GetRedisBaiduV1Getscript(ctx iris.Context) {
	ak := ctx.Params().Get("ak")
	if !isValidAK(ak) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid AK")
		return
	}
	redisClient := client.InitRedis()
	body, _ := client.GetValue(redisClient, utils.RedisKeyV2lScript(ak))
	if body == "" {
		ctx.Params().Set("v2l_flg", "getscript")
		GetBaiduV1Webgl(ctx)
		return
	}
	ctx.ContentType("application/javascript")
	ctx.WriteString(body)
}

// 获取redis缓存的css
func GetRedisBaiduV1GetCSS(ctx iris.Context) {
	ak := ctx.Params().Get("ak")
	if !isValidAK(ak) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid AK")
		return
	}
	redisClient := client.InitRedis()
	body, _ := client.GetValue(redisClient, utils.RedisKeyV2lCSS(ak))
	if body == "" {
		ctx.Params().Set("v2l_flg", ".css")
		GetBaiduV1Webgl(ctx)
	}
	ctx.ContentType("application/javascript")
	ctx.WriteString(body)
}

// 删除redis缓存
func DeleteRedisBaiduV1Key(ctx iris.Context) {
	ak := ctx.Params().Get("ak")
	if !isValidAK(ak) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid AK")
		return
	}
	ctx.StatusCode(iris.StatusBadRequest)
	if utils.GetMd5(utils.GetMd5(ak)) == ctx.URLParam("sk") {
		clearRedisBaiduV1Key(ak)
		ctx.WriteString("ok")
		return
	}
	ctx.WriteString(ak)
}

func GetCiphertext(ctx iris.Context) {
	var body v2lstructs.Decrypt
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Failed to read request body: " + err.Error())
	}
	pwd := utils.GetDesEncrypt(body.Password)
	ctx.WriteString(pwd)
}
