package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/todo", &controllers.TodoController{}, "get:Index")
	beego.Router("/todo", &controllers.TodoController{}, "post:Store")
	beego.Router("/todo/?:id:int", &controllers.TodoController{}, "get:Show")
	beego.Router("/todo/?:id:int", &controllers.TodoController{}, "post:Update")
}
