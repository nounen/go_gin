package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/models"
)

type TodoController struct {
	beego.Controller
}

// TODO: 控制器的 init 如何使用
//func (c *TodoController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
//	c.Data["message"] = ""
//}

// TODO: 此处应为数据库查询, 且分页
func (c *TodoController) getTodos() []interface{} {
	todos, _ := models.GetAllTodo(
		map[string]string{},
		[]string{"Id", "Title", "Sort", "Status"},
		[]string{"sort"},
		[]string{"asc"},
		0,
		10,
	)

	return todos
}

func (c *TodoController) Index() {
	c.Data["todos"] = c.getTodos()
	c.Data["message"] = ""
	c.TplName = "todo/index.tpl"
}

func (c *TodoController) getTodoFromPOST() models.Todo {
	todo := models.Todo{}
	todo.Title = c.GetString("title")
	todo.Sort, _ = c.GetInt("sort", 0)
	todo.Status, _ = c.GetInt("status", 1)
	return todo
}

// TODO: 此处应添加数据库校验
func (c *TodoController) Store() {
	todo := c.getTodoFromPOST()

	beego.Debug(todo)

	if _, err := models.AddTodo(&todo); err == nil {
		c.Data["message"] = "添加成功"
	} else {
		c.Data["message"] = "添加失败"
	}

	c.Data["todos"] = c.getTodos()

	c.TplName = "todo/index.tpl"
}
