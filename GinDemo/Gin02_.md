## 1. 实现用户注册

### 1.1获取参数&对post接收到的数据进行验证

![image-20221126204134931](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204134931.png)

![image-20221126204156161](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204156161.png)

![image-20221126204220382](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204220382.png)

### 1.2 gorm连接数据库

```
go get -u github.com/jinzhu/gorm
go get 	github.com/go-sql-driver/mysql
```

#### 1.2.1 定义模型

![image-20221126204535363](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204535363.png)

#### 1.2.2创建数据库连接

![image-20221126204758654](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204758654.png)

 ![image-20221126204825613](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204825613.png)

![image-20221126205713553](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126205713553.png)

![image-20221126204910224](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126204910224.png)

#### 1.2.3 校验手机号

![image-20221126205215824](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126205215824.png)

![image-20221126205520372](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126205520372.png) 

## 2.项目重构关注分离

重构1.中的代码让其更具结构性，有更好的维护性

### 2.1.创建git仓库

2.1.1  创建.ignore文件

2.1.2  git init 初始化git仓库

### 2.2 重构代码

![image-20221126211057961](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126211057961.png)

## 3.用户登录

## 4.jwt配合中间件用户认证

```go
go get github.com/dgrijalva/jwt-go
```

