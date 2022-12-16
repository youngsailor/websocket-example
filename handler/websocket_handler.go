/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: websocket_handler
 * @Date: 2022/11/24 9:46
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/olahol/melody"
	"github.com/youngsailor/websocket/iface"
	"github.com/youngsailor/websocket/impl"
	"github.com/youngsailor/websocket/types"
	"net/http"
)

var sfNode, _ = snowflake.NewNode(1)

type handler struct {
	m      *melody.Melody
	server iface.IServer
}

func (h *handler) HandleClose(session *melody.Session, i int, s string) error {
	fmt.Println("close")
	return nil
}

func (h *handler) HandleConnect(session *melody.Session) {
	sessionId := sfNode.Generate()
	session.Set("session_id", sessionId)
}

func (h *handler) HandleDisconnect(session *melody.Session) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleError(session *melody.Session, err error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleMessage(session *melody.Session, bytes []byte) {
	ctx := session.Request.Context()
	var message types.Message
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	newMsg := impl.NewMsg(message.BizType, []byte(message.Data))
	request := impl.NewRequest(newMsg, session, h.m)
	h.server.(*impl.Server).MsgHandler.HandleMsg(ctx, request)
}

func (h *handler) HandleMessageBinary(session *melody.Session, bytes []byte) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandlePong(session *melody.Session) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleSentMessage(session *melody.Session, bytes []byte) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleSentMessageBinary(session *melody.Session, bytes []byte) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return h.m.HandleRequest(w, r)
}

func (h *handler) HandleRequestWithKeys(w http.ResponseWriter, r *http.Request, keys map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleHttpRequest(r *ghttp.Request) {
	err := h.HandleRequest(r.Response.Writer, r.Request)
	if err != nil {
		panic(err)
	}
}

func NewHandlerImpl(m *melody.Melody, server iface.IServer) iface.IConnHandler {
	return &handler{
		m:      m,
		server: server,
	}
}
