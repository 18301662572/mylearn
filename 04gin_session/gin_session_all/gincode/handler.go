package main

import (
	"code.oldbody.com/studygolang/mylearn/04gin_session/gin_session_all/ginsession"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//UserInfo 用户表映射
type UserInfo struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

//编写一个校验用户是否登陆的中间件
//起始就是从上下文中取到session data ,从session data 取isLogin
func AuthMiddleware(c *gin.Context){
	//1.先从上下文中获取session data
	tmpSD,_:= c.Get(ginsession.SessionContextName)
	sd:=tmpSD.(ginsession.SessionData)
	value,err:= sd.Get("isLogin")
	if err!=nil{
		fmt.Println(err)
		//取不到就是没有登陆
		c.Redirect(http.StatusFound,"/login")
		return
	}
	fmt.Println(value)
	isLogin,ok:=value.(bool)
	if !ok{
		c.Redirect(http.StatusFound,"/login")
		return
	}
	if !isLogin{
		c.Redirect(http.StatusFound,"/login")
		return
	}
	fmt.Println(isLogin)
	c.Next()
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		//获取中间件中的next参数，跳转到上个页面
		toPath := c.DefaultQuery("next", "/index")
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码不能为空",
			})
			return
		}
		if u.UserName == "guan" && u.PassWord == "123" {
			//登陆成功,在当前这个用户的sessiondata保存一个键值对：isLogin=true
			//1.先从上下文中获取session data
			 tmpSD,ok:= c.Get(ginsession.SessionContextName)
			 if !ok{
			 	panic("session middleware")
			 }
			sd:=tmpSD.(ginsession.SessionData)
			sd.Set("isLogin",true)
			sd.Save()
			//2.给session data 设置isLogin=true
			//跳转到index页面
			c.Redirect(http.StatusMovedPermanently, toPath)
		} else {
			//密码错误
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码有误",
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func homeHandler(c *gin.Context) {
	//在返回页面之前要先校验是否存在username的Cookie
	//获取Cookie
	username, err := c.Cookie("username")
	if err != nil {
		//直接跳转到登陆页面
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}

//基于Cookie实现用户登录认证的中间件
func cookieMiddleware(c *gin.Context) {
	//在返回页面之前要先校验是否存在username的Cookie
	//获取Cookie
	username, err := c.Cookie("username")
	if err != nil {
		//直接跳转到登陆页面
		//给中间件做一个登陆成功后，自动跳转到上个页面的功能---使用next拼接当前页面
		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, toPath)
		return
	}
	//用户已经登陆了
	c.Set("username", username) //在上下文中设置一个键值对
	c.Next()                    //继续后续的处理函数
}

func vipHandler(c *gin.Context) {
	tmpusername, ok := c.Get("username") //获取上下文中的key值
	if !ok {
		//如果取不到值，说明前面中间件出问题
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	username, ok := tmpusername.(string)
	if !ok {
		//如果取不到值，说明类型断言失败
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}
