package PortCheck

import (
	"fmt"
	"net"
	"sync"
	"tcpScanner/GetIPInfo"
	"time"
)

func CheckPortStatus() (success []string, failed []string) {
	var wg = sync.WaitGroup{}
	list, _, err := GetIPInfo.GetIPList("ip.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	for _, ip := range list {
		wg.Add(1)
		go func(ip string) {
			_, err := net.DialTimeout("tcp", ip, 4*time.Second)
			time.Sleep(time.Second * 4)
			if err != nil {
				fmt.Printf("%40v    close\n", ip)
				failed = append(failed, ip)
			} else {
				fmt.Printf("%40v    open\n", ip)
				success = append(success, ip)
			}
			wg.Done()
		}(ip)
	}
	wg.Wait()
	return success, failed
}
