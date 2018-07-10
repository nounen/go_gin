package controllers

import (
	// TODO: 如果只引入 models 下指定的文件是否执行效率更高? eg: 如只导入 todo 模型 `todo "quickstart/models"`
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

// Prepare 优先执行于其他方法
func (c *TodoController) Prepare() {
	beego.Debug("Prepare: 优先执行于下面这些函数")
}

// Index 列表
func (c *TodoController) Index() {
	c.Data["todo"] = c.getTodo(defaultId)
	c.Data["todos"] = c.getTodos()
	c.TplName = "todo/index.tpl"
}

// Store 创建数据 TODO: 此处应添加数据库校验
func (c *TodoController) Store() {
	if _, err := models.AddTodo(c.getTodoFromPOST()); err == nil {
		c.Data["message"] = "添加成功"
	} else {
		c.Data["message"] = "添加失败"
	}

	c.Data["todo"] = c.getTodo(defaultId)
	c.Data["todos"] = c.getTodos()
	c.TplName = "todo/index.tpl"
}

// Show 查看详情
func (c *TodoController) Show() {
	id := c.getId()
	c.Data["id"] = id
	c.Data["todo"] = c.getTodo(id)
	c.Data["todos"] = c.getTodos()
	c.TplName = "todo/index.tpl"
}

// Update 更新数据
func (c *TodoController) Update() {
	todo := c.getTodoFromPOST()
	todo.Id = c.getId()
	models.UpdateTodoById(todo)

	url := "/todo/" + c.getIdString()
	c.Ctx.Redirect(302, url)
}

// TODO: 以下的业务能提取出来吗? 比如提取到 services 目录下
// getId 获取路由上的 :id 参数, 并做类型转换
func (c *TodoController) getId() int64 {
	// string 转 int64, 如何直接拿到想要类型呢
	id, _ := strconv.ParseInt(c.getIdString(), 10, 64)

	return id
}

// getIdString 获取路由上的 :id 参数
func (c *TodoController) getIdString() string {
	return c.Ctx.Input.Param(":id")
}

// getTodoFromPOST 获取表单提交数据
func (c *TodoController) getTodoFromPOST() *models.Todo {
	todo := &models.Todo{}
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
