package middleware

import (
	c "edas/service/apigw/middleware/casbin"
	"github.com/casbin/casbin/v2"
	"net/http"
	"strings"
)

var (
	e *casbin.SyncedEnforcer
)

func CheckPermission(r *http.Request) bool {
	path := r.URL.Path
	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/static") {
		return true
	}

	// TODO: JWT token验证

	// 系统菜单不进行权限拦截
	if !strings.Contains(path, "/sysMenu") {
		// casbin权限拦截
		ok := c.PermissionMiddleware(e, r)
		if !ok {
			return false
		}
	}
	return false
}

func checkURL(s string) bool {
	return true
}
