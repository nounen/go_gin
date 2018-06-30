package controllers

import (
	"quickstart/models"
	"strconv"

	"github.com/astaxie/beego"
)

const defaultId int64 = 0
const defaultSort int = 0
const defaultStatus int = 1

// TodoController todo 控制器
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

// Index 列表
func (c *TodoController) Index() {
	c.Data["todo"] = c.getTodo(0)
	c.Data["todos"] = c.getTodos()
	c.TplName = "todo/index.tpl"
}

// Store 创建数据 TODO: 此处应添加数据库校验
func (c *TodoController) Store() {
	todo := c.getTodoFromPOST()

	if _, err := models.AddTodo(&todo); err == nil {
		c.Data["message"] = "添加成功"
	} else {
		c.Data["message"] = "添加失败"
	}

	c.Data["todo"] = c.getTodo(0)
	c.Data["todos"] = c.getTodos()
	c.TplName = "todo/index.tpl"
}

// Show 查看详情
func (c *TodoController) Show() {
	// string 转 int64, 如何直接拿到想要类型呢
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	c.Data["todo"] = c.getTodo(id)
	c.Data["todos"] = c.getTodos()
	c.TplName = "todo/index.tpl"
}

// getTodoFromPOST 获取表单提交数据
func (c *TodoController) getTodoFromPOST() models.Todo {
	todo := models.Todo{}
	todo.Title = c.GetString("title")
	todo.Sort, _ = c.GetInt("sort", defaultSort)
	todo.Status, _ = c.GetInt("status", defaultStatus)
	return todo
}

// getTodo 有 id 则从数据库查询, 无责赋予表单字段默认值
func (c *TodoController) getTodo(id int64) *models.Todo {
	if id > 0 {
		todo, _ := models.GetTodoById(id)
		return todo
	}

	todo := &models.Todo{
		Title:  "",
		Sort:   defaultSort,
		Status: defaultStatus,
	}
	return todo
}
