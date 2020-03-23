# grpc token认证
https://www.jishuchi.com/read/GO/695<br>
**CA根证书： 针对每个GRPC链接的认证。<br>**
即当客户端访问服务器时，只对每次访问的地址进行认证跟https类似<br>
**Token： 针对每个GRPC 方法 的认证，基于用户Token对不同的方法访问进行权限管理。<br>**
即当用户进行某个操作时，验证是否有操作权限。<br>

```text
8grpc-token              GRPC token认证 （对每个grpc方法进行权限认证）/（CS根证书是对每个grpc链接认证）
    authentication       用户名密码认证
    main                 grpc.Dial()将用户名密码传入
    pb                   grpc请求,响应结构体
    server               实现grpc方法，首先进行Authentication认证
```