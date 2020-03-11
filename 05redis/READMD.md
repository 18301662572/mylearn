#Redis
**http://www.redis.cn/**   
**https://www.runoob.com/redis/redis-install.html**

**1.是一个开源的内存数据库，提供了5种不同类型的数据结构**<br>
2.通过复制、持久化和客户端分片等特性，我们可以很方便的将Redis扩展成一个能够包含数百GB数据、每秒处理上百万次请求的系统<br>
3.数据结构：STRING(字符串)、LIST(列表)、SET(集合)、HASH(哈希)、ZSET(有序集合)<br>
4.应用场景<br>
```go
a.缓存系统，减轻主数据库（MYSQL）的压力
b.计数场景，比如微博、抖音中的关注数和粉丝数
c.热门排行榜、需要排序的场景特别合适使用ZSET
d.利用LIST可以实现队列的功能
```
5.Redis与Memcached比较<br>
```go
Memcached中的值只支持简单的字符串，Redis支持更丰富的5种数据结构类型。	
Redis性能比Memcached好很多，Redis支持RDB(快照)持久化和AOF持久化。Redis支持master/slave模式
```
    			