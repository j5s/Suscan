package nmap

import (
	"github.com/Ullaakut/nmap/v2"
	"log"
	"strconv"
	"strings"
)

type NmapScanRes struct {
	Url      []nmap.Hostname
	Ip       nmap.Address
	Port     string
	Protocol string
	Service nmap.Service
	State nmap.State
}

func NmapScan(ip,port string)  (nmapRes []NmapScanRes){

	var (
		resultBytes []byte
		errorBytes  []byte
	)

	s, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(port),
		nmap.WithSkipHostDiscovery(), //  -Pn
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
		panic(err)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			nmapsResTmp := NmapScanRes{host.Hostnames,host.Addresses[0],strconv.Itoa(int(port.ID)),port.Protocol,port.Service,port.State}
			nmapRes = append(nmapRes, nmapsResTmp)
		}
		//fmt.Println("nmap scan result is ",nmapRes)
	}
	return
}
