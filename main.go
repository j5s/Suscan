package main

import (
	"Suscan/global"
	"Suscan/models"
	"Suscan/pkg/setting"
	"Suscan/routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"strings"
	"syscall"
	"time"
)


func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("configs read fail: %v", err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split("configs/", ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Masscan", &global.MasscanSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	models.Setup()
	return nil
}

func main() {
	//实现 Golang HTTP/HTTPS 服务重新启动的零停机
	endless.DefaultReadTimeOut = global.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = global.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", global.ServerSetting.HttpPort)
	//返回一个初始化的 endlessServer 对象
	server := endless.NewServer(endPoint, routers.InitRouter())
	//在 BeforeBegin 时输出当前进程的 pid
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	//调用 ListenAndServe 将实际“启动”服务
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
