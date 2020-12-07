package main

import (
	"fmt"
	"go-api/api/bootstrap"
	"go-api/api/config"
	conf "go-api/api/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}
func main() {
	app := bootstrap.Start()
	addr := fmt.Sprintf(":%s", conf.GetString("app.port"))
	app.Run(addr)
}
