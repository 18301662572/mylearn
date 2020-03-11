# gin框架的Session扩展

1.Session中间件<br>
见图 "images/*"
```text
内存版Session中间件实现
Session服务分成两部分：
    1.大仓库:   单例模式
    2.session data： 每个用户对应自己的session data 
    3.代码练习：

    老师开源代码：https://github.com/Q1mi/ginsession
```