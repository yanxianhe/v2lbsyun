# [v2lbsyun](https://github.com/yanxianhe/lbsyun)


- pyrhon3.10.0 + 
- 本地运行需要安装依赖
~~~~~~
 pip install --upgrade pip install -r requirements.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
~~~~~~

- 项目目录说明
~~~~~~

└── lbsyun
    ├── docker-compose-v2lbsyun.yml    # docker-compose 启动文件 docker-compose -f ./docker-compose-v2lbsyun.yml up -d
    ├── Dockerfile                     # dockerfile 构建镜像文件
    ├── __init__.py                    # 空文件
    ├── main.py                        # 主入口
    ├── modules                        # 模块目录
    │   ├── db_client.py               # 数据模块文件
    │   └── __init__.py                # 空文件
    ├── public                         # 公共类目录
    │   ├── __init__.py                # 空文件
    │   ├── metadata.py                # 接口注释
    │   ├── sys_logs.py                # logs 日志
    │   ├── system.py                  # 系统配置
    │   └── utils.py                   # 工具
    ├── README.md                      # 说明
    ├── requirements.txt               # pip 依赖
    └── run.sh                         # 运行脚本
~~~~~~

- lbsyun

- 私有化部署使用外网百度地图开放平
- 需要DMZ区代理百度地图开放平域名替换成内网可以访问的地址
- 需要先引入地图的jsapi https 请求方式

```
<script src="//api.map.baidu.com/api?v=1.0&type=webgl&ak=你的密钥"></script>
```

https://api.map.baidu.com/api?v=1.0&type=webgl&ak=你的密钥

- 获取静态方法
- 根据方法中的地址获取js 及 css (由于百度地图实时更新建议定期更新)
- 把js css 中的地址替换成内部dmz区代理地址可以访问到外网百度地图资源
- 替换百度地图 jsapi 。地址为 dmz 代理地址 获取jsapi 方法

```
(function() {
	window.BMAP_PROTOCOL = "https";
	window.BMapGL_loadScriptTime = (new Date)
		.getTime();
	document.write('<script type="text/javascript" src="https://api.map.baidu.com/getscript?type=webgl&v=1.0&ak=你的密钥&services=&t=20230913103653"></script>');
	document.write('<link rel="stylesheet" type="text/css" href="https://api.map.baidu.com/res/webgl/10/bmap.css"/>');
})();

```

- 把百度地图资源加载到本地redis中 保存 24小时
- 通过替换外部域名可以通过环境变量修也可以修改 代码中的默认值
- 清空redis 可以自行添加验证。当前是两次 md5 加 ak 认证
- 当前使用 12 个替换域名如果不够自行添加

- DMZ 区域能正常访问外网地图的机器 10.10.10.10
- 部署 v2lbsyun 机器            192.168.31.202 开放 8000 端口
- DMZ区替换域名参数
- * 系统环境变量默认值

  |百度源地址 env|百度源地址|内部NG env|内部NG地址|备注|
  |----|----|----|----|----|
  |MAPOPEN_CDN_BCEBOS_COM | mapopen.cdn.bcebos.com|            NG_MAPOPEN_CDN_BCEBOS_COM |10.10.10.10:770/mapopen_cdn_bcebos_cmd|三级域名|
  |API_MAP_BAIDU_COM      | api.map.baidu.com|                 NG_API_MAP_BAIDU_COM      |10.10.10.10:770/api_map_baidu_cmd|三级域名|
  |HM_BAIDU_COM           | hm.baidu.com|                      NG_HM_BAIDU_COM           |10.10.10.10:770/hm_baidu_cmd|          二级域名|
  |MAPONLINE0_BDIMG_COM   | maponline0.bdimg.com|              NG_MAPONLINE0_BDIMG_COM   |10.10.10.10:770/maponline0_bdimg_cmd|二级域名|
  |MAPONLINE1_BDIMG_COM   | maponline1.bdimg.com|              NG_MAPONLINE1_BDIMG_COM   |10.10.10.10:770/maponline1_bdimg_cmd|二级域名|
  |MAPONLINE2_BDIMG_COM   | maponline2.bdimg.com|              NG_MAPONLINE2_BDIMG_COM   |10.10.10.10:770/maponline2_bdimg_cmd|二级域名|
  |MAPONLINE3_BDIMG_COM   | maponline3.bdimg.com|              NG_MAPONLINE3_BDIMG_COM   |10.10.10.10:770/maponline3_bdimg_cmd|二级域名|
  |WEBMAP0_BDIMG_COM      | webmap0.bdimg.com|                 NG_WEBMAP0_BDIMG_COM      |10.10.10.10:770/webmap0_bdimg_cmd|二级域名|
  |PCOR_BAIDU_COM         | pcor.baidu.com|                    NG_PCOR_BAIDU_COM         |10.10.10.10:770/pcor_baidu_cmd|二级域名|
  |MIAO_BAIDU_COM         | miao.baidu.com|                    NG_MIAO_BAIDU_COM         |10.10.10.10:770/miao_baidu_cmd|二级域名|
  |DLSWBR_BAIDU_COM       | dlswbr.baidu.com|                  NG_DLSWBR_BAIDU_COM       |10.10.10.10:770/dlswbr_baidu_cmd|二级域名|
  |MAP_BAIDU_COM          | map.baidu.com|                     NG_MAP_BAIDU_COM          |10.10.10.10:770/map_baidu_cmd|二级域名|

- 系统参数变量默认值
- 具体参数可以通过环境变量维护或直接修改代码中的默认值，建议通过环境变量维护

  |参数 env|参数|备注|
  |----|----|----|
  |HTTPS_TYPE|https://|地图默认 https 方式|
  |EXTERNAL_URL|http://192.168.31.202:8000|部署 v2lbsyun 机器 及端口 |
  |API_WEBGL|https://api.map.baidu.com/api?v=1.0&type=webgl&ak=|百度地图 jsapi 设置DMZ 区域地址|
  |LOG_LEVEL|Default|日志级别 error info success  Default|
  |REDIS_HOST|redis.local|redis 地址|
  |REDIS_PORT|6379||
  |REDIS_NAME|0||

- 需要先引入地图的jsapi https 请求方式
- API_WEBGL = "https://api.map.baidu.com/api?v=1.0&type=webgl&ak="
- 替换内部资源地址
  'API_WEBGL'="http://192.168.0.1:770/api_map_baidu_com/api?v=1.0&type=webgl&ak="



- [hub.docker.com](https://hub.docker.com/r/xianheyan/v2lbsyun)

```
docker pull xianheyan/v2lbsyun:v0.0.7 # docker pull registry.cn-beijing.aliyuncs.com/dockermg/v2lbsyun:v0.0.7 

- 启动 
docker-compose -p v2lbsyun -f ./docker-compose-v2lbsyun.yml up -d
- 停止
docker-compose -p v2lbsyun -f ./docker-compose-v2lbsyun.yml down
```
- yml docker-compose-v2lbsyun.yml

```
version: "3.7"
services:
  v2lbsyun.local:
    hostname: v2lbsyun.local
    image: xianheyan/v2lbsyun:v0.0.7
    container_name: v2lbsyun
    environment:
      - TZ=Asia/Shanghai
      - EXTERNAL_URL=http://192.168.31.202:8000
      - API_WEBGL=https://api.map.baidu.com/api?v=1.0&type=webgl&ak=
    ports:
      - "8000:8000"
  redis.local:
    hostname: redis.local
    image: redis
    restart: always
    command: redis-server --save "" --maxmemory 5gb --timeout 0 --tcp-keepalive 60 --bind 0.0.0.0
    environment:
      - TZ=Asia/Shanghai

```
### 二次开发 增加替换规则需要注意先 四级域名  三级域名  二级域名  一级域名 多级域名逐级替换具体方法 main.py

- async def getscriot_new(srt):

- 由于内网部署 FastAPI Swagger UI 打不开
- 原因使用外网静js css 态资源将静态资源下载到本地修改一下

- 如 Python 3.10.12 lib 默认库在~/.local/lib/python3.10/site-packages/fastapi/openapi/docs.py

- 修改 get_swagger_ui_html 、get_redoc_html
- 静态文件已经上传值 static 目录
* get_swagger_ui_html
~~~~~~
    swagger_js_url: str = "/static/swagger-ui-5.8.0/swagger-ui-bundle.js",
    swagger_css_url: str = "/static/swagger-ui-5.8.0/swagger-ui.css",
    swagger_favicon_url: str = "/static/favicon.png",
~~~~~~
* get_redoc_html
~~~~~~
    redoc_js_url: str = "/static/redoc-2.1.1/redoc.standalone.js",
    redoc_favicon_url: str = "/static/redoc-2.1.1/favicon.png",
~~~~~~


