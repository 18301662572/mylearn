package controllers

import (
	"code.oldbody.com/studygolang/mylearn/23websocket/test5/src/websocketpro/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type ServerController struct {
	beego.Controller
}

var (
	clients = make(map[models.Client]bool) // 用户组映射
	// 此处要设置有缓冲的通道。因为这是goroutine自己从通道中发送并接受数据。
	// 若是无缓冲的通道，该goroutine发送数据到通道后就被锁定，需要数据被接受后才能解锁，而恰恰接受数据的又只能是它自己
	join    = make(chan models.Client, 10)  // 用户加入通道
	leave   = make(chan models.Client, 10)  // 用户退出通道
	message = make(chan models.Message, 10) // 消息通道Name
)

// 用于与用户间的websocket连接(chatRoom.html发送来的websocket请求)
func (c *ServerController) WS() {
	name := c.GetString("name")
	if len(name) == 0 {
		beego.Error("name is NULL")
		c.Redirect("/", 302)
		return
	}

	// 检验http头中upgrader属性，若为websocket，则将http协议升级为websocket协议
	conn, err := (&websocket.Upgrader{}).Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

	if _, ok := err.(websocket.HandshakeError); ok {
		beego.Error("Not a websocket connection")
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	var client models.Client
	client.Name = name
	client.Conn = conn

	// 如果用户列表中没有该用户
	if !clients[client] {
		join <- client
		beego.Info("user:", client.Name, "websocket connect success!")
	}

	// 当函数返回时，将该用户加入退出通道，并断开用户连接
	defer func() {
		leave <- client
		client.Conn.Close()
	}()

	// 由于WebSocket一旦连接，便可以保持长时间通讯，则该接口函数可以一直运行下去，直到连接断开
	for {
		// 读取消息。如果连接断开，则会返回错误
		_, msgStr, err := client.Conn.ReadMessage()

		// 如果返回错误，就退出循环
		if err != nil {
			break
		}

		beego.Info("WS-----------receive: " + string(msgStr))

		// 如果没有错误，则把用户发送的信息放入message通道中
		var msg models.Message
		msg.Name = client.Name
		msg.EventType = 0
		msg.Message = string(msgStr)
		message <- msg
	}
}

//后端广播功能
//将发消息、用户加入、用户退出三种情况都广播给所有用户。
// 后两种情况经过处理，转换为第一种情况。真正发送信息给客户端的，只有第一种情况。
func broadcaster() {
	for {
		// 哪个case可以执行，则转入到该case。若都不可执行，则堵塞。
		select {
		// 消息通道中有消息则执行，否则堵塞
		case msg := <-message:
			str := fmt.Sprintf("broadcaster-----------%s send message: %s\n", msg.Name, msg.Message)
			beego.Info(str)
			// 将某个用户发出的消息发送给所有用户
			for client := range clients {
				// 将数据编码成json形式，data是[]byte类型
				// json.Marshal()只会编码结构体中公开的属性(即大写字母开头的属性)
				data, err := json.Marshal(msg)
				if err != nil {
					beego.Error("Fail to marshal message:", err)
					return
				}
				// fmt.Println("=======the json message is", string(data))  // 转换成字符串类型便于查看
				if client.Conn.WriteMessage(websocket.TextMessage, data) != nil {
					beego.Error("Fail to write message")
				}
			}

		// 有用户加入
		case client := <-join:
			str := fmt.Sprintf("broadcaster-----------%s join in the chat room\n", client.Name)
			beego.Info(str)

			clients[client] = true // 将用户加入映射

			// 将用户加入消息放入消息通道
			var msg models.Message
			msg.Name = client.Name
			msg.EventType = 1
			msg.Message = fmt.Sprintf("%s join in, there are %d preson in room", client.Name, len(clients))

			// 此处要设置有缓冲的通道。因为这是goroutine自己从通道中发送并接受数据。
			// 若是无缓冲的通道，该goroutine发送数据到通道后就被锁定，需要数据被接受后才能解锁，而恰恰接受数据的又只能是它自己
			message <- msg

		// 有用户退出
		case client := <-leave:
			str := fmt.Sprintf("broadcaster-----------%s leave the chat room\n", client.Name)
			beego.Info(str)

			// 如果该用户已经被删除
			if !clients[client] {
				beego.Info("the client had leaved, client's name:" + client.Name)
				break
			}

			delete(clients, client) // 将用户从映射中删除

			// 将用户退出消息放入消息通道
			var msg models.Message
			msg.Name = client.Name
			msg.EventType = 2
			msg.Message = fmt.Sprintf("%s leave, there are %d preson in room", client.Name, len(clients))
			message <- msg
		}
	}
}

//此处需要利用goroutine并发模式，使得该函数能独立在额外的一个线程上运作。
func init() {
	go broadcaster()
}
