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