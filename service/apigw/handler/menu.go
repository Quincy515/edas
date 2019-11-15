package handler

import (
	"context"
	permProto "edas/service/permission/proto"
	"edas/share/util"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

// MenuReq 请求结构体
type MenuReq struct {
	Record     string `json:"record"`
	Name       string `json:"name"`
	Sequence   int64  `json:"sequence"`
	Icon       string `json:"icon"`
	Router     string `json:"router"`
	Hidden     int64  `json:"hidden"`
	ParentId   string `json:"parent_id"`
	ParentPath string `json:"parent_path"`
	Creator    string `json:"creator"`
}

func QueryMenuHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetMenuHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &MenuReq{}
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
	resp, err := permCli.GetMenu(context.TODO(), &permProto.GetMenuRequest{
		Record: ubody.Record,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	// 4. 构建并响应用户数据
	rsp := util.RespMsg{
		Id:   "menu",
		Code: 200,
		Detail: MenuReq{
			Record:     resp.Record,
			Name:       resp.Name,
			Sequence:   resp.Sequence,
			Icon:       resp.Icon,
			Router:     resp.Router,
			Hidden:     resp.Hidden,
			ParentId:   resp.ParentId,
			ParentPath: resp.ParentPath,
			Creator:    resp.Creator,
		},
		Status: "OK",
	}
	_, _ = w.Write(rsp.JSONBytes())
}

func CreateMenuHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &MenuReq{}
	// 将请求的body转变为结构体
	if err := jsoniter.Unmarshal(res, ubody); err != nil {
		log.Println(err)
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "未正确传递参数")))
		return
	}
	// 2. 判断请求参数是否合法
	if len(ubody.Name) < 1 {
		_, _ = w.Write([]byte(fmt.Sprintf(
			`{"id":"menu","code":200,"detail":%s,"status":"OK"}`, "请求参数非法")))
		return
	}
	// 3. 将参数传递给yak.permission服务进行处理
	_, err := permCli.CreateMenu(context.TODO(), &permProto.CreateMenuRequest{
		Name:       ubody.Name,
		Sequence:   ubody.Sequence,
		Icon:       ubody.Icon,
		Router:     ubody.Router,
		Hidden:     ubody.Hidden,
		ParentId:   ubody.ParentId,
		ParentPath: ubody.ParentPath,
		Creator:    ubody.Creator,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("创建菜单成功"))
}

func UpdateMenuHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &MenuReq{}
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
	_, err := permCli.UpdateMenu(context.TODO(), &permProto.UpdateMenuRequest{
		Record:     ubody.Record,
		Name:       ubody.Name,
		Sequence:   ubody.Sequence,
		Icon:       ubody.Icon,
		Router:     ubody.Router,
		Hidden:     ubody.Hidden,
		ParentId:   ubody.ParentId,
		ParentPath: ubody.ParentPath,
		Creator:    ubody.Creator,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("更新菜单成功"))
}

func DeleteMenuHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1. 请求转变为结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &MenuReq{}
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
	_, err := permCli.DeleteMenu(context.TODO(), &permProto.DeleteMenuRequest{
		Record: ubody.Record,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}
	_, _ = w.Write([]byte("删除菜单成功"))
}
