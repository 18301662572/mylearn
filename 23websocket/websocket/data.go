package main

//将连接中传输的数据抽象出对象
type Data struct {
	Ip string	`json:"ip"`
	//标识信息的类型
	//login:登录
	//handshake:握手信息，刚打开网页的状态
	//system:系统信息，XXX上线了
	//logout:退出
	//user:普通信息
	Type string `json:"type"`
	//代表哪个用户说的
	Form string `json:"form"`
	//传输内容
	Content string `json:"content"`
	//用户名
	User string	`json:"user"`
	//用户列表
	UserList []string `json:"user_list"`
}


