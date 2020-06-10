package host

import (
	"fmt"
	"net/http"

	"github.com/sparrc/go-ping"
)

//PingHost performs an ICMP ping on a particular host
func PingHost(host string) error {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pinger.Count = 3
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		fmt.Println("ICMP enabled")
	}

	return nil
}

//ProbeHTTP probes a host for information regarding the HTTP protocol
func ProbeHTTP(host string) error {
	//TODO fix hardcoded protocol prefix.
	resp, err := http.Get("https://" + host)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	fmt.Printf("HTTP version: %s\n", resp.Proto)
	fmt.Printf("HTTP response status: %s\n", resp.Status)
	return nil
}
