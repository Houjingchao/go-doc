package ctrl

import (
	"github.com/kataras/iris/context"
	"code.aliyun.com/mougew/app-server/domain/resp"
)

type Login struct {
}

func (lg *Login) Login(ctx context.Context) {
	ctx.JSON(resp.Forbidden.Resp())
}
