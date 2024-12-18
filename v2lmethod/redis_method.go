package v2lmethod

import (
	"fmt"
	"v2lbsyun/public/client"
	"v2lbsyun/v2lmethod/v2lstructs"
)

func PingGetRedisKeyCount() v2lstructs.ResponseData {
	redisClient := client.InitRedis()
	keyCount := client.KeyCount(redisClient, "*")
	var message string
	var data interface{}

	if keyCount == -1 {
		message = "Redis连接失败"
		data = nil
	} else {
		message = "Redis连接成功"
		data = map[string]string{"keyCount": fmt.Sprintf("%d", keyCount)}

	}
	return v2lstructs.ResponseData{
		Message: message,
		Data:    data,
		Status:  200,
	}
}
