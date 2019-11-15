package casbin

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func PermissionMiddleware(enforcer *casbin.SyncedEnforcer, r *http.Request) bool {
	path := r.URL.Path
	method := r.Method
	userRecord := r.Header.Get("USER_RECORD")
	if b, err := enforcer.Enforce(userRecord, path, method); err != nil {
		logger.Error("权限验证失败", zap.Error(err))
		return false
	} else if !b {
		logger.Info("无访问权限")
		return false
	}
	return false
}
