## 安装 Beego 和 Bee 的开发工具:
```
$ go get -u github.com/astaxie/beego
$ go get -u github.com/beego/bee
```

## $GOPATH
```
# 如果您还没添加 $GOPATH 变量
$ echo 'export GOPATH="$HOME/go"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件

# 如果您已经添加了 $GOPATH 变量
$ echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件
$ exec $SHELL
```

## 简单示例
* quick_start.go:
    ```go
    package main

    import (
        "github.com/astaxie/beego"
    )

    type MainController struct {
        beego.Controller
    }

    func (this *MainController) Get() {
        this.Ctx.WriteString("hello world")
    }

    func main() {
        beego.Router("/", &MainController{})
        beego.Run()
    }
    ```

* 编译执行:
    ```
    go build -o hello quick_start.go
    ./hello
    ```

* 更多介绍: https://beego.me/quickstart

* ab 测试: `$ ab -n 10000 -c 100 http://127.0.0.1:8080/`


## bee 工具简介
* 安装 bee 工具: `go get github.com/beego/bee`
    * bee 可执行文件默认存放在 $GOPATH/bin 里面

* `bee new <项目名>` 就可以创建一个新的项目 (web 项目)
    * 但是注意该命令必须在 `$GOPATH/src` 下执行。最后会在 `$GOPATH/src` 相应目录下生成如下目录结构的项目

* `bee api <项目名>` 就可以创建一个新的项目 (api 项目)

* `bee run <项目路径>`: bee run 命令是监控 beego 的项目，通过 fsnotify监控文件系统 (也就是文件改动会自己编译)

* 更多 bee 命令:
    ```go
    bee pack
    bee bale
    bee version
    bee generate
    bee migrate
    ```


## 快速入门
* 设置多个 gopath 目录
    * 在当前目录创建 `src` 目录, 执行 `source SetGoPath.sh`
    * 进入 `src` 目录 执行 `bee new xxx` 创建新项目

1. bee 工具新建项目
    * 进入 `src` 目录 执行 `bee new quickstart`

    * main.go 是入口文件

    * `cd quickstart; bee run` 运行项目 

2. 路由设置
    * 入口文件 main.go, 它引入了路由文件 `_ "quickstart/routers"`, 此时会自动执行 routers.go 的 `init()` 函数

    * 再回来看看 main 函数里面的 beego.Run， beego.Run 执行之后，我们看到的效果好像只是监听服务端口这个过程，但是它内部做了很多事情：
        * 解析配置文件: conf/app.conf` (开启的端口，是否开启 session，应用名称等信息)
        * 执行用户的 hookfun
        * 是否开启 session
        * 是否编译模板
        * 是否开启文档功能
        * 是否启动管理模块
        * 监听服务端口
    
    * 一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 `8080` 作为对外服务，另一个 `8088` 端口实行对内监控。

3. controller 运行机制

4. model 逻辑

5. view 渲染

6. 静态文件处理
