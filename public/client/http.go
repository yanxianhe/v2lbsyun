package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
	"v2lbsyun/config"
)

// sendHTTPSGetRequest 发起 HTTPS GET 请求并返回响应内容
func SendHTTPSGetRequest(url string) (string, error) {
	// // 创建 HTTP 客户端，并设置超时时间为20秒
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.Baidu.RY_MAP_BAIDU_TIMEOUT) * time.Second,
		// 创建 Transport 对象并禁用 SSL 证书验证
		// 注意：在生产环境中，不推荐禁用证书验证
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 发起 GET 请求
	fmt.Println("SendHTTPSGetRequest :: url:", url)
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 返回响应体内容
	return string(body), nil
}
