package handler

import (
	"context"
	"edas/service/user/db"
	proto "edas/service/user/proto"
	"edas/share/errors"
	"edas/share/log"
	"edas/share/util"
	"go.uber.org/zap"
)

// UserService 用于实现UserServiceHandler接口的对象
type UserService struct {
	logger *zap.Logger
}

type test struct {
	code    string
	message string
}

func NewUser() *UserService {
	return &UserService{
		logger: log.Instance(),
	}
}

// RegisterUser 账号注册
func (u *UserService) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest, resp *proto.RegisterUserResponse) error {
	// 1. 获取请求参数
	username := req.UserName
	password := req.Password
	// 2. 判断该用户名是否已经存在
	user, err := db.SelectUserByUserName(username)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user != nil {
		return errors.ErrorUserAlreadyExists
	}
	// 3. 未查询到记录就进行注册操作
	// 对密码hash之后保存到数据库中
	passwd := util.SHA1HashString(password)
	record := util.MustUUID()
	err = db.InsertUser(username, passwd, record)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed // 注册用户失败
	}
	return nil
}

// LoginUser 账号登录
func (u *UserService) LoginUser(ctx context.Context, req *proto.LoginUserRequest, rsp *proto.LoginUserResponse) error {
	username := req.UserName
	password := req.Password
	passwd := util.SHA1HashString(password)
	user, err := db.SelectUserByUsernamePassword(username, passwd)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user == nil {
		return errors.ErrorUserLoginFailed
	}
	rsp.UserName = user.UserName
	rsp.Record = user.RecordId
	rsp.Email = user.Email
	rsp.Phone = user.Phone
	return nil
}

// ResetUser 重置密码
func (u *UserService) ResetUser(ctx context.Context, req *proto.ResetUserRequest, rsp *proto.ResetUserResponse) error {
	// 1. 获取请求参数
	username := req.UserName
	password := req.Password
	newPassword := req.NewPassword
	record := req.Record
	// 2. 查看原账号密码是否正确
	passwd := util.SHA1HashString(password)
	user, err := db.SelectUserByUsernamePassword(username, passwd)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user == nil {
		return errors.ErrorUserLoginFailed
	}
	// 3. 修改密码
	newPasswd := util.SHA1HashString(newPassword)
	err = db.UpdateUserPassword(username, newPasswd, record)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	return nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, req *proto.UpdateUserProfileRequest, rsp *proto.UpdateUserProfileResponse) error {
	userEmail := req.UserEmail
	userName := req.UserName
	userPhone := req.UserPhone
	record := req.Record
	if userEmail != "" {
		// 判断该邮箱是否已经被绑定
		user, err := db.SelectUserByUserEmail(userEmail)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
		if user != nil {
			return errors.ErrorEmailAlreadyExists
		}
		// 若该邮箱不存在则可以进行更新
		err = db.UpdateUserEmailProfile(userEmail, record)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userName != "" {
		// 判断该用户名是否已经存在
		user, err := db.SelectUserByUserName(userName)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
		if user != nil {
			return errors.ErrorUserAlreadyExists
		}
		// 若该用户名不存在则可以进行更新
		err = db.UpdateUserNameProfile(userName, record)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userPhone != "" {
		// 判断该手机号是否已经存在
		user, err := db.SelectUserByUserPhone(userPhone)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
		if user != nil {
			return errors.ErrorPhoneAlreadyExists
		}
		// 若该手机号不存在则可以进行更新
		err = db.UpdateUserPhoneProfile(userPhone, record)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	return nil
}
