**logtransfer工作流程**

#### log transfer===> 从kafka里消费日志数据；写入ES数据库里（index）；Kibana从ES获取数据展示日志信息

“images/*”是systransfer的工作流程<br>


日志收集系统架构图（LogAgent -> Kafka ->log transfer->ES->Kibana->浏览器）<br>
系统信息收集架构图 （SysAgent（gopsutil包收集服务器性能信息）->Kafka->sys transfer->influxDB->Grafana->浏览器） 