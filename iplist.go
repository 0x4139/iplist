package iplist

import (
	"io"
	"bufio"
	"fmt"
	"strings"
	"net"
)

type IPList struct {
	networks []*net.IPNet
}

func (iplist *IPList) Contains(ip string) (bool) {
	parsedIP := net.ParseIP(ip)
	for _, network := range iplist.networks {
		if network.Contains(parsedIP) {
			return true
		}
	}
	return false
}

func NewFromReader(reader io.Reader) (ipList *IPList, err error) {
	scanner := bufio.NewScanner(reader)
	lineNum := 1
	ipList = &IPList{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		network, lineErr := parseLine(line)
		if lineErr != nil {
			err = fmt.Errorf("error parsing line %d: %s", lineNum, lineErr)
			return
		}
		lineNum++
		ipList.networks = append(ipList.networks, network)
	}
	err = scanner.Err()
	if err != nil {
		return
	}

	return
}

func parseLine(line string) (network *net.IPNet, err error) {
	line = strings.TrimSpace(line)

	if !strings.Contains(line, "/") {
		line += "/32"
	}
	_, network, err = net.ParseCIDR(line)
	return
}
