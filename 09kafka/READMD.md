# kafka 与 kibana

### kafka
1.kafka默认用zookeeper做集群管理  (Kafka Cluster)<br>
2.往kafka里面发消息： 基于sarama第三⽅库开发的kafka client （sarama版本v1.19.0）<br>
3.从kafka里面读消息：tailf包<br>

### zookeeper
zookeeper是一个分布式的，开放源码的分布式应用程序协调服务，是集群的管理者，监视着集群中各个节点的状态根据节点提交的反馈进行下一步合理操作。最终将简单易用的接口和性能高效、功能稳定的系统提供给用户。<br>
**kafka默认用zookeeper做集群管理**<br>
1.服务注册发现 (见图)<br>
2.分布式锁，借助zookeeper的树结构，谁是树节点下的第一个子节点，这个子节点就获取了这个锁.（见图）<br>

见图“images/*”