## 分分钟 构建 go-api 服务

### 使用
拉取代码到go工作区  
`git clone https://github.com/lofreer/go-api`  

### 目录结构
```
app
    controllers #控制器
    middleware  #中间件
    models      #模型
boorstrap       #启动
config          #配置
pkg             #工具包
routes          #路由
.env.example
main.go         #入口
curd.go        #自动构建模型的CURD接口入口
```

### 开始
首先把 `.env.example` 改为 `.env` 并配置数据库连接，其他信息随情况修改  
默认项目包含2个接口(见路由文件routes/api.go)  
1、登录 `api/admin/login`  
2、登录用户信息 `api/admin/auth`  

初始没有用户数据没法登录  
运行 `go run curd.go admin` 构建CURD控制器  
运行 `go run curd.go admin route` 构建CURD路由  
默认构建路由：  
`api/admin/create (POST)`,`api/admin/update (POST)`,`api/admin/delete (GET)`,`api/admin/info (GET)`,`api/admin/paginate (POST)`。  

完成后 运行 `go run main.go` 启动服务  
使用 `api/admin/create` 接口 POST参数:email,password,name 创建数据。  

再次使用登录接口`api/admin/login` POST email 和 password 登录。

成功后使用token 访问 `api/admin/auth` 获取登录用户信息。

### 构建CURD工具使用
要使用该工具分3步  
1、创建数据表模型
```go
type Admin struct {
	models.BaseModel
	Email    string `validate:"required,min:6,email"search:"like"`
	Password string `validate:"min:6"pwd:"pwd"`
	Name     string `validate:"required,min:6"search:"like"`
}
//valiedate标签控制create和update的参数验证，valiedate参数详情请看 github.com/thedevsaddam/govalidator
//search标签控制分页列表的查询方式（目前支持:like,=,>,<,!=。like为str%）
```  
2、bootstrap/auto_migrate.go `init` 函数中`MigrateStruct` map数据添加模型  
```
MigrateStruct["admin"] = admin.Admin{}
```
3、此时就可以执行 `go run curd.go admin` 构建CURD控制器了（默认目录app/controllers/model/curd.go,model为你的模型名称小写）,然后构建路由 `go run curd.go admin route`  
注：该工具暂时只能使用单模型，模型关联（hasOne,hasMany等）暂时还不支持，后续可能会加上  
