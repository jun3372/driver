# driver

#### 介绍

> 该项目是基于 GORM Driver 接口实现各种方言数据库包

#### 安装教程

```shell
# 达梦数据库
go get gitee.com/jun3372/driver/dm
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
