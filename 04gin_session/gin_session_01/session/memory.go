package session

import "fmt"

//内存版Session服务

//SessionData支持的操作

//Get 根据key获取值
func (s *SessionData) Get(key string) (value interface{}, err error) {
	//获取读锁
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value, ok := s.Data[key]
	if !ok {
		err = fmt.Errorf("invaild Key")
		return
	}
	return
}

//Set 给key设置值
func (s *SessionData) Set(key string, value interface{}) {
	//写入锁
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	s.Data[key] = value
}

//Del 删除key对应的键值对
func (s *SessionData) Del(key string) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	delete(s.Data, key)
}
