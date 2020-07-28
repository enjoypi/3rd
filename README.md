# Template for cobra

input in term

```bash
go build -v && ./3rd --config.file=app.yaml child --child.str=changed
```

out

```text
INFO[2019-01-17T18:38:09+08:00] reading from 127.0.0.1:2379
WARN[2019-01-17T18:38:09+08:00] current log level: info
INFO[2019-01-17T18:38:09+08:00] settings on child: {Config:{File:app.yaml} Child:{Bool:true Test:false Str:changed}}
```

## TODO

* 配置工具——配置数据导出、配置校验、配置类型生成
* Deepcopy
* Router
* MQ
* 1、可重入
  2、异步
  3、网络
  4、单元测试
  5、测试目标：正确性、希望尽量的覆盖全，可用性、压力测试，基本的功能就行了
  
  单元测试的原则：
  代码覆盖率
  算法层面
  单元测试的效率、想法
  组件层的测试
  Controller的测试维护性较差
  Module的Mock
  为测试执行创建玩家的
  底层可独立测试，可测试性
  
  FW先出unittest
  package的公开接口应该测试
  
  
  
  dirty，diff merge bson 中间交互消息放在里面 deepcopy