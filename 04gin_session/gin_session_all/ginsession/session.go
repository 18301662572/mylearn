package ginsession

import (
	"github.com/gin-gonic/gin"
)

// 自己实现的gin框架的session中间件
//Session服务

//课上这个版本存在的问题
//1.调用save 的时候不管修没有修改都会保存数据库，session data 定义一个标志位：r.modifyFlog
//2.从redis中加载数据，会存在一个问题 sync.Once 用来加载一次！
//3.过期时间
//老师开源的代码：https://github.com/Q1mi/ginsession

type Options struct {
	MaxAge int
	Path string
	Domain string
	Secure bool
	HttpOnly bool
}

const (
	SessionCookieName="session_id"  //session_id 在Cookie中对应的名字
	SessionContextName="session"  //session data 在gin上下文中对应的key
)

var (
	//MgrObj 全局的Session管理对象（大仓库）
	MgrObj Mgr
)

type SessionData interface {
	GetSessionID()string//返回自己的SessionID
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{})
	Del(key string)
	Save() //保存
	SetExpire(int) //设置过期时间
}


//Mgr 所有类型的大仓库都应该遵循的接口
type Mgr interface {
	Init(addr string,options ...string)//所有支持的后端都必须实现Init（）来执行具体的链接
	GetSessionData(sessionID string) (sd SessionData, err error)
	CreateSession() (sd SessionData)
}

//InitMgr Mgr构造函数
func InitMgr(name,addr string,options...string){
	switch name {
	case "memory":
		MgrObj=NewMemoryMgr()
	case "redis":
		MgrObj=NewRedisMgr()
	}
	MgrObj.Init(addr,options...)//初始化Mgr
}
//实现一个gin框架的中间件
//所有流经我这个中间件的请求，他的上下文中肯定会有一个Session -> session data
func SessionMiddleware(mgrObj Mgr,options *Options)gin.HandlerFunc{
	if mgrObj==nil{
		panic("must call InitMger before use it.")
	}
	//1. 从请求的Cookie中获取SessionID
	//1.1 取不到SessionID,给这个新用户创建一个新的SessionData,同时分配一个SessionID
	//1.2 取SessionID
	//2.  根据sessionid 去Session 大仓库中取到对应的SessionData
	//3.  如何实现让后续所有的处理请求的方法都能拿到sessiondata
	//3.  利用gin 的c.Set("session",sessiondata)
	return func(c *gin.Context){
		//1. 从请求的Cookie中获取SessionID
		var sd SessionData //session data
		sessionID,err:= c.Cookie(SessionCookieName)
		if err!=nil{
			//1.1 取不到SessionID,给这个新用户创建一个新的SessionData,同时分配一个SessionID
			sd= mgrObj.CreateSession()
			sessionID=sd.GetSessionID()
		}else{
			//1.2 取SessionID
			//2.  根据sessionid 去Session 大仓库中取到对应的SessionData
			sd,err= mgrObj.GetSessionData(sessionID)
			if err!=nil{
				//2.1 根据用户传过来的session_id在大仓库中取不到sessio data
				sd=mgrObj.CreateSession()//sd
				//2.2 更新用户Cookie 中保存的session_id
				sessionID=sd.GetSessionID()
			}
		}
		//3.  如何实现让后续所有的处理请求的方法都能拿到sessiondata
		//3.  利用gin 的c.Set("session",sessiondata)
		c.Set(SessionContextName,sd)
		//在 gin 框架中，要回写Cookie 必须在处理请求的函数返回之前
		//c.SetCookie(SessionCookieName,sessionID,3600,"/","127.0.0.1",false,true)
		c.SetCookie(SessionCookieName,sessionID,options.MaxAge,options.Path,options.Domain,options.Secure,options.HttpOnly)
		c.Next() //执行后续的请求处理方法 c.HTML() 时已经把响应头写好了
	}
}

