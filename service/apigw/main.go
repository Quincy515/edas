package main

import (
	"edas/service/apigw/route"
	"edas/share/log"
	"go.uber.org/zap"
	"net/http"
)

var (
	cors   = map[string]bool{"*": true}
	logger *zap.Logger
)

func init() {
	logger = log.Init("api")
}

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handleRPC)
	r := route.RegisterHandlers()
	mh := route.NewMiddleWareHandler(r)
	err := http.ListenAndServe(":7001", mh)
	if err != nil {
		logger.Error("start api gateway server failed")
	}
	logger.Info("Listen on :7001")
}
