package main

import (
	"edas/service/permission/db"
	"edas/service/permission/handler"
	proto "edas/service/permission/proto"

	"edas/share/config"
	"edas/share/log"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"go.uber.org/zap"
	"time"
)

func main() {
	log.Init("permission")
	logger := log.Instance()
	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name(config.NameSpace+config.ServiceNamePermission),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	// 初始化Service，解析命令行参数
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("permission-srv", "permission-srv is start..."))
			// 初始化db链接
			db.Init(config.MySQLSource)
			// 给微服务绑定handler
			_ = proto.RegisterPermissionServiceHandler(service.Server(), new(handler.PermissionServiceHandler))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("permission-srv", "permission-srv is stop..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	// 启动Service，通过Run开启
	if err := service.Run(); err != nil {
		logger.Panic("permission-srv服务启动失败...")
	}
}
