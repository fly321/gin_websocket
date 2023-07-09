package routers

import (
	"gin-websocket/services"
	"github.com/gin-gonic/gin"
)

func WsRouter(r *gin.Engine) {
	r.GET("/ws", services.WsRouter{}.IndexService)
	// 接收消息
	/*r.GET("/ws", func(context *gin.Context) {
		//services.WsRouter{}.IndexService(context)

	})*/
}
