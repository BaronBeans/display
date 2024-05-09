package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	ipAddr := getIpAddress()
	fmt.Println(ipAddr)
}

func getIpAddress() string {
	ipconfig, err := exec.Command("ipconfig.exe").Output()
	if err != nil {
		log.Fatal(err)
	}

	var ipAddr string
	lines := strings.Split(string(ipconfig), "\n")
	found := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "WSL") {
			found = true
			continue
		}
		if found {
			if strings.Contains(line, "IPv4 Address") {
				ipAddr = strings.Split(line, ":")[1]
				ipAddr = strings.TrimSpace(ipAddr)
				break
			}
		}
	}

	return ipAddr
}
