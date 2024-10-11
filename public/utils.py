
#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   utils.py
@Time    :   2023/08/09 00:24:42
@Author  :   yxh 
@Version :   1.0
@Contact :   xianhe_yan@sina.com
'''

import os,uuid,base64,hashlib,hmac,random,string
from Crypto.Cipher import AES,DES
class Configs :
    SECRET_KEY = "948440a0cf7c6dfd2ceeb56f9e7571201f2c6834d301ebb9c441ca267765331b"

srt_key = os.environ.get("RY_MAP_SECRET_KEY", Configs.SECRET_KEY)



class UtilsTools(object) :
    
    def getUuid1(self):
        return (str(uuid.uuid1()).replace("-", ""))
    def generate_ak_sk():
        # 生成随机的 AK 和 SK
        ak = os.urandom(16).hex().upper()
        sk = os.urandom(32).hex()
        return ak, sk
#### base64_tools begin
class base64_tools(object) :
    ###### base64  ######
    def srt_base64(srt_obj) :
        ## 字符串转 base64
        bs = base64.b64encode(srt_obj.encode("UTF-8"))
        # bytes 转 字符串
        return str(bs,encoding="UTF-8")
    def base64_srt(srt_obj) :
        # 字符串 转 bytes
        srt_bs = bytes(srt_obj,encoding="UTF-8")
        # base64 还原
        return base64.b64decode(srt_bs).decode("utf-8")
    def get_base64(srt_obj) :
        return  bytes(srt_obj,encoding="UTF-8")
    def get_srt(bytes_obj) :
        return  str(bytes_obj,encoding="UTF-8")
    ###### base64 end  ######
#### base64_tools end

#### pycryptos begin
class pycryptos(object) :
    #### DES #### 
    # 使用 DES加密数据的长度须为8的的倍数
    def des_encrypt(srt) :
        srt = base64_tools.srt_base64(srt)
        keys = hashlib.md5(srt_key.encode(encoding='UTF_8')).hexdigest()[0:8]
        key = bytes(keys, encoding="utf8")
        if len(srt) % 8 != 0:
            srt = srt + " " * (8 - len(srt) % 8)
        des = DES.new(key, DES.MODE_ECB)
        pas_enc = des.encrypt(srt.encode()).hex()
        return pas_enc

    def des_decrypt(pas_en) :
        
        keys = hashlib.md5(srt_key.encode(encoding='UTF_8')).hexdigest()[0:8]
        key = bytes(keys, encoding="utf8")

        des = DES.new(key, DES.MODE_ECB)
        pas_dec = des.decrypt(bytes.fromhex(pas_en))

        dtr_dec = str(pas_dec, encoding='utf-8').replace(" ","")
        return base64_tools.base64_srt(dtr_dec)
    #### DES end #### 

    #### AES ####
    def aes_encrypt(srt) :
        srt = base64_tools.srt_base64(srt)
        keys = hashlib.md5(srt_key.encode(encoding='UTF_8')).hexdigest()[0:16]
        key = bytes(keys, encoding="utf8")
        aes =  AES.new(key,AES.MODE_CFB,key)
        return aes.encrypt(srt.encode()).hex()
    def aes_decrypt(pas_en) :
        keys = hashlib.md5(srt_key.encode(encoding='UTF_8')).hexdigest()[0:16]
        key = bytes(keys, encoding="utf8")
        aes = AES.new(key, AES.MODE_CFB, key)
        pas_en = aes.decrypt(bytes.fromhex(pas_en))
        dtr_dec = str(pas_en, encoding='utf-8')
        return base64_tools.base64_srt(dtr_dec)
    #### pycryptos end
if __name__ == "__main__":
    print(UtilsTools().getUuid1())
    srt = "yanxianhe"
    # des 
    enp = pycryptos.aes_encrypt(srt)
    print(f"aes加密::[{enp}]")
    pas = pycryptos.aes_decrypt(enp)
    print(f"aes解密::[{pas}]")