package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"v2lbsyun/config"
	"v2lbsyun/v2lmethod/v2lstructs"
)

// GetMd5 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
// 参数:
//
//	str: 要计算 MD5 哈希值的字符串
//
// 返回值:
//
//	返回给定字符串的 MD5 哈希值的十六进制字符串表示形式
func GetMd5(str string) string {
	hash := md5.Sum([]byte(str))
	md5String := hex.EncodeToString(hash[:])
	return md5String
}

// BaiduScript 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
func RedisKeyBaiduWebgl(ak string) string {
	baidu_Script := fmt.Sprintf("baidu_webgl%s", ak)
	return GetMd5(baidu_Script)
}

// V2lScript 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
func RedisKeyV2lWebgl(ak string) string {
	baidu_Script := fmt.Sprintf("v2l_webgl%s", ak)
	return GetMd5(baidu_Script)
}

// BaiduScript 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
func RedisKeyBaiduScript(ak string) string {
	baidu_Script := fmt.Sprintf("baidu_getscript%s", ak)
	return GetMd5(baidu_Script)
}

// V2lScript 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
func RedisKeyV2lScript(ak string) string {
	baidu_Script := fmt.Sprintf("v2l_getscript%s", ak)
	return GetMd5(baidu_Script)
}

// BaiduCSS 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
func RedisKeyBaiduCSS(ak string) string {
	baidu_Script := fmt.Sprintf("baidu_css%s", ak)
	return GetMd5(baidu_Script)
}

// V2lCSS 计算给定字符串的 MD5 哈希值，并返回其十六进制字符串表示形式
func RedisKeyV2lCSS(ak string) string {
	baidu_Script := fmt.Sprintf("v2l_css%s", ak)
	return GetMd5(baidu_Script)
}

// DesEncrypt 使用DES算法加密给定的数据
//
// 参数：
//
//	data: 需要加密的数据
//	key: DES算法的密钥
//
// 返回值：
//
//	加密后的数据切片
//	可能发生的错误
func DesEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData := PKCS5Padding(data, block.BlockSize())
	crypted := make([]byte, len(origData))
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// DesDecrypt 使用DES算法解密给定的密文
//
// 参数:
//
//	crypted: 需要解密的密文数据
//	key: 解密密钥
//
// 返回值:
//
//	解密后的明文数据
//	如果解密过程中出现错误，则返回错误信息
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// PKCS5Padding 对给定的密文进行PKCS#5填充
// 参数：
// ciphertext: 待填充的密文切片
// blockSize: 块大小
// 返回值：填充后的密文切片
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 去除PKCS5填充的字节
// origData: 需要去除PKCS5填充的字节切片
// 返回值: 去除PKCS5填充后的字节切片
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// GetDesEncrypt 对传入的密码进行DES加密，并返回加密后的十六进制字符串
// 参数：
//
//	password string: 需要加密的密码字符串
//
// 返回值：
//
//	string: 加密后的十六进制字符串
func GetDesEncrypt(password string) string {
	key := []byte(config.Cfg.Secret.RY_MAP_SECRET_KEY)[:8]
	for len(key) < 8 {
		key = append(key, '0') // 使用'0'填充
	}
	pwd := []byte(password)
	crypted, _ := DesEncrypt(pwd, key)
	return hex.EncodeToString(crypted)
}

// GetDesDecrypt 对输入的密码进行DES解密，并返回解密后的字符串
// 参数：
//
//	password: 需要解密的密码（应为十六进制字符串）
//
// 返回值：
//
//	string: 解密后的字符串
func GetDesDecrypt(password string) string {
	key := []byte(config.Cfg.Secret.RY_MAP_SECRET_KEY)[:8]
	for len(key) < 8 {
		key = append(key, '0') // 使用'0'填充
	}
	pwd, _ := hex.DecodeString(password)
	encrypt, _ := DesDecrypt(pwd, key)
	return string(encrypt)
}

// 百度webgl 转 dmz地址
func GetUirBaiduWebgl(ak string) string {
	baidu_url := config.Cfg.Baidu.RY_MAP_BAIDU_API_URL
	dmz_ip := fmt.Sprintf("http://%s%s",
		config.Cfg.Dmz.RY_MAP_DMZ_EXTERNAL_HOST,
		config.Cfg.Dmz.RY_MAP_DMZ_EXTERNAL_PORTS_PREFIX)
	replace_url := strings.Replace(strings.Replace(
		baidu_url, "https://", dmz_ip, -1),
		"api.map.baidu.com", "/api_map_baidu_com", -1)
	dmz_url := fmt.Sprintf("%s%s", replace_url, ak)
	return dmz_url
}

// 百度Script 转 dmz地址
func LocalScript(ak string) string {
	local_url := fmt.Sprintf("http://%s:%d%s%s",
		config.Cfg.Host.RY_MAP_EXTERNAL_HOST,
		config.Cfg.Host.RY_MAP_EXTERNAL_PORT,
		"/v2lbsyun/getscript/",
		ak)
	return local_url
}

// 百度CSS 转 dmz地址
func LocalCSS(ak string) string {
	local_url := fmt.Sprintf("http://%s:%d%s%s",
		config.Cfg.Host.RY_MAP_EXTERNAL_HOST,
		config.Cfg.Host.RY_MAP_EXTERNAL_PORT,
		"/v2lbsyun/getbmapcss/", ak)
	return local_url
}

func DomainsMap() map[string]string {
	dmz_ip := fmt.Sprintf("%s%s",
		config.Cfg.Dmz.RY_MAP_DMZ_EXTERNAL_HOST,
		config.Cfg.Dmz.RY_MAP_DMZ_EXTERNAL_PORTS_PREFIX)
	// 定义替换映射
	var domainMap = map[string]string{
		// 四级域名
		// 三级域名
		"mapopen.cdn.bcebos.com": fmt.Sprintf("%s%s", dmz_ip, "/mapopen_cdn_bcebos_com"),
		"api.map.baidu.com":      fmt.Sprintf("%s%s", dmz_ip, "/api_map_baidu_com"),
		// 二级级域名
		"hm.baidu.com":         fmt.Sprintf("%s%s", dmz_ip, "/hm_baidu_com"),
		"maponline0.bdimg.com": fmt.Sprintf("%s%s", dmz_ip, "/maponline0_bdimg_com"),
		"maponline1.bdimg.com": fmt.Sprintf("%s%s", dmz_ip, "/maponline1_bdimg_com"),
		"maponline2.bdimg.com": fmt.Sprintf("%s%s", dmz_ip, "/maponline2_bdimg_com"),
		"maponline3.bdimg.com": fmt.Sprintf("%s%s", dmz_ip, "/maponline3_bdimg_com"),
		"webmap0.bdimg.com":    fmt.Sprintf("%s%s", dmz_ip, "/webmap0_bdimg_com"),
		"pcor.baidu.com":       fmt.Sprintf("%s%s", dmz_ip, "/pcor_baidu_com"),
		"miao.baidu.com":       fmt.Sprintf("%s%s", dmz_ip, "/miao_baidu_com"),
		"dlswbr.baidu.com":     fmt.Sprintf("%s%s", dmz_ip, "/dlswbr_baidu_com"),
		"map.baidu.com":        fmt.Sprintf("%s%s", dmz_ip, "/map_baidu_com"),
	}
	return domainMap
}

func ReplaceDomains(str string) string {
	domainMap := DomainsMap()
	// 获取排序后的域名级别列表
	var levels []v2lstructs.DomainLevel
	for domain := range domainMap {
		levels = append(levels, v2lstructs.DomainLevel{Domain: domain, Level: strings.Count(domain, ".")})
	}
	// 从高级别到低级别排序
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].Level > levels[j].Level
	})
	// 替换域名
	for _, level := range levels {
		str = strings.ReplaceAll(str, level.Domain, domainMap[level.Domain])
	}
	return str
}
