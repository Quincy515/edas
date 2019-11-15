package handler

import (
	"context"
	permProto "edas/service/permission/proto"
	"edas/share/util"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

// MenuReq 请求结构体
type RoleReq struct {
	Record   string `json:"record"`
	Sequence int64  `json:"sequence"`
	Name     string `json:"name"`
	Memo     string `json:"memo"`
	Creator  string `json:"creator"`
}

func CreateRoleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &RoleReq{}
	// 将请求的body转变为结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"role","code":200,"detail":%s,"status":"OK"}`, "未正确传递参数")))
		return
	}
	// 2. 判断请求参数是否合法
	if len(ubody.Name) == 0 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"role","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.permission服务进行处理
	_, err := permCli.CreateRole(context.TODO(), &permProto.CreateRoleRequest{
		Name:     ubody.Name,
		Sequence: ubody.Sequence,
		Memo:     ubody.Memo,
		Creator:  ubody.Creator,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("创建角色成功"))
}

func DeleteRoleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &RoleReq{}
	// 将请求的body转变为结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "未正确传递参数")))
		return
	}
	// 2. 判断请求参数是否合法
	if len(ubody.Record) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.permission服务进行处理
	_, err := permCli.DeleteRole(context.TODO(), &permProto.DeleteRoleRequest{
		Record: ubody.Record,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("删除角色成功"))
}

func UpdateRoleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &RoleReq{}
	// 将请求的body转变为结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "未正确传递参数")))
		return
	}
	// 2. 判断请求参数是否合法
	if len(ubody.Record) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.permission服务进行处理
	_, err := permCli.UpdateRole(context.TODO(), &permProto.UpdateRoleRequest{
		Record:   ubody.Record,
		Name:     ubody.Name,
		Sequence: ubody.Sequence,
		Memo:     ubody.Memo,
		Creator:  ubody.Creator,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("更新角色成功"))
}

func QueryRoleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func GetRoleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &RoleReq{}
	// 将请求的body转变为结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "未正确传递参数")))
		return
	}
	// 2. 判断请求参数是否合法
	if len(ubody.Record) < 6 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.permission服务进行处理
	resp, err := permCli.GetRole(context.TODO(), &permProto.GetRoleRequest{
		Record: ubody.Record,
		Name:   ubody.Name,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	// 4. 构建并响应用户数据
	rsp := util.RespMsg{
		Id:   "role",
		Code: 200,
		Detail: RoleReq{
			Record:   resp.Record,
			Name:     resp.Name,
			Sequence: resp.Sequence,
			Memo:     resp.Memo,
			Creator:  resp.Creator,
		},
		Status: "OK",
	}
	_, _ = w.Write(rsp.JSONBytes())
}
