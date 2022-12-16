/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: main
 * @Date: 2022/11/15 17:35
 */

package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/olahol/melody"
	"github.com/youngsailor/websocket-example/handler"
	"github.com/youngsailor/websocket-example/router"
	"github.com/youngsailor/websocket/impl"
	"github.com/youngsailor/websocket/vars"
)

var Websocket = gcmd.Command{
	Name:  "socket",
	Usage: "mini socket",
	Brief: "mini start socket server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		server := g.Server()
		httpRouter := server.Group("/")
		vars.WsHub = melody.New()
		wsServer := impl.NewServer(ctx)
		wsHandler := handler.NewHandlerImpl(vars.WsHub, wsServer)
		// biz router
		router.WebsocketRouterInit(ctx, wsServer)

		httpRouter.GET("/ws", wsHandler.HandleHttpRequest)

		vars.WsHub.HandleConnect(wsHandler.HandleConnect)
		vars.WsHub.HandleMessage(wsHandler.HandleMessage)
		vars.WsHub.HandleClose(wsHandler.HandleClose)
		vars.WsHub.HandleError(wsHandler.HandleError)
		vars.WsHub.HandleDisconnect(wsHandler.HandleDisconnect)
		vars.WsHub.HandleClose(wsHandler.HandleClose)

		server.SetPort(9999)
		server.Run()
		return nil
	},
}

func main() {
	Websocket.Run(gctx.New())
}
