package route

import (
    "edas/service/apigw/handler"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type middleWareHandler struct {
    r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
    m := middleWareHandler{}
    m.r = r
    return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 加入中间件
    // TODO: 用于校验token的拦截器
    m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
    router := httprouter.New()
    // 用户注册
    router.POST("/api/v1/register", handler.RegisterUserHandler)
    // 用户登录
    router.POST("/api/v1/login", handler.LoginUserHandler)
    // 重置密码
    router.POST("/api/v1/reset", handler.ResetUserHandler)
    // 修改用户信息
    router.POST("/api/v1/update", handler.UpdateUserProfileHandler)
    return router
}
