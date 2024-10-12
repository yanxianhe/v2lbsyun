#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   system.py
@Time    :   2023/08/03 00:20:45
@Author  :   yxh 
@Version :   1.0
@Contact :   xianhe_yan@sina.com
'''

# openssl rand -hex 32

# ####################################################################################################
# ##*******************************************************************************************#######
# ##*************************                 os env              *****************************#######
# ##*******************************************************************************************#######

import os

## redis key 定义

V2l_JSAPI_KEY = "jsapi"
V2l_GETSCRIPT_KEY = "getscript"
V2l_GETBMAPCSS_KEY = "getbmapcss"
V2l_JSAPI_NEW_KEY = "jsapinew"
V2l_GETSCRIPT_NEW_KEY = "getscriptnew"
V2l_GETBMAPCSS_NEW_KEY = "getbmapcssnew"

## 代理地址 前缀
PREFIX_MAPOPEN_CDN_BCEBOS_COM="mapopen_cdn_bcebos_com"
PREFIX_HM_BAIDU_COM="hm_baidu_com"
PREFIX_API_MAP_BAIDU_COM="api_map_baidu_com"
PREFIX_MAPONLINE0_BDIMG_COM="maponline0_bdimg_com"
PREFIX_MAPONLINE1_BDIMG_COM="maponline1_bdimg_com"
PREFIX_MAPONLINE2_BDIMG_COM="maponline2_bdimg_com"
PREFIX_MAPONLINE3_BDIMG_COM="maponline3_bdimg_com"
PREFIX_WEBMAP0_BDIMG_COM="webmap0_bdimg_com"
PREFIX_PCOR_BAIDU_COM="pcor_baidu_com"
PREFIX_MIAO_BAIDU_COM="miao_baidu_com"
PREFIX_DLSWBR_BAIDU_COM="dlswbr_baidu_com"
PREFIX_MAP_BAIDU_COM="map_baidu_com"

## - 本机器IP地址、【访问请求地址】
EXTERNAL_HOST = os.environ.get("RY_MAP_EXTERNAL_HOST", "10.138.3.71")
EXTERNAL_PORT = os.environ.get("RY_MAP_EXTERNAL_PORT", "80")
if EXTERNAL_PORT == "80" or EXTERNAL_PORT == "443":
    EXTERNAL_URL_DEF = ("%s") % (EXTERNAL_HOST)
else:
    EXTERNAL_URL_DEF = ("%s:%s") % (EXTERNAL_HOST, EXTERNAL_PORT)

EXTERNAL_URL = os.environ.get("RY_MAP_EXTERNAL_URL", f"http://{EXTERNAL_URL_DEF}")
# # ## http DMZ 区域代理
DMZ_EXTERNAL_HOST = os.environ.get("RY_MAP_DMZ_EXTERNAL_HOST", "10.138.3.36")
DMZ_EXTERNAL_PORTS_PREFIX = os.environ.get("RY_MAP_DMZ_EXTERNAL_PORTS_PREFIX", "/11443")
DMZ_EXTERNAL_PORT_PREFIX = os.environ.get("RY_MAP_DMZ_EXTERNAL_PORT_PREFIX", "/11080")
DMZ_EXTERNAL_DEFS = ("%s%s") % (DMZ_EXTERNAL_HOST, DMZ_EXTERNAL_PORTS_PREFIX)
DMZ_EXTERNAL_DEF = ("%s%s") % (DMZ_EXTERNAL_HOST, DMZ_EXTERNAL_PORT_PREFIX)

# # ## 提供服务地址
HTTPS_TYPE = os.environ.get("RY_MAP_HTTPS_TYPE", "https://")
# # ## http 百度源地址 "https://api.map.baidu.com/api?v=1.0&type=webgl&ak="
## export RY_MAP_API_WEBGL="https://api.map.baidu.com/api?v=1.0&type=webgl&ak="
API_WEBGL = os.environ.get("RY_MAP_API_WEBGL")
# # ## 内部访问 百度源地址
V2l_WEBGL_ = (f"http://%s/api_map_baidu_com/api?v=1.0&type=webgl&ak=" % DMZ_EXTERNAL_DEFS) 
# ## 四级域名

# ## 三级域名
MAPOPEN_CDN_BCEBOS_COM = os.environ.get("RY_MAP_MAPOPEN_CDN_BCEBOS_COM", "mapopen.cdn.bcebos.com")
API_MAP_BAIDU_COM = os.environ.get("RY_MAP_API_MAP_BAIDU_COM", "api.map.baidu.com")
# ## 二级级域名
HM_BAIDU_COM = os.environ.get("RY_MAP_HM_BAIDU_COM", "hm.baidu.com")
MAPONLINE0_BDIMG_COM = os.environ.get("RY_MAP_MAPONLINE0_BDIMG_COM", "maponline0.bdimg.com")
MAPONLINE1_BDIMG_COM = os.environ.get("RY_MAP_MAPONLINE1_BDIMG_COM", "maponline1.bdimg.com")
MAPONLINE2_BDIMG_COM = os.environ.get("RY_MAP_MAPONLINE2_BDIMG_COM", "maponline2.bdimg.com")
MAPONLINE3_BDIMG_COM = os.environ.get("RY_MAP_MAPONLINE3_BDIMG_COM", "maponline3.bdimg.com")
WEBMAP0_BDIMG_COM = os.environ.get("RY_MAP_WEBMAP0_BDIMG_COM", "webmap0.bdimg.com")
PCOR_BAIDU_COM = os.environ.get("RY_MAP_PCOR_BAIDU_COM", "pcor.baidu.com")
MIAO_BAIDU_COM = os.environ.get("RY_MAP_MIAO_BAIDU_COM", "miao.baidu.com")
DLSWBR_BAIDU_COM = os.environ.get("RY_MAP_DLSWBR_BAIDU_COM", "dlswbr.baidu.com")
MAP_BAIDU_COM = os.environ.get("RY_MAP_MAP_BAIDU_COM", "map.baidu.com")


    


