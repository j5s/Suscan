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
	go InitNmapscan()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//实现端口扫描
func InitNmapscan() {
	start := time.Now()
	//数据库读取要扫描的资产
	result := models.GetAllAsset()
	for _,r := range result{
		ipstmp := r.Target
		ips = append(ips, ipstmp)
	}
	//数据库读取要扫描的端口
	port := models.GetSettingPort("port")
	//数据库读取nmap -T 的值
	t := models.GetSettingTiming("timetemplate")
	t1, _ := strconv.Atoi(t)
	NmapStart(ips,port,t1)
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

func NmapStart(ips []string, port string,t int)  {
	//数据库获取nmap个数的配置
	cmd := models.GetSettingCmd("cmd")
	//转为int
	cmdInt, _ := strconv.Atoi(cmd)

	//开启多个nmap协程
	ch := make(chan int,cmdInt)
	for _, ip := range ips {
		ch <- 1
		go NmapScan(ip,port,t,ch)
	}
}

func NmapScan(ip,port string,t int,ch chan int)  {

	nmapRes := nmap.NmapScan(ip,port,t)
	//fmt.Println(nmapRes)
	// 并发处理扫描结果
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
	// 生产完成后，从生产方关闭task
	close(taskChan)
	wg.Wait()
	<- ch
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
		data["state"] = fmt.Sprintf("%s",target.State)
		data["protocol"] = target.Protocol
		data["service"] = fmt.Sprintf("%s",target.Service)
		data["res_code"] = target.Res_code
		data["res_result"] = target.Res_result
		data["res_type"] = target.Res_type
		data["res_url"] = target.Res_url
		data["res_title"] = target.Res_title

		//扫描结果入库前对比
		ok, id := models.ExistIplist(target.Ip, target.Port)
		if ok {
			//fmt.Println(target.Ip, target.Port, "更新")
			nowTime := time.Now().Format("20060102150405")
			dataUpdate["updated_time"] = nowTime
			models.EditIplist(id, dataUpdate)
			wg.Done()
		} else {
			if target.State == "open" {
				models.AddIplist(data)
				wg.Done()
			fmt.Println(target.Ip, target.Port, "插入")
			}else {
				wg.Done()
			}
		}
	}
}