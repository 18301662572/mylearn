package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//cookie 示例

//UserInfo 用户表映射
type UserInfo struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
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
			//登陆成功
			//1.设置Cookie
			c.SetCookie("username", u.UserName, 20, "/", "127.0.0.1", false, true)
			//跳转页面
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

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", cookieMiddleware, vipHandler)
	r.Run()
}
