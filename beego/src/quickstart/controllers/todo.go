package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/models"
)

type TodoController struct {
	beego.Controller
}

type Todo struct {
	Id int
	Title string
	Sort int
	Status int
}

// TODO: 控制器的 init 如何使用
//func (c *TodoController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
//	c.Data["message"] = ""
//}

// TODO: 此处应为数据库查询, 且分页
func (c *TodoController) getFakerTodos() []Todo {
	todos := []Todo {
		{1, "title1", 1, 1},
		{2, "title2", 2, 1},
	}

	// TODO: 表示不懂的调用这个方法
	//todos, _ := models.GetAllTodo()

	return todos
}

func (c *TodoController) Index() {
	c.Data["todos"] = c.getFakerTodos()
	c.Data["message"] = ""

	c.TplName = "todo/index.tpl"
}

func (c *TodoController) getTodoFromPOST() models.Todo {
	var todo models.Todo
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

	c.Data["todos"] = c.getFakerTodos()

	c.TplName = "todo/index.tpl"
}
