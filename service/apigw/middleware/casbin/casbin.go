package casbin

//import (
//    "net/http"
//    "strings"
//)

//func CheckPermission(r *http.Request) bool {
//    path := r.URL.Path
//    // 过滤静态资源、login接口、首页等...不需要验证
//    if checkURL(path) || strings.Contains(path, "/static") {
//        return true
//    }
//
//    // TODO: JWT token验证
//
//}