package handlers

import (
	"code.oldbody.com/studygolang/mylearn/04cookieandsession/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//登录处理函数
func LoginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/index")
		var u model.User
		if err := c.ShouldBind(&u); err == nil {
			//判断登录信息
			if strings.ToLower(u.UserName) == "test" && u.PassWord == "123456" {
				//设置Cookie
				c.SetCookie("username", u.UserName, 20, "/", "127.0.0.1", false, true)
				//跳转到index页面
				c.Redirect(http.StatusFound, toPath) //301：StatusMovedPermanently 永久重定向;应该使用302
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "username or password failed",
				})
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "get form data failed",
			})
			return
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

//主界面
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//基于Cookie实现用户登录认证的中间件
func CookieMiddleware(c *gin.Context) {
	//在返回页面之前要先校验是否存在username的Cookie
	//获取Cookie
	username, err := c.Cookie("username")
	if err != nil {
		//直接跳转到登陆页面
		//给中间件做一个登陆成功后，自动跳转到上个页面的功能---使用next拼接当前页面
		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		c.Redirect(http.StatusFound, toPath) //301：StatusMovedPermanently 永久重定向;应该使用302
	}
	//用户已经登陆了
	c.Set("username", username) //在上下文中设置一个键值对
	c.Next()                    //继续后续的处理函数
}

//VIP界面
func VipHandler(c *gin.Context) {
	//获取上下文中的键值对
	tmpusername, ok := c.Get("username")
	if !ok {
		//如果取不到值，说明前面中间件出问题,跳转到登录界面
		c.JSON(http.StatusOK, gin.H{
			"msg": "CookieMiddleware failed",
		})
		c.Redirect(http.StatusFound, "/login") //301：StatusMovedPermanently 永久重定向;应该使用302
	}
	username, ok := tmpusername.(string)
	if !ok {
		//如果取不到值，说明类型断言失败
		c.Redirect(http.StatusFound, "/login") //301：StatusMovedPermanently 永久重定向;应该使用302
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}
