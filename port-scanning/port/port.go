package port

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port    string
	State   string
	Service string
}

var mutex sync.Mutex
var results []ScanResult

/*
function which will iterate through ports on a host and try and to open a UDP/TCP connection on those ports
*/

/**
#################### Uncomment below to run with goroutine #########################################
*/

func InitialScan(hostname string) []ScanResult {
	//var results []ScanResult
	var wg sync.WaitGroup

	// scan from 0 to 1024
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go ScanPort("tcp", hostname, i, &wg)
	}
	wg.Wait()

	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go ScanPort("udp", hostname, i, &wg)
	}
	wg.Wait()

	return results
}

func WideScan(hostname string) []ScanResult {
	//var results []ScanResult
	var wg sync.WaitGroup

	for i := 0; i <= 49152; i++ {
		wg.Add(1)
		go ScanPort("udp", hostname, i, &wg)
	}
	wg.Wait()

	for i := 0; i <= 49152; i++ {
		wg.Add(1)
		go ScanPort("tcp", hostname, i, &wg)
	}
	wg.Wait()

	return results
}

func ScanPort(protocol, hostname string, port int, wg *sync.WaitGroup) {
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
	}
	if conn != nil {
		defer conn.Close()
	}
	result.State = "Open"

	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()

	wg.Done()
}

/**
#################### Uncomment below to run in serial #########################################
*/

//func InitialScan(hostname string) []ScanResult {
//
//	var results []ScanResult
//
//	for i := 0; i <= 1024; i++ {
//		results = append(results, ScanPort("udp", hostname, i))
//	}
//
//	for i := 0; i <= 1024; i++ {
//		results = append(results, ScanPort("tcp", hostname, i))
//	}
//
//	return results
//}
//
//func ScanPort(protocol, hostname string, port int) ScanResult {
//	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
//	address := hostname + ":" + strconv.Itoa(port)
//	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
//
//	if err != nil {
//		result.State = "Closed"
//		return result
//	}
//	defer conn.Close()
//	result.State = "Open"
//	return result
//}
//
//func WideScan(hostname string) []ScanResult {
//	var results []ScanResult
//
//	for i := 0; i <= 49152; i++ {
//		results = append(results, ScanPort("udp", hostname, i))
//	}
//
//	for i := 0; i <= 49152; i++ {
//		results = append(results, ScanPort("tcp", hostname, i))
//	}
//
//	return results
//}
