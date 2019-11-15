# EDAS

[Api在线文档](https://documenter.getpostman.com/view/4827715/SW7W5Vdo)

#### 介绍
企业级分布式应用服务EDAS（Enterprise Distributed Application Service）
期望做一个基于go-micro + casbin + jwt 的用户认证和权限的微服务 [参考](https://github.com/winyh/accbase)

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
go get github.com/casbin/casbin
go get github.com/casbin/casbin-server
```

2.  项目结构概览

```
├─.gitee
├─service  服务划分
│  ├─apigw 网关
│  │  ├─handler     
│  │  ├─middleware
│  │  │  └─casbin
│  │  └─route
│  ├─permission 权限微服务
│  │  ├─db
│  │  ├─handler
│  │  └─proto
│  └─user       用户微服务
│      ├─db
│      ├─handler
│      └─proto
├─share
│  ├─config  项目配置
│  ├─errors  错误说明
│  ├─log     zap日志配置
│  └─util    工具集合
└─sql 数据库字段
```

#### 进度说明
- [*] 账号系统，注册/登录/查询用户数据
- [ ] JWT认证 未完成
- [*] 菜单/角色/用户的增删改查操作
- [ ] 菜单动作/菜单资源/角色与菜单的绑定/用户与角色绑定 未完成
- [ ] casbin 的 RBAC 访问控制 未完成

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
- 一个Go Web Api 后端 简单例子,包含 用户、权限、菜单、JWT 、 RBAC(Casbin)等: [github源码](https://github.com/hequan2017/go-admin)
- Example of role-based HTTP Authorization with casbin: [github源码](https://github.com/zupzup/casbin-http-role-example)
- 基于casbin实现的身份认证及验证授权服务例子: [github源码](https://github.com/Soontao/go-simple-api-gateway)
- 基于Go Micro + Casbin + Gin + Gorm + Casbin + JWT 用户认证和权限微服务: [正在开发](https://github.com/winyh/accbase)
我的手记博客 http://cn.blog.cn/custer

#### 共同学习
> 该项目是源于自己的项目积累及个人思考，因为个人能力有限，希望有小伙伴可以一起参与共同学习进步
1. Fork 本仓库
2. 新建分支
3. 提交代码
4. 新建 Pull Request

> 期待与您一起学习进步，下面是我的微信二维码：

![微信二维码](./wechat.png)
