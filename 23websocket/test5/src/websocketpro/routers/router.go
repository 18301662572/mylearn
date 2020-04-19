package routers

import (
	"code.oldbody.com/studygolang/mylearn/23websocket/test5/src/websocketpro/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ServerController{})
}
