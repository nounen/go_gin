package main

import (
	"github.com/astaxie/beego"
)

// 定义 Controller，这里我们定义了一个 struct 为 MainController，
// 充分利用了 Go 语言的 __组合__ 的概念，匿名包含了 beego.Controller，这样我们的 MainController 就拥有了 beego.Controller 的所有方法
type MainController struct {
	beego.Controller
}

// 这里我们定义了 MainController 的 Get 方法用来重写继承的 Get 函数，这样当用户发起 GET 请求的时候就会执行该函数
func (controller *MainController) Get() {
	controller.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})

	beego.Run()
}
