package common

//CollectEntry 要收集的日志的配置项的结构体
type CollectEntry struct {
	Path  string `json:"path"` //去哪个日志读取日志文件
	Topic string `json:"topic"` //日志文件发往kafka中的哪个topic
}