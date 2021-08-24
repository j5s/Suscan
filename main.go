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
	//err = s.ReadSection("Masscan", &global.MasscanSetting)
	//if err != nil {
	//	return err
	//}
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	models.Setup()
	return nil
}

func main() {
	endless.DefaultReadTimeOut = global.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = global.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", global.ServerSetting.HttpPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
