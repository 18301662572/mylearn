package tailfile

import (
	"code.oldbody.com/studygolang/mylearn/11logagentallip/kafka"
	"context"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/prometheus/common/log"
	"strings"
	"time"
)

//tail相关

type tailTask struct {
	path   string
	topic  string
	tObj   *tail.Tail
	ctx    context.Context
	cancel context.CancelFunc
}

func newTailTask(path, topic string) *tailTask {
	ctx, cancel := context.WithCancel(context.Background())
	return &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (t *tailTask) Init() (err error) {
	//tail 配置
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	//打开文件获取数据
	t.tObj, err = tail.TailFile(t.path, cfg)
	return
}

//读取日志发往kafka
func (t *tailTask) run() {
	log.Infof("collect for path %s is running ...", t.path)
	for {
		select {
		//这个task停掉，之后的操作不执行
		case <-t.ctx.Done(): //只要调用t.cancel()函数，t.ctx.Done()里面就有值
			log.Infof("path:%s is stopping...", t.path)
			return
		//循环读数据
		case line, ok := <-t.tObj.Lines: //chan tail.Line
			if !ok {
				log.Warnf("tail file close reopen,filename:%s\n", t.path)
				time.Sleep(time.Second) //读取出错，等一秒
				continue
			}
			//如果是空行就略过
			if len(strings.Trim(line.Text, "\r")) <= 0 {
				log.Info("出现空行，直接跳过...")
				continue
			}
			log.Infof("msg:", line.Text)
			//利用通道将同步的代码改为异步的
			//把读出来的一行日志包装成kafka里面的msg类型，丢到通道中
			msg := &sarama.ProducerMessage{}
			msg.Topic = t.topic
			msg.Value = sarama.StringEncoder(line.Text)
			//丢到通道中
			kafka.ToMsgChan(msg)
		}
	}
}
