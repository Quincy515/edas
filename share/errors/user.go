package errors

import (
	"edas/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeUserSuccess = 200
)

var (
	ErrorUserSuccess = errors.New(
		config.ServiceNameUser, "操作成功", errorCodeUserSuccess)

	ErrorUserFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeUserSuccess)

	ErrorUserAlreadyExists = errors.New(
		config.ServiceNameUser, "该用户名已经被注册", errorCodeUserSuccess)

	ErrorEmailAlreadyExists = errors.New(
		config.ServiceNameUser, "该邮箱已经被绑定", errorCodeUserSuccess)

	ErrorPhoneAlreadyExists = errors.New(
		config.ServiceNameUser, "该手机号已经被绑定", errorCodeUserSuccess)

	ErrorUserLoginFailed = errors.New(
		config.ServiceNameUser, "密码或者用户名错误", errorCodeUserSuccess)
)
