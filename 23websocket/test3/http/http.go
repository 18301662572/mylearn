package main

import "net/http"

func wsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {
	//访问: http://localhost:8080/ws
	//配置路由，定义一个回调函数
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
