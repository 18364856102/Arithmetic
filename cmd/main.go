package main

import (
	"ArithmeticOperation/router"
	"ArithmeticOperation/service"
	"net/http"

	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"os"
)

var logger log.Logger

const (
	logFile = "main.log"
)

func main() {
	path := logFile + ".txt"
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("open file path: %s err %s", path, err.Error()))
	}

	logger = log.NewLogfmtLogger(logFile)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// ===== 创建路由 =====
	httpHandler := mux.NewRouter()
	svc := service.New(logger)
	router.HandlerArithmeticOperation(httpHandler, svc)

	// 开启监听
	err = http.ListenAndServe(":8081", httpHandler)
	if err != nil {
		_ = logger.Log("[main]", "Listen Error", err)
	}

	//var g run.Group
	//{
	//	httpListener, err := net.Listen("tcp", ":8081")
	//	if err != nil {
	//		_ = logger.Log("[main]", "transport HTTP", "Listen Error", err)
	//		os.Exit(1)
	//	}
	//	g.Add(func() error {
	//		_ = logger.Log("[main]", "transport", "HTTPAddr", ":8081")
	//		return http.Serve(httpListener, httpHandler) // http
	//
	//	}, func(error) {
	//		_ = httpListener.Close()
	//	})
	//}
}
