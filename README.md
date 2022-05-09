# driver

#### 介绍
> 该项目是基于GORM Driver接口实现各种方言数据库包


#### 安装教程

```shell
# 达梦数据库
g go get gitee.com/jun3372/driver/dm
```

#### 使用说明

```go
import (
"gitee.com/jun3372/driver/dm"
"gorm.io/gorm"
)

func main() {
dsn := "dm://SYSDBA:SYSDBA@localhost:5236"
db, err := gorm.Open(dm.Open(dsn), &gorm.Config{})
}
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
