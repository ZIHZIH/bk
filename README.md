# 介绍
此项目为 wzh 的练习博客，只有相应的服务端代码，简单提供文章以及用户相应功能。

# 代码目录结构
1. ability目录为一些公用组件 暂时未用 2
2. test目录为平时测试组件目录，不必理会
3. tools目录为一些工具，暂时里面用不到，也不必理会
4. 本项目分为了5个服务：gateway article comment like user 每个目录对应一个单体服务
5. 每个微服务的整体布局为：api目录 cmd目录 config目录 infra目录 internal目录 ，每个微服务项目中会有相应的dockerfile文件负责build镜像
api目录：存放proto文件以及生成的grpc包文件
cmd目录：存放着微服务启动的main文件
config目录：存放着整个微服务会用到的配置文件
infra目录：负责基础设置的一些初始化，比如连接数据库等
internal目录：负责着DDD设计思想的 应用层 领域层 基础层。dao负责的是po持久化对象的落盘; dto负责着一些转换，比如po->dto do->dto;service目录负责应用层，也就是重写相应的grpc接口。

关于mongodb的练习，这里采用文章存储。传入任意字段格式的文章，当有多种属性也可以存储。

jwt鉴权具体实现为利用gin中间件进行token检查，在登陆的时候生成一个token给用户，目前只用在/import路由。

kafka简单引用，当作消息队列使用，在article服务中，当查询文章方法被调用打印日志，这时将打印内容写至kafka，再启动一个协程去读相应读kafka将日志信息打印到相应日志文件。
(假想日志落盘为耗时操作，先写入kafka后续再消费提升接口相应速度)
