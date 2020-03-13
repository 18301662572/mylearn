package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

//websocket 服务端进阶1

var (
	upgrade = websocket.Upgrader{
		//握手过程中，总是允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//在回调函数中，将http头部添加upgrade,转换成websocket协议
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	//Upgrade:websocket
	if conn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	//websocket.Conn 长连接
	//进行数据收发
	for {
		//message类型：Text,Binary
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}

func main() {
	//访问: 打开websocket.html
	//配置路由，定义一个回调函数
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
