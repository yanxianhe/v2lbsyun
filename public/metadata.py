#!/usr/bin/env python

#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   metadata.py
@Time    :   2023/09/16 00:21:58
@Author  :   yxh 
@Version :   1.0
@Contact :   xianhe_yan@sina.com
'''



#####################################################################################################
###*******************************************************************************************#######
###*************************                接口注释              *****************************#######
###*******************************************************************************************#######

class Tags(object) :
    def tags_metadata():
        tags_metadata = [
            {
                "name": "ping",
                "description": "服务探针[存活及数据库]",
            },
            {
                "name": "jsapi",
                "description": "地图的jsapi 具体看百度地址api https://lbsyun.baidu.com/solutions/mapvdata",
            },
            {
                "name": "getscript",
                "description": "获取百度地图JavaScript ",
            },
            {
                "name": "getbmapcss",
                "description": "获取百度地图 css ",
            },
            {
                "name": "clear",
                "description": "清理换成重新拉取资源 ",
            },
            {
                "name": "togeographic",
                "description": "大地2000坐标转换为经纬度 ",
            },
            {
                "name": "crscgcs2000",
                "description": "CGCS2000坐标与经纬度（WGS84）之间的转换 WGS84转换为百度坐标（BD-09）只读取只读取第一个工作表 第3 4 列的值转换 ",
            }
        ]
        return tags_metadata
    
