package ginsession

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
	"sync"
	"github.com/go-redis/redis"
	"time"
)

//redis 版session服务

//ReidsSessionData 表示一个具体的用户Session数据
type ReidsSD struct {
	SessionID string
	Data      map[string]interface{}
	rwLock    sync.RWMutex //读写锁，锁的是上面的Data
	expired int//过期时间
	client *redis.Client//redis链接
}

//NewRedisSessionData 构造函数
func NewRedisSessionData(id string,client *redis.Client) SessionData {
	return &ReidsSD{
		SessionID: id,
		Data:      make(map[string]interface{}, 8),
		client:client,
	}
}

func (r *ReidsSD) GetSessionID()string{
	return r.SessionID
}

func (r *ReidsSD) Get(key string) (value interface{}, err error){
	//获取读锁
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	value, ok := r.Data[key]
	if !ok {
		err = fmt.Errorf("invaild Key")
		return
	}
	return
}
func (r *ReidsSD) Set(key string, value interface{}){
	//写入锁
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	r.Data[key] = value
}
func (r *ReidsSD) Del(key string){
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	delete(r.Data, key)
}
func (r *ReidsSD) Save(){
	//将最新的session data 保存到redis中
	value,err:=json.Marshal(r.Data)
	if err!=nil{
		//序列化session data 失败
		fmt.Printf("marshal session data failed,err:%v\n",err)
	}
	//将数据保存到redis
	r.client.Set(r.SessionID,value,time.Second*time.Duration(r.expired))
}

//设置过期时间
func (r *ReidsSD) SetExpire(expired int){
	r.expired=expired
}

type RedisMgr struct{
	Session map[string] SessionData
	rwLock  sync.RWMutex
	client *redis.Client  //redis 连接池
}

//NewRedisMgr RedisMgr的构造函数
func NewRedisMgr()(mgr Mgr){
	return &RedisMgr{
		Session:make(map[string] SessionData, 1024), // 初始化1024红色的小框用来存取用户的session data
	}
}

func (r *RedisMgr)Init(addr string,options ...string){
	//初始化redis 连接
	var(
		password string
		db string
	)
	if len(options)==1{
		password=options[0]
	}else if len(options)==2{
		password=options[0]
		db=options[1]
	}
	dbValue,err:= strconv.Atoi(db)
	if err!=nil{
		dbValue=0
	}
	redis.NewClient(&redis.Options{
		Addr:addr,
		Password:password,
		DB:dbValue,
	})
	_,err=r.client.Ping().Result()
	if err!=nil{
		panic(err)
	}
}

func (r *RedisMgr)loadFormRedis(sessionID string)(err error)  {
	//1.连接redis
	value,err:= r.client.Get(sessionID).Result()
	if err!=nil{
		//redis 中没有该session_id对应的session data
		return
	}
	err= json.Unmarshal([]byte(value),&r.Session)
	if err!=nil{
		//从redis 取出来的数据反序列化失败
		return
	}
	//2.根据sessionID 找到对应的数据
	//3.把数据取出来反序列化到r.data
	return
}

//GetSessionData 获取sessionid对应的session data
func (r *RedisMgr) GetSessionData(sessionID string) (sd SessionData, err error) {
	//1.r.Session 中必须已经从Redis里面加载出来数据
	//2.r.Session[sessionID] 拿到对应的session data
	if r.Session==nil{
		err=r.loadFormRedis(sessionID)
		if err!=nil{
			return nil,err
		}
	}
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	sd,ok:=r.Session[sessionID]
	if !ok{
		err=fmt.Errorf("invaild session id")
		return
	}
	return
}

func (r *RedisMgr) CreateSession() (sd SessionData) {
	//1.造一个SessionID
	uuidObj,err:= uuid.NewV4()
	if err!=nil{
		panic("create invalid session id")
	}
	//2.造一个和他对应的SessionData
	sd =NewRedisSessionData(uuidObj.String(),r.client)
	r.Session[sd.GetSessionID()]=sd //把新创建的session data 保存到大仓库中
	//3.返回SessionData
	return
}


