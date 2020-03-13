## Socket 和 Websocket
网站：https://blog.csdn.net/dodod2012/article/details/81744526<br>
 **Socket: tcp/udp协议的抽象，底层socket （test1）<br>**
 **Websocket: http实现长连接的协议，是tcp协议的上层协议，在应用层的。(test2)<br>**
 **Websocket: http->websocket 学习进阶1，2(test3)<br>**
 **websocket协议交互，见图"image/"，<br>
 即 客户端发起请求，在http协议头部添加一个 upgrade，服务端发现客户端 http头部中的upgrade后，
 返回给客户端一个 switching, 代表服务端允许客户端将http协议转成websocket协议。达成一致后，底层的tcp仍然保持连接状态，
 双方发送message消息<br>**