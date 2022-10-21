package main

import (
	"fmt"
	"tcpScanner/PortCheck"
)

func main() {
	s, f := PortCheck.CheckPortStatus()
	fmt.Println(s)
	fmt.Println(f)
}
