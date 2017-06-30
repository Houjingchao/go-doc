package router

import (
	"go-doc/ctrl"
	"github.com/cocotyty/summer"
	"github.com/kataras/iris/core/router"
	"go-doc/filters"
)

func init() {
	summer.Put(&Router{})
}

type Router struct {
	TokenFilter *filters.Token  `sm:"*"`
	Login       *ctrl.Login `sm:"*"`
}

func (r *Router) login(m router.Party) {
	p := m.Party("/login", r.TokenFilter.Empty) //handler
	{
		p.Post("/verify", r.Login.Login)
	}
}
func (r *Router) Config(mu router.Party) {
	//添加模块的位置
	r.login(mu)
}
