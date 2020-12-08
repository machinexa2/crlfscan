package runner

import (
	"fmt"
	"net/http"
	"time"
)

func checkHeader(resp *http.Response) bool{
	for r := resp; r != nil; r = r.Request.Response {
		if s := r.Header.Get("crlfscan"); s != "" {
			return true
		}
	}
	return false
}

func Scan(scanUrl string) string {
	if Settings.Verbose != false {
		fmt.Printf("%s Trying url: %s\n", Color.Information, Color.Blue(scanUrl))
	}
	var httpClient = &http.Client{
		Timeout: time.Second * 16,
	}
	resp, err := httpClient.Get(scanUrl)
	if err != nil {
		fmt.Printf("%s Error fetching response of: %s\n", Color.Bad, Color.Red(scanUrl))
	}
	vuln := checkHeader(resp);
	if vuln == true {
		return scanUrl
	}
	return "";
}
