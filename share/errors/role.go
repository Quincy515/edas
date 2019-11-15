package errors

import (
	"edas/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeRoleFailed = 401
)

var (
	ErrorRoleFailed = errors.New(
		config.ServiceNamePermission, "操作异常", errorCodeRoleFailed)

	ErrorRoleAlreadyExists = errors.New(
		config.ServiceNamePermission, "角色已存在", errorCodeRoleFailed)

	ErrorCreateRoleFailed = errors.New(
		config.ServiceNamePermission, "创建角色失败", errorCodeRoleFailed)

	ErrorNotFoundRole = errors.New(
		config.ServiceNamePermission, "角色不存在", errorCodeRoleFailed)

	ErrorDeleteRoleFailed = errors.New(
		config.ServiceNamePermission, "删除角色失败", errorCodeRoleFailed)

	ErrorUpdateRoleFailed = errors.New(
		config.ServiceNamePermission, "更新角色失败", errorCodeRoleFailed)
)
