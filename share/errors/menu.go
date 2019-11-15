package errors

import (
	"edas/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeMenuFailed = 401
)

var (
	ErrorMenuFailed = errors.New(
		config.ServiceNamePermission, "操作异常", errorCodeMenuFailed)

	ErrorMenuFailedParams = errors.New(
		config.ServiceNamePermission, "参数异常", errorCodeMenuFailed)

	ErrorMenuForbidden = errors.New(
		config.ServiceNamePermission, "没有权限", errorCodeMenuFailed)

	ErrorMenuAlreadyExists = errors.New(
		config.ServiceNamePermission, "该菜单已存在", errorCodeMenuFailed)

	ErrorNotFoundMenuItem = errors.New(
		config.ServiceNamePermission, "该菜单不存在", errorCodeMenuFailed)

	ErrorNotAllowDeleteWithChild = errors.New(
		config.ServiceNamePermission, "含有子级，不能删除", errorCodeMenuFailed)

	ErrorUpdateMenuFailed = errors.New(
		config.ServiceNamePermission, "更新菜单失败", errorCodeMenuFailed)
)
