package data

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "data.html", gin.H{})
}

// 往下是 Web Socket 代码

// 存储全部客户端连接
var connClient = make(map[*websocket.Conn]bool)

// 去除跨域限制
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端连接
func Ws(c *gin.Context) {
	var err error
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		// 创建 WebSocket 失败
		return
	}

	connClient[ws] = true

	defer ws.Close()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			// 客户端断开
			connClient[ws] = false
			break
		}
	}
}

// 发送消息
func Send(data map[string]interface{}) {
	for k, v := range connClient {
		if v {
			err := k.WriteJSON(data)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}

// 生成数据 JSON
func MakeDataJson(typex string, data map[string]interface{}) map[string]interface{} {
	/**
	//左侧三个：攻击类型汇总 攻击地区(国家)汇总 IP汇总
	//中间两个：攻击地图 最新攻击数据
	//右侧三个：最新账号密码 攻击时间段统计 集群统计

	left_type
	left_country
	left_ip

	center_map
	center_data

	right_account
	right_time
	right_colony
	**/

	result := map[string]interface{}{
		"type": typex,
		"data": data,
	}

	return result
}
