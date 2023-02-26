## Gin

### Gin安装

1. 下载并安装：

```
go get -u github.com/gin-gonic/gin
```

2. 将gin引入到代码中：

```
import "github.com/gin-goinc/gin"
```

3. （可选）如果使用诸如http.StatusOK之类的常量，则需要引入`net/http`包：

```
import "net/http"
```

4. go env 

```
go env -w 设置环境变量
go env -u 还原环境变量
```

### Goland中GOPATH需要单独在preference中设置

### Gin使用

#### 1.创建一个默认的路由引擎

```go
r. := gin.Default()
```

![image-20221126183016734](/Users/heyuyang/Library/Application Support/typora-user-images/image-20221126183016734.png)