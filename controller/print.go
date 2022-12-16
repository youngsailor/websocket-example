/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: print
 * @Date: 2022/12/16 16:24
 */

package controller

import (
	"context"
	"github.com/youngsailor/websocket/iface"
	"github.com/youngsailor/websocket/util"
)

var PrintRouter = &printRouter{}

type printRouter struct {
	Print
}

type Print struct {
	iface.BaseRouter
}

func (p *Print) Handle(ctx context.Context, request iface.IRequest) {
	util.WriteSusJson(request.GetSession(), "hello", "hello")
}
