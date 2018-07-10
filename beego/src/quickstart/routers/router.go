package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	FilterMethod()

	beego.Router("/", &controllers.MainController{})

	beego.Router("/todo", &controllers.TodoController{}, "get:Index")
	beego.Router("/todo", &controllers.TodoController{}, "post:Store")
	beego.Router("/todo/?:id:int", &controllers.TodoController{}, "get:Show")
	beego.Router("/todo/?:id:int", &controllers.TodoController{}, "put:Update")
}

// FilterMethod 过滤器, 使请求支持更多 method: put, patch, delete 等
func FilterMethod()  {
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") !=  "" && ctx.Input.IsPost(){
			ctx.Request.Method = ctx.Input.Query("_method")
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
}
