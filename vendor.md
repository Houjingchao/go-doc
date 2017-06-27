# Vendor Tool for go
## 下载vendor
`go get -u github.com/kardianos/govendor`
复制dependencies  `govendor add/update`
### 项目必须在`$GOPATH/src`

- 初始化:

  `govendor init`
- 增加依赖包:

  `govendor add github.com/jmoiron/sqlx`
- 更新一个包到最新：

  `govendor fetch golang.org/x/net/context`