package route

import (
	"edas/service/apigw/handler"
	"edas/service/apigw/middleware"
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
	middleware.CheckPermission(r)
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

	// 菜单查询
	router.GET("/api/v1/menus/query", handler.QueryMenuHandler)
	// 查询指定菜单详情
	router.GET("/api/v1/menus/get", handler.GetMenuHandler)
	// 新增菜单
	router.POST("/api/v1/menus/create", handler.CreateMenuHandler)
	// 更新菜单
	router.PUT("/api/v1/menus/update", handler.UpdateMenuHandler)
	// 删除菜单
	router.DELETE("/api/v1/menus/delete", handler.DeleteMenuHandler)

	// 角色查询
	router.GET("/api/v1/role/query", handler.QueryRoleHandler)
	// 查询指定角色
	router.GET("/api/v1/role/get", handler.GetRoleHandler)
	// 新增角色
	router.POST("/api/v1/role/create", handler.CreateRoleHandler)
	// 更新角色
	router.PUT("/api/v1/role/update", handler.UpdateRoleHandler)
	// 删除角色
	router.DELETE("/api/v1/role/delete", handler.DeleteRoleHandler)
	return router
}
