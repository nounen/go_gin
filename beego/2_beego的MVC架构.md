## controller 设计

### 1. 参数配置
* 默认解析 conf/app.conf

* 获取配置: `beego.AppConfig.String("mysqluser")`, 更多方法见手册

* 不同运行模式优先读取不同的配置

* 多个配置文件

* 从环境变量中获取配置值

* 框架默认配置:
```go
# 常见的

beego.BConfig.AppName = "beego"

beego.BConfig.RunMode = "dev" // 应用运行模式 prod || dev

beego.BConfig.EnableGzip = false // 是否开启 gzip 支持

beego.BConfig.MaxMemory = 1 << 26 // 文件上传默认内存缓存大小，默认值是 1 << 26(64M)

beego.BConfig.EnableErrorsShow = true

beego.BConfig.WebConfig.TemplateLeft="{{" // 模板左标签，默认值是{{

beego.BConfig.WebConfig.EnableXSRF = false // 是否开启 XSRF，默认为 false，不开启

beego.BConfig.WebConfig.XSRFExpire = 0 // 

beego.BConfig.Listen.EnableAdmin = false // 是否开启进程内监控模块，默认 false 关闭

beego.BConfig.Listen.EnableFcgi = false // 是否启用 fastcgi ， 默认是 false。 TODO: 怎么用?

Session配置 ...

Log配置 ...
```
