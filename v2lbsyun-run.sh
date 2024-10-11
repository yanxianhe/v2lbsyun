#!/usr/bin/env bash
set -e

filePath=`dirname $0`
export RY_MAP_WORK=${filePath}

# - 指定日志路径
export RY_MAP_LOG_DIR=${RY_MAP_WORK}/logs

##########################REDIS db配置##########################
## 是否开启AES/DES加密，Default: False
export RY_MAP_DB_ENCRYPTION="DES"
## - 盐值 
export RY_MAP_SECRET_KEY="948440a0cf7c6dfd2ceeb56f9e7571201f2c6834d301ebb9c441ca267765331b"

## - redis 配置
export RY_MAP_REDIS_HOST=192.168.1.202
export RY_MAP_REDIS_PORT=6379
export RY_MAP_REDIS_NAME=1
export RY_MAP_REDIS_PWD=ae6209c6ad713f7760266d5ea375989b
##########################REDIS db配置##########################

##########################本地IP 端口，用于提供v2lbsyun请求地址##########################
export RY_MAP_EXTERNAL_HOST=192.168.0.202
export RY_MAP_EXTERNAL_PORT=9095
##########################本地IP 端口，用于提供v2lbsyun请求地址##########################


########################## https / http DMZ 区域代理 ##########################

### RY_MAP_DMZ_EXTERNAL_xx_PREFIX port前缀，
##用于服务请求地址，默认为空，如需设置，请按照格式设置，如：":80/11443" 或者 "/11080"
# - 默认
#export RY_MAP_DMZ_EXTERNAL_HOST=10.138.3.36
#export RY_MAP_DMZ_EXTERNAL_PORTS_PREFIX="/11443"
#export RY_MAP_DMZ_EXTERNAL_PORT_PREFIX="/11080"
 
########################## https / http DMZ 区域代理 ##########################


cd ${RY_MAP_WORK}
sh -c ./v2lbsyun
## nohup v2lbsyun > /dev/null 2>&1 &

## v2lbsyun 服务结束
## ps aux | grep 'v2lbsyun'
## kill -15 

