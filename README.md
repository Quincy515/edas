# EDAS

#### 介绍
企业级分布式应用服务EDAS（Enterprise Distributed Application Service）

#### 软件架构
软件架构说明

##### 模块划分

##### 服务划分

##### 数据库图

#### 技术实现

1.  技术方案 参考[MTBSystem](https://github.com/wiatingpub/MTBSystem)
    - 服务端 go-micro
    - 数据库 mysql
    - 缓存   redis
    - 前端   react
    - 服务器 centos7 & nginx
    - 本地环境 go 1.13.4
    - 容器   Docker
    - 进程管理 supervisor
    - 数据库备份 冷备份(rsync+mysqldump)
2.  特性
    - 遵循 RESTful API 设计规范
    - 基于原生 http net 可以方便修改为各种框架和定制服务
    - 基于 Casbin 的 RBAC 访问控制模型
    - 基于 SQL 语句的数据库访问
    - 依赖注入(基于dig)
    - 日志追踪(基于zap)
    - JWT 认证
    - 支持Swagger文档(基于swaggo)
    - 单元测试(基于net/http/httptest包，覆盖所有接口层的测试)
3.  xxxx

#### 使用说明

1.  使用的开源第三方库

如下：
```
go get github.com/jmoiron/sqlx
go get go.uber.org/zap
go get go.uber.org/dig
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/json-iterator/go
go get github.com/stretchr/testify
go get github.com/swaggo/swag
go get google.golang.org/grpc
go get github.com/micro/go-micro
go get github.com/micro/go-plugins
go get github.com/golang/protobuf/proto
go get github.com/golang/protoc-gen-go
go get github.com/micro/protoc-gen-micro
go get github.com/google/uuid
go get github.com/julienschmidt/httprouter
go get github.com/juju/ratelimit
```

2.  项目结构概览

```
├─.gitee
├─.idea
│  └─dictionaries
└─service
    └─account

```

3.  docker部署服务注册中心consul

##### 当前版本
docker: docker-ce19.03.4 consul: 1.6.1

##### consul配置参数说明

```
–net=host docker参数, 使得docker容器越过了net namespace的隔离，免去手动指定端口映射的步骤
-server consul支持以server或client的模式运行, server是服务发现模块的核心, client主要用于转发请求
-advertise 将本机私有IP传递到consul
-retry-join 指定要加入的consul节点地址，失败会重试, 可多次指定不同的地址
-client consul绑定在哪个client地址上，这个地址提供HTTP、DNS、RPC等服务，默认是127.0.0.1
-bind 绑定服务器的ip地址；该地址用来在集群内部的通讯，集群内的所有节点到地址都必须是可达的，默认是0.0.0.0
allow_stale 设置为true, 表明可以从consul集群的任一server节点获取dns信息, false则表明每次请求都会经过consul server leade
-bootstrap-expect 数据中心中预期的服务器数。提供后，Consul将等待指定数量的服务器可用，然后启动群集。允许自动选举leader，但不能与传统-bootstrap标志一起使用, 需要在服务端模式下运行。
-data-dir 数据存放位置，用于持久化保存集群状态
-node 群集中此节点的名称，这在群集中必须是唯一的，默认情况下是节点的主机名。
-config-dir 指定配置文件，当这个目录下有 .json 结尾的文件就会被加载，详细可参考https://www.consul.io/docs/agent/options.html#configuration_files
-enable-script-checks 检查服务是否处于活动状态，类似开启心跳
-datacenter 数据中心名称
-ui 开启ui界面
-join 加入到已有的集群中
```

##### consul端口用途说明
- 8500 : http 端口，用于 http 接口和 web ui访问
- 8300 : server rpc 端口，同一数据中心 consul server 之间通过该端口通信
- 8301 : serf lan 端口，同一数据中心 consul client 通过该端口通信; 用于处理当前datacenter中LAN的gossip
- 8302 : serf wan 端口，不同数据中心 consul server 通过该端口通信; agent Server使用，处理与其他datacenter的gossip
- 8600 : dns 端口，用于已注册的服务发现

##### 启动一个server节点
`docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul agent -server -bootstrap-expect 2 -ui -bind=0.0.0.0 -client=0.0.0.0`

##### 启动第二个server节点，并加入consul1
- 查看第一个server节点的ip地址
`$docker inspect --format '{{ .NetworkSettings.IPAddress }}' consul1`
`172.17.0.2`

- 启动第二个server节点
`docker run --name consul2 -d -p 8501:8500 consul agent -server -ui -bind 0.0.0.0 -client 0.0.0.0 -join 172.17.0.2`

##### 启动第三个server节点, 并加入consul
`docker run --name consul3 -d -p 8502:8500 consul agent -server -ui -bind 0.0.0.0 -client 0.0.0.0 -join 172.17.0.2`

##### 查看consul集群成员信息

```
$docker exec -it consul1 consul members
Node          Address          Status  Type    Build  Protocol  DC   Segment
140b6da57d90  172.17.0.2:8301  alive   server  1.6.1  2         dc1  <all>
43553ae1faa6  172.17.0.4:8301  alive   server  1.6.1  2         dc1  <all>
556b7c6dcd72  172.17.0.3:8301  alive   server  1.6.1  2         dc1  <all>
```

##### 进入ui界面
通过`http://localhost:8500`可以打开ui界面；(8501或8502也可以)

#### 进度说明

#### 感谢以下的开源支持
- Go入门: [语言之旅](https://tour.go-zh.org/welcome/1)
- MySQL: [偶然翻到的一位大牛翻译的使用手册](https://chhy2009.github.io/document/mysql-reference-manual.pdf)
- Redis: [命令手册](http://redisdoc.com/)
- RabbitMQ: [英文官方](https://www.rabbitmq.com/getstarted.html) [一个中文版文档](http://rabbitmq.mr-ping.com/)
- 阿里云OSS: [文档首页](https://help.aliyun.com/product/31815.html?spm=a2c4g.750001.3.1.47287b13LQI3Ah)
- gRPC: [官方文档中文版](http://doc.oschina.net/grpc?t=56831)
- Casbin： [官网](https://casbin.org/)
- Dig： [官网](http://go.uber.org/dig)
- go-micro微服务框架: [github源码](https://github.com/micro/go-micro)
- go-micro微服务实现的在线电影院订票系统: [github源码](https://github.com/wiatingpub/MTBSystem)
我的手记博客 http://cn.blog.cn/custer