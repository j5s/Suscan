package nmap

import (
	"Suscan/pkg/utils"
	"fmt"
	"github.com/Ullaakut/nmap/v2"
	"log"
	"strconv"
	"strings"
)

type NmapScanRes struct {
	Url       string
	Ip        string
	Port      string
	Protocol  string
	Service   string
	State     string
	ResCode   string
	ResResult string
	ResType   string
	ResUrl    string
	ResTitle  string
}

type IdentifyResult struct {
	Type     string
	RespCode string
	Result   string
	ResultNc string
	Url      string
	Title    string
}

func NmapScan(ip, port string) (nmapRes []NmapScanRes) {

	var (
		resultBytes []byte
		errorBytes  []byte
	)

	s, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(port),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	if err := s.RunAsync(); err != nil {
		panic(err)
	}

	stdout := s.GetStdout()

	stderr := s.GetStderr()

	go func() {
		for stdout.Scan() {
			resultBytes = append(resultBytes, stdout.Bytes()...)
		}
	}()

	go func() {
		for stderr.Scan() {
			errorBytes = append(errorBytes, stderr.Bytes()...)
		}
	}()

	if err := s.Wait(); err != nil {
		panic(err)
	}

	result, err := nmap.Parse(resultBytes)

	result.NmapErrors = strings.Split(string(errorBytes), "\n")

	if err != nil {
		//panic(err)
		return
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		for _, port := range host.Ports {
			//格式化nmap扫描返回的字段为string
			hostname := fmt.Sprintf("%s", host.Hostnames)
			context := strings.Fields(hostname)
			hostname = context[0]
			hostname = strings.Trim(hostname, "[")
			hostname = strings.Trim(hostname, "]")
			ip := fmt.Sprintf("%s", host.Addresses[0])
			portchange := strconv.Itoa(int(port.ID))
			service := fmt.Sprintf("%s", port.Service)
			state := fmt.Sprintf("%s", port.State)
			if state == "open" {
				if hostname != "" {
					url := utils.ParseUrl(hostname, portchange)
					for _, results := range utils.Identify(url, 5) {
						//fmt.Println("识别url的results为： ")
						//fmt.Println(results)
						nmapsResTmp := NmapScanRes{hostname, ip, portchange, port.Protocol, service, state, results.RespCode, results.Result, results.Type, results.Url, results.Title}
						nmapRes = append(nmapRes, nmapsResTmp)
					}
				} else {
					url := utils.ParseUrl(ip, portchange)
					for _, results := range utils.Identify(url, 5) {
						//fmt.Println("识别ip的results为： ")
						//fmt.Println(results)
						nmapsResTmp := NmapScanRes{hostname, ip, portchange, port.Protocol, service, state, results.RespCode, results.Result, results.Type, results.Url, results.Title}
						nmapRes = append(nmapRes, nmapsResTmp)
					}
				}
				return
			}
		}
	}
	return 
}