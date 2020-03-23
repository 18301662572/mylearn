package pb

import (
	"fmt"
	"github.com/moby/moby/pkg/pubsub" //"github.com/docker/docker/pkg/pubsub"
	"strings"
	"time"
)

//docker项目中提供了一个pubsub的极简实现的demo

//下面是基于pubsub包实现的本地发布订阅代码：
//其中pubsub.NewPublisher构造一个发布对象，p.SubscribeTopic()可以通过函数筛选感兴趣的主题进行订阅。

func main() {
	//构造一个发布对象
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	//订阅golang:开头的主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	//订阅docker:开头的主题
	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})
	//发布信息
	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(1)
	go func() {
		fmt.Println("golang topic:", <-golang)
	}()
	go func() {
		fmt.Println("docker topic:", <-docker)
	}()
	//阻塞
	<-make(chan bool)
}
