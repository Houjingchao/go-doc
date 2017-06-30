package filters

import (
	"github.com/cocotyty/summer"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/context"
	"io/ioutil"
	"go-doc/consts"
)

func init() {
	summer.Put(new(Token))
}

type Token struct {
	DB *sqlx.DB          `sm:"@.DBApp"`
}

func (t *Token) Empty(ctx context.Context) {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil { //暂时不处理
		return
	}
	ctx.Values().Set(consts.Request, body)
	ctx.Next()
}
