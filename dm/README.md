# GORM DMSQL Driver

## Quick Start

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

## Configuration

```go
import (
"gitee.com/jun3372/driver/dm"
"gorm.io/gorm"
)

var datetimePrecision = 2
func main() {
db, err := gorm.Open(dm.New(dm.Config{
DSN: "dm://SYSDBA:SYSDBA@localhost:5236", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
DefaultStringSize: 256,                   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
DisableDatetimePrecision: true,               // disable datetime precision support, which not supported before MySQL 5.6
DefaultDatetimePrecision: &datetimePrecision, // default datetime precision
DontSupportRenameIndex: true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
DontSupportRenameColumn: true, // use change when rename column, rename rename not supported before MySQL 8, MariaDB
SkipInitializeWithVersion: false, // smart configure based on used version
}), &gorm.Config{})
}
```