package router

import (
	"context"
	"github.com/youngsailor/websocket-example/controller"
	"github.com/youngsailor/websocket/iface"
)

func WebsocketRouterInit(ctx context.Context, wsServer iface.IServer) {
	wsServer.AddRouter(ctx, "print", &controller.PrintRouter.Print)
}
