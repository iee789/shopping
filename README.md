### 简介及技术栈

本项目实现了一个电商的核心功能内容包括：用户登录、用户注册、jwt鉴权、商品分类管理、商品管理、订单管理、购物车管理。

技术栈：mysql+gorm、gin、jwt、viper配置库、swagger api接口、air热加载、分页。

### 导库

加密库bcrypt

````
go get golang.org/x/crypto/bcrypt
````

gin

````
go get github.com/gin-gonic/gin
````

uuid库

````
go get github.com/google/uuid
````

viper包

````
go get github.com/spf13/viper
````

jwt包

````
go get github.com/dgrijalva/jwt-go
````

swagger包

````
go install github.com/swaggo/swag/cmd/swag@v1.6.7
# 项目配置后执行swag init
# 浏览器中输入http://localhost:8080/swagger/index.html
````

### 项目启动

要启动这个项目需要进行以下操作：

#### 修改配置文件

```
Env: "dev"

DatabaseSettings:
  DatabaseURI: "root:123456@tcp(127.0.0.1:3306)/go_database?parseTime=True&loc=Local"
  DatabaseName: "go_database"
  Username: "root"
  Password: "123456"


JwtSettings:
  SecretKey: "golang-tech-stack.com"
```

#### 创建数据库

```
create database go_database
```

#### 访问

```
http://localhost:8080/swagger/index.html
```

