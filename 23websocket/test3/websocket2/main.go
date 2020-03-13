package main

import (
	"code.oldbody.com/studygolang/mylearn/23websocket/test3/websocket2/impl"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

//websocket 服务端进阶2-封装Conn
//websocket不是线程安全的，channal是线程安全的

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
		wsconn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	//Upgrade:websocket
	if wsconn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	//初始化Connection
	if conn, err = impl.InitConnection(wsconn); err != nil {
		goto ERR
	}
	//创建一个goroutine 每秒给websocket写入一次心跳
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	//循环读取websocket接收和发送的数据
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	//TODO:关闭连接操作
	conn.Close()
}

func main() {
	//访问: 打开websocket.html
	//配置路由，定义一个回调函数
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
