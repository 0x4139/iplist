package iplist

import (
	"testing"
	"net/http"
)

func TestShouldCheckIfIpIsFoundInBlockList(t *testing.T) {
	response, err := http.Get("https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset")
	if err != nil {
		t.Fatalf("error while obtaining the list %s", err.Error())
	}
	defer response.Body.Close()

	iplist, err := NewFromReader(response.Body)
	if err != nil {
		t.Fatalf("Error while building the iplist %s", err.Error())
	}
	ip := "2.50.11.28"

	if iplist.Contains(ip) == false {
		t.Fatalf("List should contain the ip %s", ip)
	}

}

func TestShouldCheckIfIpIsFoundInBlockListRange(t *testing.T) {
	response, err := http.Get("https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset")
	if err != nil {
		t.Fatalf("error while obtaining the list %s", err.Error())
	}
	defer response.Body.Close()

	iplist, err := NewFromReader(response.Body)
	if err != nil {
		t.Fatalf("Error while building the iplist %s", err.Error())
	}
	ip := "109.232.160.1"

	if iplist.Contains(ip) == false {
		t.Fatalf("List should contain the ip %s", ip)
	}

}

func TestShouldCheckIfIpIsNotFoundInBlockList(t *testing.T) {
	response, err := http.Get("https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset")
	if err != nil {
		t.Fatalf("error while obtaining the list %s", err.Error())
	}
	defer response.Body.Close()

	iplist, err := NewFromReader(response.Body)
	if err != nil {
		t.Fatalf("Error while building the iplist %s", err.Error())
	}
	ip := "93.120.46.119"
	if iplist.Contains(ip) == true {
		t.Fatalf("List should not contain the ip %s", ip)
	}

}
