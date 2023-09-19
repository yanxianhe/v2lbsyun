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
                "description": "服务探针",
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
                "description": "获取百度地图JavaScript ",
            },
            {
                "name": "clear",
                "description": "清理换成重新拉取资源 ",
            }
        ]
        return tags_metadata
    
