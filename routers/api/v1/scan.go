package v1

import (
	"Suscan/models"
	"Suscan/pkg/e"
	"Suscan/pkg/nmap"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// 开始nmap端口扫描
func ScanPort(c *gin.Context) {

	code := e.SUCCESS
	go InitNmapscan()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}



func InitNmapscan() {

	start := time.Now()
	ips :=[]string{"baidu.com","qq.com","tuya.com","taobao.com","sohu.com"}
	NmapStart(ips)
	costTime := time.Since(start)
	data := make(map[string]interface{})
	data["task_name"] = "PortScan"
	data["task_type"] = "nmapscan"
	data["all_num"] = len(ips)
	data["run_time"] = fmt.Sprintf("%s", costTime)
	models.AddLog(data)

}

func NmapStart(ips []string)  {
	//开启多个nmap协程
	ch := make(chan int,1)
	for _, ip := range ips {
		ch <- 1
		go NmapScan(ip,"1-65535",ch)
	}
}

func NmapScan(ip,port string,ch chan int)  {

	nmapRes := nmap.NmapScan(ip,port)
	fmt.Println(nmapRes)
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

		data["url"] = fmt.Sprintf("%s",target.Url)
		data["ip"] = target.Ip.String()
		data["port"] = target.Port
		data["state"] = fmt.Sprintf("%s",target.State)
		data["protocol"] = target.Protocol
		data["service"] = fmt.Sprintf("%s",target.Service)

		//扫描结果入库前对比
		ok, id := models.ExistIplist(target.Ip.String(), target.Port)
		if ok {
			//fmt.Println(target.Ip, target.Port, "更新")
			nowTime := time.Now().Format("20060102150405")
			dataUpdate["updated_time"] = nowTime
			models.EditIplist(id, dataUpdate)
			wg.Done()
		} else {
			if target.State.State == "open" {
				models.AddIplist(data)
				wg.Done()
				//fmt.Println(target.Ip, target.Port, "插入")
			}else {
				wg.Done()
			}
		}
	}
}