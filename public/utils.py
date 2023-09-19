
#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   utils.py
@Time    :   2023/08/09 00:24:42
@Author  :   yxh 
@Version :   1.0
@Contact :   xianhe_yan@sina.com
'''

import os,uuid


class UtilsTools(object) :
    
    def getUuid1(self):
        return (str(uuid.uuid1()).replace("-", ""))
    def generate_ak_sk():
        # 生成随机的 AK 和 SK
        ak = os.urandom(16).hex().upper()
        sk = os.urandom(32).hex()
        return ak, sk

