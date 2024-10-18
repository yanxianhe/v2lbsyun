package v2lmethod

import (
	"fmt"
	"v2lbsyun/public/client"
	"v2lbsyun/v2lmethod/v2lstructs"
)

func PingGetRedisKeyCount() v2lstructs.ResponseData {
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
