package util

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

// RespMsg http响应数据的通用结构
type RespMsg struct {
	Id     string      `json:"id"`
	Code   int         `json:"code"`
	Detail interface{} `json:"detail"`
	Status string      `json:"status"`
}

// NewRespMsg 生成response对象
func NewRespMsg(id string, code int, detail interface{}, status string) *RespMsg {
	return &RespMsg{
		Id:     id,
		Code:   code,
		Detail: detail,
		Status: status,
	}
}

// JSONBytes 对象转json格式的二进制数组
func (rsp *RespMsg) JSONBytes() []byte {
	r, err := jsoniter.Marshal(rsp)
	if err != nil {
		logger.Error("对象转json格式的二进制数组失败")
	}
	return r
}

func (rsp *RespMsg) JSONString() string {
	r, err := jsoniter.Marshal(rsp)
	if err != nil {
		logger.Error("对象转json格式的字符串失败")
	}
	return string(r)
}
