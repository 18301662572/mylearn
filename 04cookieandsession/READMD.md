# Cookie和Session
  [博客：https://www.liwenzhou.com/posts/Go/Cookie_Session/]<br>
  
**A.Cookie**  
```text
HTTP请求是无状态的，服务端让用户的客户端（浏览器）保存一小段数据

Cookie作用机制:
    1.是由服务端保存在客户端的键值对数据（客户端可以阻止服务端保存Cookie）
    2.每次客户端发请求的时候会自动携带该域名下的Cookie
    3.不同域名间的Cookie是不能共享的

Go操作Cookie
    net/http包	
    查询Cookie:  http.Cookie("key")
    设置Cookie: http.SetCookie(w  http.ResponseWriter,cookie *http.Cookie)

Gin框架操作Cookie
    查询Cookie：`c.Cookie("key")`
    设置Cookie：`c.SetCook("key","value",domain,path,maxAge,secure,httpOnly)`

Cookie的应用场景
    1.保存HTTP请求的状态
        1.保存用户登录的状态
        2.保存用户购物车的状态
        3.保存用于定制化的需求
    2.Cookie的缺点
        1.数据量最大4K
        2.保存在客户端（浏览器），不安全
总结：
1.是保存在浏览器端的键值对（保存在客户机上的文本信息）
2.用途：标识请求的，弥补HTTP请求是无状态的缺陷！场景：自定义页面效果，登录，购物车...
3.特点
    1.浏览器发送请求的时候，自动把携带该站点之前存储的Cookie信息
    2.服务端可以设置Cookie数据
    3.Cookie是针对单个域名的，不同域名之间的Cookie是独立的
    4.Cookie数据可以配置过期时间，过期的Cookie数据会被系统清除。
4.Go操作Cookie
	net/http包
```

**B.Session**
```go
前提：用户登录成功后，我们在服务端为每个用户创建一个特定的session和  唯一 的标识，他们一一对应。
其中：
1.Session是在服务端保存的一个数据结构，用来跟踪用户的状态，这个数据可以保存在集群、数据库、文件中；
2.唯一标识通常为 Session ID 会写入用户的 Cookie 中**这样该用户后续再次访问时，请求会自动携带Cookie数据（其中包含Session ID）,服务器通过该（Seesion ID）就能找到与之对应的Session数 据，也就知道来的是“谁”

1.是保存在服务端的键值对（内存、关系型数据库、Redis、文件）
2.必须依赖于Cookie 才能使用
注：Session和Cookie相比的优势：
    1.数据量不受限
	2.数据是保存在服务端，是相对安全的
    3.但是需要后端维护一个Session服务

Session:保存在服务端的键值对数据
Session的存在必须依赖于Cookie，Cookie中保存了每个用户Seseion的唯一标识
Session的特点
    1.保存在服务端，数据量可以存很大（只要服务器支持）
    2.保存在服务端也相对保存在客户端更安全
    3.需要自己去维护一个Session服务，会提高系统的复杂度
Session的工作流程
见图 “images/*”

```
  ​		