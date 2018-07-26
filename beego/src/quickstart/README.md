## 初始化
```shell
source ../../SetGoPath.sh

go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
go get -u github.com/lib/pq
go get -u github.com/go-sql-driver/mysql

cp conf/app.conf.back conf/app.conf # 并填写数据库信息

bee migrate -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/beego_todo" # 这里要填写数据库信息
bee run
```

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
