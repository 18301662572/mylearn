package ginsession

import (
	"fmt"
	"github.com/satori/go.uuid"
	"sync"
)

//内存版Session服务
//仅供参考使用

//MemSessionData 表示一个具体的用户Session数据
type MemSD struct {
	SessionID string
	Data      map[string]interface{}
	rwLock    sync.RWMutex //读写锁，锁的是上面的Data
	//过期时间
}

//NewMemorySessionData MenSD 的构造函数
func NewMemorySessionData(id string) SessionData {
	return &MemSD{
		SessionID: id,
		Data:      make(map[string]interface{}, 8),
	}
}

func (m *MemSD) GetSessionID()string{
	return m.SessionID
}

//SessionData支持的操作
//Get 根据key获取值
func (m *MemSD) Get(key string) (value interface{}, err error) {
	//获取读锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	value, ok := m.Data[key]
	if !ok {
		err = fmt.Errorf("invaild Key")
		return
	}
	return
}

//Set 给key设置值
func (m *MemSD) Set(key string, value interface{}) {
	//写入锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	m.Data[key] = value
}

//Del 删除key对应的键值对
func (m *MemSD) Del(key string) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	delete(m.Data, key)
}

//Save 保存session data
func (m *MemSD) Save(){
	return
}

//设置过期时间
func (m *MemSD) SetExpire(expired int){
	return
}

//Mgr 是一个全局的Session管理
type MemoryMgr struct {
	Session map[string]SessionData
	rwLock  sync.RWMutex
}
//NewMemoryMgr 内存版session大仓库的构造函数
func NewMemoryMgr()(Mgr){
	return &MemoryMgr{
		Session:make(map[string]SessionData, 1024), // 初始化1024红色的小框用来存取用户的session data
	}
}

func (m *MemoryMgr)Init(addr string,options ...string){
	return
}

//GetSessionData 根据传进来的SessionID找到对应的SessionData
func (m *MemoryMgr) GetSessionData(sessionID string) (sd SessionData, err error) {
	//取之前加锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	sd, ok := m.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return
}

//CreateSession 创建一条Session记录
func (m *MemoryMgr) CreateSession() (sd SessionData) {
	//1.造一个SessionID
	uuidObj,err:= uuid.NewV4()
	if err!=nil{
		panic("create invalid session id")
	}
	//2.造一个和他对应的SessionData
	sd = NewMemorySessionData(uuidObj.String())
	m.Session[sd.GetSessionID()]=sd //把新创建的session data 保存到大仓库中
	//3.返回SessionData
	return
}



