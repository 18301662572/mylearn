package handlers

import (
	"code.oldbody.com/studygolang/mylearn/02mysqlx/db"
	"code.oldbody.com/studygolang/mylearn/02mysqlx/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//HandlerFunc

//UserListHandler 查看用户列表Handler
func UserListHandler(c *gin.Context) {
	//查询全部数据
	list, err := db.QueryUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err,
		})
		return
	}
	fmt.Printf("list:%v\n", list)
	if list == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "list is nil",
		})
		return
	}
	//返回给浏览器
	c.HTML(http.StatusOK, "templates/userlist.html", gin.H{
		"code": 0,
		"data": list,
	})
}

//AddUserHandler 跳转到添加用户界面
func AddUserHandler(c *gin.Context) {
	//返回给用户html界面

	// 通过HTML函数返回html代码
	// 第二个参数是模版文件名字
	// 第三个参数是map类型，代表模版参数 gin.H 是map[string]interface{}类型的别名
	c.HTML(http.StatusOK, "templates/adduser.html", nil)
}

//添加用户
func CreateUserHandler(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err == nil {
		fmt.Printf("u:age=%v name=%v\n", u.Age, u.Name)
		_, err := db.AddUser(u.Age, u.Name)
		if err != nil {
			msg := "添加用户失败！"
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
			return
		}
		//跳转到userlist页面  StatusMovedPermanently 301:重定向
		c.Redirect(http.StatusMovedPermanently, "/user/list")
	} else {
		msg := "获取表单数据失败！"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}

	//nameVal := c.PostForm("name")
	//参数1 数字的字符串形式 参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制 参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	//ageVal, _ := strconv.ParseInt(c.PostForm("age"), 10, 64)
	//_, err := db.AddUser(ageVal, nameVal)
	//if err != nil {
	//	msg := "添加用户失败！"
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": msg,
	//	})
	//	return
	//}
	////跳转到userlist页面  StatusMovedPermanently 301:重定向
	//c.Redirect(http.StatusMovedPermanently, "/user/list")
}

//添加用户界面
func AddUserAnyHandler(c *gin.Context) {
	//c.Request.Method 判断是 post 还是 get
	if c.Request.Method == "POST" {
		//也可以用c.ShouldBind(&u) 将postform参数反射到model.User 中，同（CreateUserHandler）
		nameVal := c.PostForm("name")
		ageVal, _ := strconv.ParseInt(c.PostForm("age"), 10, 64)
		_, err := db.AddUser(ageVal, nameVal)
		if err != nil {
			msg := "添加用户失败！"
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
			return
		}
		//跳转到userlist页面  StatusMovedPermanently 301:重定向
		c.Redirect(http.StatusMovedPermanently, "/user/list")
	} else {
		// 通过HTML函数返回html代码
		// 第二个参数是模版文件名字
		// 第三个参数是map类型，代表模版参数 gin.H 是map[string]interface{}类型的别名
		c.HTML(http.StatusOK, "templates/adduser.html", nil)
	}
}

//跳转到修改用户界面
func UpdateUserHandler(c *gin.Context) {
	//获取query数据
	theid, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取修改用户id失败",
		})
		return
	}
	//获取用户信息
	u, err := db.QueryUser(theid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取修改用户信息失败",
		})
		return
	}
	//返回给浏览器
	c.HTML(http.StatusOK, "templates/updateuser.html", gin.H{
		"ID":   u.ID,
		"Name": u.Name,
		"Age":  u.Age,
	})
}

//修改用户信息
func UpdateUserPostHandler(c *gin.Context) {
	idVal, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	nameVal := c.PostForm("name")
	ageVal, _ := strconv.ParseInt(c.PostForm("age"), 10, 64)
	_, err := db.UpdateUser(nameVal, ageVal, idVal)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "修改用户失败",
		})
		return
	}
	//跳转到userlist页面
	c.Redirect(http.StatusMovedPermanently, "/user/list")
}

//删除用户信息
func DeleteUserHandler(c *gin.Context) {
	//获取query id
	theid, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	//删除用户
	_, err := db.DeleteUser(theid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除用户失败",
		})
		return
	}
	//刷新本界面
	c.Redirect(http.StatusMovedPermanently, "/user/list")
}
