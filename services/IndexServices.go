package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
)

type WsRouter struct {
}

var upgrade = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

type RequestJson struct {
	Action string        `json:"action"`
	Data   []interface{} `json:"data"`
}

// IndexService ...

/**
* @api {get} /ws 服务端
 * @apiName ws
 * @apiGroup ws
 * @apiVersion 1.0.0
 * @apiDescription 服务端
 * @apiParam {string} action 动作
 * @apiParam {string} data 数据
 * @apiParamExample {json} 请求样例：
 * {
 *     "action": "LogAdd",
 *     "data": [
 *         {
 *             "id": 1,
 *             "name": "张三"
 *         }
 *     ]
 * }
 * @apiSuccessExample {json} 返回样例:
 * {
 *     "code": 200,
 *     "data": [],
 *     "message": "成功"
 * }
*/
func (Wss WsRouter) IndexService(c *gin.Context) {
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		var err = ws.Close()
		if err != nil {
			log.Println(err)
		}
	}(ws)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var RequestJson = &RequestJson{}
		err = json.Unmarshal(message, RequestJson)
		if err != nil {
			log.Println("json.Unmarshal:", err)
			break
		}

		var messageService = &MessageService{}
		vm := reflect.ValueOf(messageService)

		name := vm.MethodByName(RequestJson.Action)
		if name.IsValid() {
			// 反射调用
			var params = []reflect.Value{
				reflect.ValueOf(ws),
				reflect.ValueOf(RequestJson),
			}
			// 调用方法
			name.Call(params)
		} else {
			err := ws.WriteMessage(1, []byte("没有找到方法"))
			if err != nil {
				return
			}
		}

	}
}
