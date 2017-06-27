# go model并创建sql

### 定义struct
```go
type Student struct {
	Id   int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Sex  string `db:"sex" json:"sex"`
}
```
### 建立相关的数据库操作