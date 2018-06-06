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