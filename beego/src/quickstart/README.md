## 一个 beego TODO list 案例
* 数据库: 
    * 创建数据库: 迁移文件
    * 模型
    * ORM: 曾 删 改 查, 分页

* 路由

* 控制器

* 业务逻辑

* 模板渲染


### 可参: 打包, 代码生成(models, controller, view, docs)
* https://www.cnblogs.com/zhangboyu/p/7760693.html


### ab 测试 -- 感觉挺客观的, 虽然只是一个列表查询
* ab -n 1000 -c100 http://localhost:8000/todo
