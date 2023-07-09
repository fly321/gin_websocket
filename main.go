package main

import (
	"gin-websocket/routers"
	"github.com/gin-gonic/gin"
)

func cmdWebSocket(c *gin.Context) {

}
func main() {
	r := gin.Default()
	routers.WsRouter(r)
	r.Run(":8080")
}
