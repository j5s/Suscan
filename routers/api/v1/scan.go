package v1

import (
	"Suscan/models"
	"Suscan/pkg/e"
	"Suscan/pkg/nmap"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var ips []string

// 端口扫描api

func ScanPort(c *gin.Context) {
	code := e.SUCCESS
	go InitNmapScan()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//实现端口扫描

func InitNmapScan() {
	start := time.Now()
	//数据库读取要扫描的资产
	result := models.GetAllAsset()
	for _,r := range result{
		ipstmp := r.Target
		ips = append(ips, ipstmp)
	}
	//数据库读取要扫描的端口
	port := models.GetSettingPort("port")

	NmapStart(ips,port)
	costTime := time.Since(start)
	data := make(map[string]interface{})
	data["task_name"] = "PortScan"
	data["task_type"] = "nmapscan"
	data["all_num"] = len(ips)
	data["run_time"] = fmt.Sprintf("%s", costTime)
	models.AddLog(data)
}

type NmapScanRes struct {
	Ip       string
	Port     string
	Protocol string
}

func NmapStart(ips []string, port string)  {
	//配置nmap个数
	cmd := models.GetSettingCmd("cmd")
	//转为int
	cmdInt, _ := strconv.Atoi(cmd)

	//开启多个nmap协程
	ch := make(chan int,cmdInt)

	for _, ip := range ips {
		ch <- 1
		go NmapScan(ip,port,ch)
	}
}

func NmapScan(ip,port string,ch chan int)  {
	noping := models.GetSettingNoPing("noping")
	if noping == "0" {
		nmapRes := nmap.NmapScan(ip,port)
		wg := &sync.WaitGroup{}
		// 创建一个buffer为thread * 2的channel
		thread := 2
		taskChan := make(chan nmap.NmapScanRes, 50*2)
		// 创建Thread个协程
		for i := 0; i < thread; i++ {
			go ScanResult(taskChan, wg)
		}
		for _, task := range nmapRes {
			wg.Add(1)
			taskChan <- task
		}
		close(taskChan)
		wg.Wait()
		<- ch
	}else {
		nmapRes := nmap.NmapScanPn(ip,port)
		wg := &sync.WaitGroup{}
		// 创建一个buffer为thread * 2的channel
		thread := 2
		taskChan := make(chan nmap.NmapScanRes, 50*2)
		// 创建Thread个协程
		for i := 0; i < thread; i++ {
			go ScanResult(taskChan, wg)
		}
		for _, task := range nmapRes {
			wg.Add(1)
			taskChan <- task
		}
		close(taskChan)
		wg.Wait()
		<- ch
	}
}

func ScanResult(taskChan chan nmap.NmapScanRes, wg *sync.WaitGroup) {
	data := make(map[string]interface{})
	dataUpdate := make(map[string]interface{})
	for target := range taskChan {
		defer func() {
			if err := recover(); err != nil {
				wg.Done()
			}
		}()

		data["url"] = target.Url
		data["ip"] = target.Ip
		data["port"] = target.Port
		data["state"] = target.State
		data["protocol"] = target.Protocol
		data["service"] = target.Service
		data["res_code"] = target.ResCode
		data["res_result"] = target.ResResult
		data["res_type"] = target.ResType
		data["res_url"] = target.ResUrl
		data["res_title"] = target.ResTitle

		//扫描结果入库前对比
		ok, id := models.ExistIplist(target.Ip, target.Port)
		if ok {
			nowTime := time.Now().Format("20060102150405")
			dataUpdate["updated_time"] = nowTime
			models.EditIplist(id, dataUpdate)
			wg.Done()
		} else {
			if target.State == "open" {
				models.AddIplist(data)
				wg.Done()
			}else {
				wg.Done()
			}
		}
	}
}