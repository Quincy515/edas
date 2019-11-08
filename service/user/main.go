package main

import (
    "edas/service/user/db"
    "edas/service/user/handler"
    proto "edas/service/user/proto"
    "edas/share/config"
    "edas/share/log"
    "github.com/micro/cli"
    "github.com/micro/go-micro"
    "go.uber.org/zap"
    "time"
)

func main() {
    log.Init("user")
    logger := log.Instance()
    // 创建Service，并定义一些参数
    service := micro.NewService(
        micro.Name(config.NameSpace+config.ServiceNameUser),
        micro.RegisterTTL(time.Second*10),
        micro.RegisterInterval(time.Second*5),
    )
    // 初始化service,解析命令行参数
    service.Init(
        micro.Action(func(c *cli.Context) {
            logger.Info("Info", zap.Any("user-srv", "user-srv is start ..."))

            // 初始化db 链接
            db.Init(config.MySQLSource)
            // 给微服务绑定handler
            _ = proto.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
        }),
        micro.AfterStop(func() error {
            logger.Info("Info", zap.Any("user-srv", "user-srv is stop..."))
            return nil
        }),
        micro.AfterStart(func() error {
            return nil
        }),
    )

    // 启动service，通过Run开启
    if err := service.Run(); err != nil {
        logger.Panic("user-sev服务启动失败...")
    }
}
