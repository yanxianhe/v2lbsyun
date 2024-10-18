package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	Secret struct {
		RY_MAP_SECRET_KEY string
	}
	Redis struct {
		RY_MAP_REDIS_HOST     string
		RY_MAP_REDIS_NAME     int
		RY_MAP_REDIS_PWD      string
		RY_MAP_REDIS_DURATION int64
	}
	Host struct {
		RY_MAP_EXTERNAL_HOST string
		RY_MAP_EXTERNAL_PORT int
	}
	Dmz struct {
		RY_MAP_DMZ_EXTERNAL_HOST         string
		RY_MAP_DMZ_EXTERNAL_PORTS_PREFIX string
		RY_MAP_DMZ_EXTERNAL_PORT_PREFIX  string
	}
	Baidu struct {
		RY_MAP_BAIDU_API_URL string
	}
}

var Cfg Configs

// 初始化配置
func InitConfig() error {
	fmt.Printf("InitConfig config.toml...\n")
	if _, err := toml.DecodeFile("config.toml", &Cfg); err != nil {
		fmt.Printf("Error decoding TOML file: %v\n", err)
		return err
	}
	fmt.Printf("Configs: %+v\n", Cfg)
	return nil
}
