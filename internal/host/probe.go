package host

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/pimg/host-analyzer/internal/constants"
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

	var myClient *http.Client

	//Way to set the TLS version on the HTTP client
	defaultTransport := http.DefaultTransport.(*http.Transport) // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.TLSClientConfig = &tls.Config{MaxVersion: tls.VersionTLS11}

	//TODO cycle through TLS versions

	myClient = &http.Client{Transport: defaultTransport}

	//TODO fix hardcoded protocol prefix.
	resp, err := myClient.Get("https://" + host)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	fmt.Printf("HTTP version: %s\n", resp.Proto)
	fmt.Printf("HTTP response status: %s\n", resp.Status)
	fmt.Printf("The Cipher suite is: %s\n", constants.CipherSuites[resp.TLS.CipherSuite])
	fmt.Printf("The TLS version is: %s\n", constants.TLSVersions[resp.TLS.Version])
	return nil
}
