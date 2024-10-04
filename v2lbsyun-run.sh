#!/usr/bin/env bash
set -e
# - 指定日志路径
filePath=`dirname $0`
export RY_MAP_WORK=${filePath}
export RY_MAP_LOG_DIR=${RY_MAP_WORK}/logs
# - 服务请求地址
export RY_MAP_EXTERNAL_HOST=192.168.0.12
export RY_MAP_EXTERNAL_PORT=9095
# - https / http DMZ 区域代理
export RY_MAP_DMZ_EXTERNAL_HOST=10.138.4.198
export RY_MAP_DMZ_EXTERNAL_PORTS=11443
export RY_MAP_DMZ_EXTERNAL_PORT=11080
# - redis 配置
export RY_MAP_REDIS_HOST=127.0.0.1
# RY_MAP_REDIS_PORT=6379
# RY_MAP_REDIS_NAME=1
# RY_MAP_REDIS_PWD=redis1

cd ${RY_MAP_WORK}
sh -c ./v2lbsyun
#nohup v2lbsyun > /dev/null 2>&1 &

# v2lbsyun 服务结束
# ps aux | grep 'v2lbsyun'
# kill -15 

