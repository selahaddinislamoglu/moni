package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type websocketController struct {
	websocketService service.Websocket
}

func NewWebsocketController() *websocketController {
	return &websocketController{}
}

func (w *websocketController) Setup(websocketService service.Websocket) {
	w.websocketService = websocketService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (w *websocketController) Connect(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = w.websocketService.Connect(ws)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
