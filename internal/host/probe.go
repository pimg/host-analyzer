package host

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

//ProbeTLS probes a host for information regarding the TLS configuration, like TLS version and supported Cipher suites.
func ProbeTLS(host string, port int) error {

	for tlsVersion, tlsVersionString := range constants.TLSVersions {
		fmt.Println("Trying with TLS version: ", tlsVersionString)

		conf := &tls.Config{
			MaxVersion:               tlsVersion,
			PreferServerCipherSuites: false,
		}
		conn, err := tls.Dial("tcp", host+":"+strconv.Itoa(port), conf)
		if err != nil {
			log.Println(tlsVersionString + "is not supported.")
		} else {
			fmt.Println(tlsVersionString + "is supported by this site.")
			conn.Close()
		}

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
	fmt.Printf("The default Cipher suite is: %s\n", constants.CipherSuites[resp.TLS.CipherSuite])
	fmt.Printf("The default TLS version is: %s\n", constants.TLSVersions[resp.TLS.Version])
	return nil
}
