# [v2lbsyun](https://github.com/yanxianhe/lbsyun)

- golang


- 目录结构

~~~~~~
.
├── config                                # 初始化 config 文件配置
│   └── InitConfig.go                     
├── config.toml                           # 配置文件
├── go.mod                                # 依赖包
├── go.sum                                # 依赖包
├── main.go                               # 入口文件
├── main_test.go                          # 测试文件
├── Makefile                              # 编译文件
├── public                                # 公共方法
│   ├── client                            # 客户端
│   │   ├── http.go                       # http 客户端
│   │   └── redis.go                      # redis 
│   └── utils                             # 工具类
│       └── utils.go
├── v2l_context                           # 上下文
│   └── centext.go
└── v2lmethod                             # 接口方法
    ├── baiduv1_method.go
    ├── redis_method.go
    └── v2lstructs                        # 结构体
        └── response_struct.go
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
- 当前使用 12 个替换域名如果不够自行添加
- 通过替换外部域名 utils/utils.go DomainsMap
- 清空redis 可以自行添加验证。当前是两次 md5 加 ak 认证
- 具体看config.toml   配置文件