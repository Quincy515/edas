package handler

import (
	"context"
	userProto "edas/service/user/proto"
	"edas/share/config"
	"edas/share/log"
	"edas/share/util"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	//ratelimit2 "github.com/juju/ratelimit"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro"
	//"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var (
	userCli userProto.UserService
	logger  *zap.Logger
)

func init() {
	logger = log.Init("api")
	// 配置请求容量及QPS
	//bRate := ratelimit2.NewBucketWithRate(100, 1000)
	service := micro.NewService(
	//加入限流功能, false为不等待(超限即返回请求失败)
	//micro.WrapClient(ratelimit.NewClientWrapper(bRate, false)),
	// 加入熔断功能, 处理rpc调用失败的情况(cirucuit breaker)
	//micro.WrapClient(hystrix.NewClientWrapper()),
	)
	// 初始化，解析命令行参数等
	service.Init()
	// 初始化服务的客户端
	cli := service.Client()

	// 初始化一个user服务的客户端
	userCli = userProto.NewUserService(config.NameSpace+config.ServiceNameUser, cli)
}

// 请求request的结构体
type UserReq struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
	Record      string `json:"record"`
	UserEmail   string `json:"user_email"`
	UserPhone   string `json:"user_phone"`
}

// RegisterUserHandler 处理用户注册请求
func RegisterUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body) // 获取请求的body
	ubody := &UserReq{}              // 构造接收请求的结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "未传递参数")))
		return // 将请求的body转变为结构体
	}
	// 2. 判断请求参数中是否合法
	if len(ubody.UserName) < 1 || len(ubody.Password) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.user服务进行处理
	_, err := userCli.RegisterUser(context.TODO(), &userProto.RegisterUserRequest{
		UserName: ubody.UserName,
		Password: ubody.Password,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("注册成功"))
}

// LoginUserHandler 处理用户登录请求
func LoginUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body) // 获取请求的body
	ubody := &UserReq{}              // 构造接收请求的结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "未传递参数")))
		return // 将请求的body转变为结构体
	}
	// 2. 判断请求参数中是否合法
	if len(ubody.UserName) < 1 || len(ubody.Password) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.user服务进行处理
	resp, err := userCli.LoginUser(context.TODO(), &userProto.LoginUserRequest{
		UserName: ubody.UserName,
		Password: ubody.Password,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	// 4. 构建并响应用户数据
	rsp := util.RespMsg{
		Id:   "user",
		Code: 200,
		Detail: struct {
			UserName string
			Record   string
			Email    string
			Phone    string
		}{
			UserName: resp.UserName,
			Record:   resp.Record,
			Email:    resp.Email,
			Phone:    resp.Phone,
		},
		Status: "OK",
	}
	_, _ = w.Write(rsp.JSONBytes())
}

// ResetUserHandler 处理用户重置密码操作
func ResetUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body) // 获取请求的body
	ubody := &UserReq{}              // 构造接收请求的结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "未传递参数")))
		return // 将请求的body转变为结构体
	}
	// 2. 判断请求参数中是否合法
	if len(ubody.UserName) < 1 || len(ubody.Password) < 6 ||
		len(ubody.NewPassword) < 6 || len(ubody.Record) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.user服务进行处理
	_, err := userCli.ResetUser(context.TODO(), &userProto.ResetUserRequest{
		UserName:    ubody.UserName,
		Password:    ubody.Password,
		NewPassword: ubody.NewPassword,
		Record:      ubody.Record,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("reset user success"))
}

// UpdateUserProfileHandler 修改用户信息
func UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body) // 获取请求的body
	ubody := &UserReq{}              // 构造接收请求的结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "未传递参数")))
		return // 将请求的body转变为结构体
	}
	// 2. 判断请求参数中是否合法
	if len(ubody.Record) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"user","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.user服务进行处理
	_, err := userCli.UpdateUserProfile(context.TODO(), &userProto.UpdateUserProfileRequest{
		UserName:  ubody.UserName,
		UserEmail: ubody.UserEmail,
		UserPhone: ubody.UserPhone,
		Record:    ubody.Record,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("update user success"))
}
